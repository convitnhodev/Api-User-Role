package common

const (
	DbTypeUser = iota + 1
	DbTypeRole
	DbTypePermission
	DbTypeDepartment
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
