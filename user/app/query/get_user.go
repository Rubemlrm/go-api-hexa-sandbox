package query

import "context"

type GetUser struct{}

type getUserHandler struct {
	readModel GetUserReadModel
}

type GetUserReadModel interface {
	GetUser(ctx context.Context) (User, error)
}

func (h getUserHandler) Handle(ctx context.Context, _ GetUser) (u User, err error) {
	return h.readModel.GetUser(ctx)
}
