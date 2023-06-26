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

	req := &proto.GetFullURLRequest{
		ShortUrl: "zzzzzzzzzz",
	}

	resp, err := client.GetFullURL(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)

}
