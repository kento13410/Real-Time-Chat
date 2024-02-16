package grpc

type Server struct {
	UnimplementedAuthServiceServer
}

func NewServer() *Server {
	return &Server{}
}