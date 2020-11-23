package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"question_2/db"
	"question_2/handler"
	"question_2/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	const portREST string = ":9000"
	const portGRPC string = ":4040"
	err := checkEnv()
	if err != nil {
		log.Fatal(err)
	}
	dc := db.Connect()
	defer dc.Close()

	forever := make(chan bool)
	// grpc
	go func() {
		log.Println("GRPC Server is running at port", portGRPC)
		listener, err := net.Listen("tcp", portGRPC)
		if err != nil {
			log.Fatal(err)
		}
		srv := grpc.NewServer()
		pb.RegisterMovieServiceServer(srv, &handler.GrpcServer{Db: dc})
		reflection.Register(srv)
		if err := srv.Serve(listener); err != nil {
			panic(err)
		}
	}()

	// rest
	go func() {
		h := handler.NewRestServer(dc)
		http.HandleFunc("/", h.GetData)
		log.Println("REST Server is running at port", portREST)
		log.Fatal(http.ListenAndServe(portREST, nil))
	}()
	<-forever
}

func checkEnv() error {
	log.Println("Checking environment.....")

	_, ok := os.LookupEnv("DATA_SOURCE")
	if !ok {
		return errors.New("env:DATA_SOURCE does not exist")
	}
	_, ok = os.LookupEnv("BASE_URL")
	if !ok {
		return errors.New("env:BASE_URL does not exist")
	}
	_, ok = os.LookupEnv("API_KEY")
	if !ok {
		return errors.New("env:API_KEY does not exist")
	}

	log.Println("Environment is complete.")
	return nil
}
