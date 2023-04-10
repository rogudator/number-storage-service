package repository

import "github.com/jmoiron/sqlx"

type NumberStoragePostgreSQL struct {
	db *sqlx.DB
}

func NewNumberStoragePostgreSQL(db *sqlx.DB) *NumberStoragePostgreSQL {
	return &NumberStoragePostgreSQL{db: db}
}

type NumberStorage interface {
	GetNumber() (int64, error)
	UpdateNumber(number int64) (error)
}

// This method gets stored number from db.
func (r *NumberStoragePostgreSQL) GetNumber() (int64, error) {
	query := `
	SELECT number FROM storage
	WHERE id=$1;
	`

	var number int64
	if err := r.db.Get(&number, query, 1); err != nil {
		return 0, err
	}

	return number, nil
}

// This method gets the calculated new number and updates the stored number.
func (r *NumberStoragePostgreSQL) UpdateNumber(number int64) (error) {
	query := `
	UPDATE storage
	SET number = $1
	WHERE id = 1
	`
	_, err := r.db.Exec(query, number)
	if err != nil {
		return err
	}
	
	return nil
}