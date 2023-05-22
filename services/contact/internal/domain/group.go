package domain

import "advanced.microservices/pkg/validator"

type Group struct {
	ID        int64  `json:"id"`
	GroupName string `json:"group_name"`
}

func ValidateGroupName(v *validator.Validator, groupName string) {
	v.Check(len(groupName) <= 250, "group name", "must not be longer than 250 characters")
}
