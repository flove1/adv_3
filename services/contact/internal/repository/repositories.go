package repository

import (
	"database/sql"
)

type Repositories struct {
	Contacts ContactRepository
	Groups   GroupRepository
}

func NewRepositories(db *sql.DB) Repositories {
	return Repositories{
		Contacts: ContactRepository{DB: db},
		Groups:   GroupRepository{DB: db},
	}
}
