package repository

import (
	"context"
	"database/sql"

	"advanced.microservices/services/contact/internal/domain"
)

type SQLContactRepository struct {
	DB *sql.DB
}

func NewContactRepository(conn *sql.DB) domain.ContactRepository {
	return &SQLContactRepository{conn}
}

// GetByID implements domain.ContactRepository
func (*SQLContactRepository) GetByID(ctx context.Context, id int64) (domain.Contact, error) {
	panic("unimplemented")
}

// GetByPhone implements domain.ContactRepository
func (*SQLContactRepository) GetByPhone(ctx context.Context, title string) (domain.Contact, error) {
	panic("unimplemented")
}

// List implements domain.ContactRepository
func (*SQLContactRepository) List(ctx context.Context) ([]domain.Contact, error) {
	panic("unimplemented")
}

// Update implements domain.ContactRepository
func (*SQLContactRepository) Update(ctx context.Context, contact *domain.Contact) error {
	panic("unimplemented")
}

// Delete implements domain.ContactRepository
func (*SQLContactRepository) Delete(ctx context.Context, id int64) error {
	panic("unimplemented")
}
