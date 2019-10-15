package main

import (
    "os"
    "fmt"
    "log"
    "net"
    //"time"
    //"context"
    "google.golang.org/grpc"
    "./protobuf"
    app "./app"
)

func main() {
    
    Host:=os.Getenv("HOST")
    Port:=os.Getenv("PORT")
    LocalPort:=os.Getenv("PORT_LOCAL")

    lis,err:= net.Listen("tcp", fmt.Sprintf(":%s", Port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

	grpcServer:= grpc.NewServer(
        grpc.UnaryInterceptor(protobuf.AccessLogInterceptor),
    )
    s:=&protobuf.Server{}

    //Подключение к БД
    s.DbConnect()
    defer s.DbDisconnect()

    //Подключение к сервису User
    s.UserConnect()

    //Подключение к сервису Category
    s.CategoryConnect()

    //Подключение к сервису Comment
    s.CommentConnect()

    // attach the user service to the server
    protobuf.RegisterPostServiceServer(grpcServer, s)
	
	log.Printf("%s service started on  %s:%s (localhos:%s)",app.SERVICE_NAME,Host,Port,LocalPort)
	err= grpcServer.Serve(lis)
    if err!= nil{
        log.Fatalf("failed to serve: %s", err)
    }
}

// func withServerUnaryInterceptor() grpc.ServerOption {
// 	return grpc.UnaryInterceptor(protobuf.AccessLogInterceptor)
// }