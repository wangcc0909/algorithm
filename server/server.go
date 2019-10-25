package main

import (
	"context"
	"github.com/algorithm/gtls"
	pb "github.com/algorithm/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strings"
)

type searchService struct {
}

func (s *searchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	//strconv.Atoi(r.GetRequest())
	/*if r.GetRequest() == "gRPC" {
		panic("dev")
	}*/
	return &pb.SearchResponse{Response: r.GetRequest() + " HTTP Server"}, nil
}

const PORT = ":9001"

func main() {
	certFile := "conf/server/server.pem"
	keyFile := "conf/server/server.key"
	t := gtls.Server{CerFile:certFile,KeyFile:keyFile}
	c,err := t.GetTLSCredentials()
	if err != nil {
		log.Fatalf("GetTLSCredentialsByCA err: %v",err.Error())
	}
	mux := GetHTTPServerMux()
	entry := logrus.WithField("algorithm","grpc")
	opts := []grpc.ServerOption{
		grpc.Creds(c),
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(entry)),
	}
	server := grpc.NewServer(opts...)
	pb.RegisterSearchServiceServer(server, &searchService{})
	/*l, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("net.Listen err %v: " + err.Error())
	}
	server.Serve(l)*/
	err = http.ListenAndServeTLS(PORT,
		certFile,keyFile,
		http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"),"application/grpc") {
				server.ServeHTTP(w,r)
			}else {
				mux.ServeHTTP(w,r)
			}
		}))
	log.Fatalf("http.ListenAndServeTLS error: %v",err.Error())
}

func GetHTTPServerMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("我是http请求结果"))
	})
	return mux
}