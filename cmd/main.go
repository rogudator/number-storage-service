package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/rogudator/number-storage-service/internal/repository"
	"github.com/rogudator/number-storage-service/internal/service"
	"github.com/rogudator/number-storage-service/internal/transport"
	"github.com/rogudator/number-storage-service/pkg/logger"
	"github.com/rogudator/number-storage-service/proto/number_storage"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	// First we initialize logger. It is logging all the information to logs folder.
	log := logger.NewLogger()
	log.Info("Starting number-service-storage...")

	// Connecting to database
	db, err := repository.NewPosgresDB(initConfig())
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	// Project is seperated in three layers. It makes it easier for dependency injection.
	// Repository layer is repsonsible for database communication.
	repos := repository.NewRepository(db)
	// Service layer stores business logic.
	services := service.NewService(repos)
	// Transport layer recieves requests and sends responds.
	api := transport.NewTransport(services)

	// Make a listener for app shutdown signal.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt)

	// We start REST api and gRPC api concurrently. If something
	// goes wrong on any goroutine we have error channel to shutdown
	// service gracefully.
	errChannel := make(chan error)
	grpcServer := runGrpc(api, log, errChannel)
	restServer := runRest(api, log, errChannel)

	select {
	case s := <-sig:
		log.Error("Error from channel: ", s)
	case err := <-errChannel:
		log.Error("Got unexpected error: ", err)
	}

	// Graceful shutdown in case of REST or gRPC malfunction
	if restServer != nil {
		if err := restServer.Shutdown(context.Background()); err != nil {
			closeDB(db, log)
			log.Error("err closing rest server", err)
		}
		log.Info("REST connection gracefully closed")
	}
	if grpcServer != nil {
		closeDB(db, log)
		grpcServer.GracefulStop()
		log.Info("gRPC connection gracefully closed")
	}

}

// This function is needed to start gRPC api concurrently.
func runGrpc(api *transport.Transport, log *zap.SugaredLogger, errCh chan error) *grpc.Server {

	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		errCh <- err
		return nil
	}

	gRPCServer := grpc.NewServer()

	number_storage.RegisterNumberStorageServiceServer(gRPCServer, api)

	go func() {
		errCh <- gRPCServer.Serve(listen)
	}()

	log.Info("gRPC started on: 0.0.0.0:50051")

	return gRPCServer
}

// This function is needed to start REST api concurrently.
func runRest(api *transport.Transport, log *zap.SugaredLogger, errCh chan error) *http.Server {

	mux := runtime.NewServeMux()

	restServer := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	err := number_storage.RegisterNumberStorageServiceHandlerServer(context.Background(), mux, api)
	if err != nil {
		errCh <- err
		return nil
	}

	go func() {
		errCh <- restServer.ListenAndServe()
	}()

	log.Info("REST started on 0.0.0.0:8080")

	return &restServer
}

// Configuration for database connection
func initConfig() repository.Config {
	return repository.Config{
		Host:     "database",
		Port:     "5432",
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  "disable",
	}
}

func closeDB(db *sqlx.DB, log *zap.SugaredLogger) {
	if err := db.Close(); err != nil {
		log.Error("an err while closing db connection: %s", err.Error())
	} else {
		log.Info("db connection gracefully closed")
	}
}