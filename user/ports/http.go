package ports

import "github.com/rubemlrm/go-api-bootstrap/user/app"

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}
