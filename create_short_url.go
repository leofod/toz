package main

import (
	"context"
	"fmt"

	proto "toz/pkg/proto" // Импортируйте ваш пакет прото

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := proto.NewUrlClient(conn)

	req := &proto.CreateShotrUrlRequest{
		FullUrl: "github.com/asd",
	}

	resp, err := client.CreateShotrUrl(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)

}
