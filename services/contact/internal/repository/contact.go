package repository

import "database/sql"

type ContactRepository struct {
	DB *sql.DB
}
