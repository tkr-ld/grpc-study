package handler

import (
	"context"
	grpc_weather "github.com/tkr-ld/grpc-study/server/gen/api/proto"
)

type WeatherHandler struct {
	grpc_weather.UnimplementedWeatherServiceServer
}

func (h *WeatherHandler) GetWeather(context.Context, *grpc_weather.GetRequest) (*grpc_weather.WeatherResponse, error) {
	// 一旦固定で返す
	return &grpc_weather.WeatherResponse{
		Weather: grpc_weather.WeatherType_SUNNY,
	}, nil
}

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{}
}
