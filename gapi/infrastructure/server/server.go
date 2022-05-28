package server

type Server interface {
	StartServer() error
	StopServer() error
}
