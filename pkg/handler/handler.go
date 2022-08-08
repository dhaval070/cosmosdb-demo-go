package handler

import (
	"cosmosdb-demo/pkg/handler/interfaces"
	"io"
	"net/http"

	hello "cosmosdb-demo/pkg/service"
	service "cosmosdb-demo/pkg/service/interfaces"

	"github.com/google/wire"
)

type api struct {
	service service.Hello
}

func NewHandler(s service.Hello) *api {
	return &api{
		s,
	}
}

func (h api) Hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, h.service.Hello())
}

var Wired = wire.NewSet(NewHandler, hello.Wired, wire.Bind(new(interfaces.Handler), new(*api)))
