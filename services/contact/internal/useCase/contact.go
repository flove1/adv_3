package useCase

import (
	"context"
	"time"

	"advanced.microservices/services/contact/internal/domain"
)

type contactUsecase struct {
	contactRepo    domain.ContactRepository
	contextTimeout time.Duration
}

// NewArticleUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewContactUsecase(c domain.ContactRepository, timeout time.Duration) domain.ContactUseCase {
	return &contactUsecase{
		contactRepo:    c,
		contextTimeout: timeout,
	}
}

// GetByID implements domain.ContactUseCase
func (*contactUsecase) GetByID(ctx context.Context, id int64) (domain.Contact, error) {
	panic("unimplemented")
}

// GetByPhone implements domain.ContactUseCase
func (*contactUsecase) GetByPhone(ctx context.Context, title string) (domain.Contact, error) {
	panic("unimplemented")
}

// List implements domain.ContactUseCase
func (*contactUsecase) List(ctx context.Context) ([]domain.Contact, error) {
	panic("unimplemented")
}

// Update implements domain.ContactUseCase
func (*contactUsecase) Update(ctx context.Context, contact *domain.Contact) error {
	panic("unimplemented")
}

// Delete implements domain.ContactUseCase
func (*contactUsecase) Delete(ctx context.Context, id int64) error {
	panic("unimplemented")
}
