package service

import "net/http"

type Service struct {
	name   string
	port   string
	server *http.Server
}
