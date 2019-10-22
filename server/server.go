package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	pb "github.com/algorithm/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
)

type searchService struct {
}

func (s *searchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	//strconv.Atoi(r.GetRequest())
	if r.GetRequest() == "gRPC" {
		panic("dev")
	}
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = ":9001"

func main() {
	c,err := GetTLSCredentialsByCA()
	if err != nil {
		log.Fatalf("GetTLSCredentialsByCA err: %v",err.Error())
	}
	entry := logrus.WithField("algorithm","grpc")
	opts := []grpc.ServerOption{
		grpc.Creds(c),
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(entry)),
	}
	server := grpc.NewServer(opts...)
	pb.RegisterSearchServiceServer(server, &searchService{})
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("net.Listen err %v: " + err.Error())
	}
	server.Serve(l)
}

func GetTLSCredentialsByCA() (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair("conf/server/server.pem", "conf/server/server.key")
	if err != nil {
		log.Printf("tls.LoadX509KeyPari error:%+v\n", err.Error())
		return nil, err
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("conf/ca.pem")
	if err != nil {
		log.Printf("ioutil.ReadFile error: %+v \n", err.Error())
		return nil, err
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Println("certPool.AppendCertsFromPEM err")
		return nil, err
	}
	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	return c,nil
}
