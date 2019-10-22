package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	pb "github.com/algorithm/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
)

type searchService struct {

}

func (s *searchService) Search(ctx context.Context,r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response:r.GetRequest()+" Server"}, nil
}

const PORT = ":9001"
func main() {
	cert,err := tls.LoadX509KeyPair("conf/server/server.pem","conf/server/server.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPari error:%+v\n",err.Error())
	}
	certPool := x509.NewCertPool()
	ca,err := ioutil.ReadFile("conf/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error: %+v \n",err.Error())
	}
	if ok := certPool.AppendCertsFromPEM(ca);!ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}
	c := credentials.NewTLS(&tls.Config{
		Certificates:[]tls.Certificate{cert},
		ClientAuth:tls.RequireAndVerifyClientCert,
		ClientCAs:certPool,
	})

	server := grpc.NewServer(grpc.Creds(c))
	pb.RegisterSearchServiceServer(server,&searchService{})
	l,err := net.Listen("tcp",PORT)
	if err != nil {
		log.Fatalf("net.Listen err %v: " + err.Error())
	}
	server.Serve(l)
}
