package interfaces

import "net/http"

type Handler interface {
	Hello(w http.ResponseWriter, req *http.Request)
	Family(w http.ResponseWriter, req *http.Request)
}
