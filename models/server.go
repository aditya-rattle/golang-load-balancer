package models

import (
	"net/http/httputil"
	"net/url"

	exceptionhandler "loadbalancer.com/exceptionHandler"
)

type Server struct {
	Address string
	Proxy *httputil.ReverseProxy
}

func NewServer(addr string) *Server {
	serverUrl ,err := url.Parse(addr)
	exceptionhandler.HandleError(err)
	return &Server{
		Address: addr,
		Proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func (server *Server) IsAlive() bool {
	return true
}

func (server *Server) GetAddress() string {
	return server.Address
}