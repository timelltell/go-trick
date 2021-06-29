package Compare

import (
	"time"
)

type Server struct {
	Addr string
	Port int
	Protocol string
	Timeout time.Duration
	Maxconns int
}
type Option func(* Server)

func Protocol(p string) Option{
	return func(s *Server){
		s.Protocol=p
	}
}
func Timeout(p time.Duration) Option{
	return func(s *Server){
		s.Timeout=p
	}
}
func Maxconns(p int) Option{
	return func(s *Server){
		s.Maxconns=p
	}
}
func NewServer(addr string,port int,options ...Option)(*Server,error){
	srv:=Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		Maxconns: 1000,
	}
	for _,option:=range options{
		option(&srv)
	}
	return &srv,nil
}
