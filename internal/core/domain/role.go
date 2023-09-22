package domain

type Role string

const (
	SyncSecrets   Role = "sync secrets"
	ManageUsers   Role = "manage-users"
	ManageGroups  Role = "manage-groups"
	ReadSecrets   Role = "read secrets"
	WriteSecrets  Role = "write secrets"
	DeleteSecrets Role = "delete secrets"
)
