package query

import "context"

type AllUsers struct{}

type AllUsersHandler struct {
	allUsersHandler
}

type allUsersHandler struct {
	readModel AllUsersReadModel
}

type AllUsersReadModel interface {
	AllUsers(ctx context.Context) ([]User, error)
}

func (h allUsersHandler) Handle(ctx context.Context, _ AllUsers) (ur []User, err error) {
	return h.readModel.AllUsers(ctx)
}

func NewAllUsersHandler() AllUsersHandler {
	return AllUsersHandler{}
}
