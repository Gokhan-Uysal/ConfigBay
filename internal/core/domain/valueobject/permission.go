package valueobject

import (
	"errors"
)

const (
	ReadProject   Permission = "read-project"
	ManageUsers   Permission = "manage-users"
	ManageGroups  Permission = "manage-groups"
	ReadSecrets   Permission = "read-secrets"
	WriteSecrets  Permission = "write-secrets"
	DeleteSecrets Permission = "delete-secrets"
)

type (
	Permission string
)

func ToPermission(name string) (Permission, error) {
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

func ToString(name Permission) string {
	return string(name)
}
