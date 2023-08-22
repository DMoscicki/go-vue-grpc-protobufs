package main

import (
	"blog/handlers"
	"blog/protogen"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/lmittmann/tint"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {

	w := os.Stderr

	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	var s = handlers.Server{}

	var grpcServer = grpc.NewServer()

	var ToCors = grpcweb.WrapServer(grpcServer)

	var httpServer = &http.Server{
		Addr: "localhost:8080",
	}

	var http2Server = &http2.Server{}

	_ = http2.ConfigureServer(httpServer, http2Server)

	protogen.RegisterTodoServiceServer(grpcServer, &s)

	reflection.Register(grpcServer)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		allowcors(writer, request)
		if ToCors.IsGrpcWebRequest(request) || ToCors.IsAcceptableGrpcCorsRequest(request) {
			ToCors.ServeHTTP(writer, request)
		}
	})

	slog.Info("HTTP server listening on " + "8080")

	if err := httpServer.ListenAndServeTLS("./certs/cert.crt", "./certs/newkey.key"); err != nil {
		slog.Error(err.Error())
	}

	//log.Fatal(httpServer.ListenAndServeTLS("./certs/cert.crt", "./certs/newkey.key"))

}

func allowcors(resp http.ResponseWriter, _ *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	resp.Header().Set("Access-Control-Expose-Headers", "grpc-status, grpc-message")
	resp.Header().Set("Content-Type", "application/grpc")
	resp.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, XMLHttpRequest, x-user-agent, x-grpc-web, grpc-status, grpc-message")
}
