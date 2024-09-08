package app

import "github.com/rubemlrm/go-api-bootstrap/user/app/query"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct{}
type Queries struct {
	AllUsers query.AllUsersHandler
}
