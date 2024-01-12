package main

import (
	"context"
	"fmt"
	grpc_weather "github.com/tkr-ld/grpc-study/client/gen/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := grpc_weather.NewWeatherServiceClient(conn)

	var number int32
	_, err = fmt.Scan(&number)
	if err != nil {
		panic(err)
	}

	req := &grpc_weather.GetRequest{
		Number: number,
	}

	res, err := client.GetWeather(context.Background(), req)
	if err != nil {
		fmt.Printf("error")
	}

	fmt.Printf("weather: %v\n", res.GetWeather())
}
