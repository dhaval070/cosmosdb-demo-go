package api

import (
	"context"
	"cosmosdb-demo/pkg/handler/interfaces"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer(h interfaces.Handler) *Server {
	srv := http.NewServeMux()
	srv.HandleFunc("/hello", h.Hello)
	srv.HandleFunc("/family", h.Family)

	return &Server{
		&http.Server{
			Addr:    ":8080",
			Handler: srv,
		},
	}
}

func (s *Server) Run() {
	go func() {
		log.Fatal(s.server.ListenAndServe())
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	<-c
	log.Println("bye")

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	s.server.Shutdown(ctx)

}
