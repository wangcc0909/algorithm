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
)

const PORT = ":9001"

func main() {
	cert,err := tls.LoadX509KeyPair("conf/client/client.pem","conf/client/client.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %+v \n",err.Error())
	}
	certPool := x509.NewCertPool()
	ca,err := ioutil.ReadFile("conf/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v\n",err.Error())
	}
	if ok := certPool.AppendCertsFromPEM(ca);!ok {
		log.Fatalf("certPool.AppendCertsFromPEM err \n")
	}
	c := credentials.NewTLS(&tls.Config{
		Certificates:[]tls.Certificate{cert},
		ServerName:"algorithm",
		RootCAs:certPool,
	})

	conn,err := grpc.Dial(PORT,grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial error %+v: " + err.Error())
	}
	defer conn.Close()
	client := pb.NewSearchServiceClient(conn)
	resp,err := client.Search(context.Background(),&pb.SearchRequest{
		Request:"gRPC",
	})
	if err != nil {
		log.Printf("client.Search error %+v: \n",err.Error())
	}
	log.Printf("resp: %s",resp.GetResponse())
}