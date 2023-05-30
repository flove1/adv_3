package useCase

import (
	"context"
	"time"

	"advanced.microservices/services/contact/internal/domain"
)

type groupUsecase struct {
	groupRepo      domain.GroupRepository
	contextTimeout time.Duration
}

// NewArticleUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewGroupUsecase(c domain.GroupRepository, timeout time.Duration) domain.GroupUseCase {
	return &groupUsecase{
		groupRepo:      c,
		contextTimeout: timeout,
	}
}

// GetByID implements domain.GroupUseCase
func (*groupUsecase) GetByID(ctx context.Context, id int64) (domain.Group, error) {
	panic("unimplemented")
}

// GetByPhone implements domain.GroupUseCase
func (*groupUsecase) GetByPhone(ctx context.Context, title string) (domain.Group, error) {
	panic("unimplemented")
}

// List implements domain.GroupUseCase
func (*groupUsecase) List(ctx context.Context) ([]domain.Group, error) {
	panic("unimplemented")
}

// Update implements domain.GroupUseCase
func (*groupUsecase) Update(ctx context.Context, contact *domain.Group) error {
	panic("unimplemented")
}

// Delete implements domain.GroupUseCase
func (*groupUsecase) Delete(ctx context.Context, id int64) error {
	panic("unimplemented")
}
