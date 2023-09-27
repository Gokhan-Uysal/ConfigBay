package valueobject

import (
	"errors"
)

const (
	ReadProject   Role = "read-project"
	ManageUsers   Role = "manage-users"
	ManageGroups  Role = "manage-groups"
	ReadSecrets   Role = "read-secrets"
	WriteSecrets  Role = "write-secrets"
	DeleteSecrets Role = "delete-secrets"
)

type (
	Role string
)

func ToRoleName(name string) (Role, error) {
	switch name {
	case string(ReadProject):
		return ReadProject, nil
	case string(ManageUsers):
		return ManageUsers, nil
	case string(ManageGroups):
		return ManageGroups, nil
	case string(ReadSecrets):
		return ReadSecrets, nil
	case string(WriteSecrets):
		return WriteSecrets, nil
	case string(DeleteSecrets):
		return DeleteSecrets, nil
	default:
		return "", errors.New("invalid role name")
	}
}

func ToString(name Role) string {
	return string(name)
}
