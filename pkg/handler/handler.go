package handler

import (
	"cosmosdb-demo/domain"
	"cosmosdb-demo/pkg/handler/interfaces"
	"encoding/json"
	"io"
	"net/http"

	hello "cosmosdb-demo/pkg/service"
	service "cosmosdb-demo/pkg/service/interfaces"

	"github.com/google/wire"
)

type handler struct {
	service service.Hello
}

func NewHandler(s service.Hello) *handler {
	return &handler{
		s,
	}
}

func (h handler) Hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, h.service.Hello())
}

func (h handler) Family(w http.ResponseWriter, req *http.Request) {
	family, err := h.service.GetFamily("AndersenFamily")

	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	resp := familyToAPIFamily(family)

	data, err := json.Marshal(resp)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, string(data))

}

func familyToAPIFamily(m *domain.Family) *domain.APIFamily {
	if m == nil {
		return nil
	}

	return &domain.APIFamily{
		ID:       m.ID,
		LastName: m.LastName,
		Address:  m.Address,
	}
}

var Wired = wire.NewSet(NewHandler, hello.Wired, wire.Bind(new(interfaces.Handler), new(*handler)))
