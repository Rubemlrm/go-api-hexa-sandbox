package main

import (
	"context"

	"github.com/rubemlrm/go-api-bootstrap/user/service"
)

func main() {
	ctx := context.Background()
	app := service.NewApp(ctx)

}
