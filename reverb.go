// Package reverb - name is cheekily synonymized from the HTTPD package
// "Echo" and aims to serve a minimally similar purpose, though for gRPC
// instead of HTTP.
package reverb

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

// Reverb ...
type Reverb struct {
	TCPListener net.Listener
	GRPCServer  *grpc.Server
}

// New ...
func New() *Reverb {
	grpcServer := grpc.NewServer()
	return &Reverb{
		GRPCServer: grpcServer,
	}
}

// Start ...
func (r *Reverb) Start(opts string) *Reverb {
	server := New()
	tcpListener, err := net.Listen("tcp", opts)
	if err != nil {
		log.Fatalf("Reverb gRPC server failed to listen: %v", err)
	}
	server.TCPListener = tcpListener
	server.Register()
	return server
}

// Register ...
func (r *Reverb) Register() {

}

// Serve ...
func (r *Reverb) Serve() {
	r.GRPCServer.Serve(r.TCPListener)
}
