package repository

import (
	"context"
	"database/sql"

	"advanced.microservices/services/contact/internal/domain"
)

type SQLGroupRepository struct {
	DB *sql.DB
}

func NewGroupRepository(conn *sql.DB) domain.GroupRepository {
	return &SQLGroupRepository{conn}
}

// GetByID implements domain.GroupRepository
func (*SQLGroupRepository) GetByID(ctx context.Context, id int64) (domain.Group, error) {
	panic("unimplemented")
}

// GetByPhone implements domain.GroupRepository
func (*SQLGroupRepository) GetByGroupName(ctx context.Context, title string) (domain.Group, error) {
	panic("unimplemented")
}

// List implements domain.GroupRepository
func (*SQLGroupRepository) List(ctx context.Context) ([]domain.Group, error) {
	panic("unimplemented")
}

// Update implements domain.GroupRepository
func (*SQLGroupRepository) Update(ctx context.Context, Group *domain.Group) error {
	panic("unimplemented")
}

// Delete implements domain.GroupRepository
func (*SQLGroupRepository) Delete(ctx context.Context, id int64) error {
	panic("unimplemented")
}
