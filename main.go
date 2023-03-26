package main

import (
	"blog/handlers"
	"blog/protogen"
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {

	s := handlers.Server{}

	grpcServer := grpc.NewServer()

	ToCors := grpcweb.WrapServer(grpcServer)

	protogen.RegisterTodoServiceServer(grpcServer, &s)

	//reflection.Register(grpcServer)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		allowcors(writer, request)
		if ToCors.IsGrpcWebRequest(request) || ToCors.IsAcceptableGrpcCorsRequest(request) {
			ToCors.ServeHTTP(writer, request)
		}
	})

	fmt.Println("HTTP server listening on", ":8080")

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalf("Failed to start http server: %s ", err)
	}

}

func allowcors(resp http.ResponseWriter, _ *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	resp.Header().Set("Access-Control-Expose-Headers", "grpc-status, grpc-message")
	//resp.Header().Set("Content-Type", "application/grpc")
	resp.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, XMLHttpRequest, x-user-agent, x-grpc-web, grpc-status, grpc-message")
}
