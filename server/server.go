package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"simple_bank/generated/proto"
)

type SimpleBankGRPCServer struct {
	proto.UnimplementedSimpleBankServer
}

func StartGRPCServer(serverAddress string) {

	tcpListener, error := net.Listen("tcp", serverAddress)
	if error != nil {
		log.Println("❌ error creating tcp listener for gRPC server")

		log.Println(error.Error( )) }

	server := grpc.NewServer( )

	// enabling gRPC reflection (similar to self documentation)
	reflection.Register(server)

	proto.RegisterSimpleBankServer(server, &SimpleBankGRPCServer{ })

	log.Printf("🚀 starting gRPC server at %s", tcpListener.Addr( ).String( ))

	error= server.Serve(tcpListener)
	if error != nil {
		log.Println("❌ server error occured")

		log.Println(error.Error( )) }
}