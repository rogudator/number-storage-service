package repository

import "github.com/jmoiron/sqlx"

type NumberStorage interface {
	GetNumber() (int64, error)
	UpdateNumber(number int64) (int64, error)
}

type NumberStoragePostgreSQL struct {
	db *sqlx.DB
}

func NewNumberStoragePostgreSQL(db *sqlx.DB) *NumberStoragePostgreSQL {
	return &NumberStoragePostgreSQL{db: db}
}

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

func (r *NumberStoragePostgreSQL) UpdateNumber(number int64) (int64, error) {
	query := `
	UPDATE storage
	SET number = $1
	WHERE id = 1
	`
	_, err := r.db.Exec(query, number)
	if err != nil {
		return 0, err
	}
	
	return number, nil
}