package gtls

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

type Server struct {
	CaFile string
	CerFile string
	KeyFile string
}

func (t *Server) GetTLSCredentialsByCA() (credentials.TransportCredentials, error)  {
	cert, err := tls.LoadX509KeyPair(t.CerFile, t.KeyFile)
	if err != nil {
		log.Printf("tls.LoadX509KeyPari error:%+v\n", err.Error())
		return nil, err
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(t.CaFile)
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

func (t *Server) GetTLSCredentials() (credentials.TransportCredentials,error) {
	c,err := credentials.NewServerTLSFromFile(t.CerFile,t.KeyFile)
	if err != nil {
		return nil, err
	}
	return c, nil
}
