package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Option func(*Server)

type Server struct {
	mux    *http.ServeMux
	logger *log.Logger
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "GET" {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	return
	// }

	s.log("%s %s", r.Method, r.URL.Path)

	s.mux.ServeHTTP(w, r)
}

func (s *Server) log(format string, v ...interface{}) {
	s.logger.Printf(format+"\n", v...)
}

func Start(logger *log.Logger) *http.Server {
	return &http.Server{
		Addr:         getAddr(),
		Handler:      newServer(logWith(logger)),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func getAddr() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}

	return ":8383"
}

func newServer(options ...Option) *Server {
	s := &Server{logger: log.New(ioutil.Discard, "", 0)}

	for _, o := range options {
		o(s)
	}

	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/users/", GetUser)
	s.mux.HandleFunc("/users/all", GetUsers)
	s.mux.HandleFunc("/users/new", AddUser)
	s.mux.HandleFunc("/users/delete/", DeleteUser)

	return s
}

func logWith(logger *log.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}
