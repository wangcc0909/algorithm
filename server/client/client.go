package main

import (
	"context"
	"github.com/algorithm/gtls"
	pb "github.com/algorithm/proto"
	"google.golang.org/grpc"
	"log"
)

var PORT = ":9001"
func main() {
	tlsClient := gtls.Client{
		ServerName:"algorithm",
		CertFile:"conf/server/server.pem",
	}
	c,err := tlsClient.GetCredentialTLS()
	if err != nil {
		log.Fatalf("tls.GetCredentialTLS %v",err)
	}
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
