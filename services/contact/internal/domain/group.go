package domain

import (
	"context"

	"advanced.microservices/pkg/validator"
)

type Group struct {
	ID        int64  `json:"id"`
	GroupName string `json:"group_name"`
}
type GroupRepository interface {
	GetByID(ctx context.Context, id int64) (Group, error)
	GetByGroupName(ctx context.Context, title string) (Group, error)
	List(ctx context.Context) ([]Group, error)
	Update(ctx context.Context, Group *Group) error
	Delete(ctx context.Context, id int64) error
}

type GroupUseCase interface {
	GetByID(ctx context.Context, id int64) (Group, error)
	GetByPhone(ctx context.Context, title string) (Group, error)
	List(ctx context.Context) ([]Group, error)
	Update(ctx context.Context, contact *Group) error
	Delete(ctx context.Context, id int64) error
}

func ValidateGroupName(v *validator.Validator, groupName string) {
	v.Check(len(groupName) <= 250, "group name", "must not be longer than 250 characters")
}
