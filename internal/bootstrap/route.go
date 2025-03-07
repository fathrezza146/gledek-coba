// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"github.com/sev-2/raiden"
	"gledekcoba/internal/controllers"
	"gledekcoba/internal/models"
	"github.com/valyala/fasthttp"
)

func RegisterRoute(server *raiden.Server) {
	server.RegisterRoute([]*raiden.Route{
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/hello/{name}",
			Methods:    []string{fasthttp.MethodGet},
			Controller: &controllers.HelloWordController{},
		},
		{
			Type:       raiden.RouteTypeRest,
			Path:       "/todo",
			Methods:    []string{},
			Controller: &controllers.ToDoController{},
			Model:      models.ToDo{},
		},
	})
}
