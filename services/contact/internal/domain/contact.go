package domain

import (
	"context"
	"regexp"
	"strings"

	"advanced.microservices/pkg/validator"
)

type Contact struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

type ContactRepository interface {
	GetByID(ctx context.Context, id int64) (Contact, error)
	GetByPhone(ctx context.Context, title string) (Contact, error)
	List(ctx context.Context) ([]Contact, error)
	Update(ctx context.Context, contact *Contact) error
	Delete(ctx context.Context, id int64) error
}

type ContactUseCase interface {
	GetByID(ctx context.Context, id int64) (Contact, error)
	GetByPhone(ctx context.Context, title string) (Contact, error)
	List(ctx context.Context) ([]Contact, error)
	Update(ctx context.Context, contact *Contact) error
	Delete(ctx context.Context, id int64) error
}

func ValidateFullName(v *validator.Validator, fullName string) {
	v.Check(len(strings.Split(fullName, " ")) == 3, "full name", "full name must contain 3 parts")
}

func ValidatePhone(v *validator.Validator, phone string) {
	v.Check(validator.Matches(phone, regexp.MustCompile(`[0-9\[\]\(\\)\+\-]`)), "email", "must be a valid email address")
}
