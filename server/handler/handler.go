package handler

import (
	"context"
	grpc_weather "github.com/tkr-ld/grpc-study/server/gen/api/proto"
)

type WeatherHandler struct {
	grpc_weather.UnimplementedWeatherServiceServer
}

func (h *WeatherHandler) GetWeather(ctx context.Context, req *grpc_weather.GetRequest) (*grpc_weather.WeatherResponse, error) {
	// 一旦固定で返す
	var weatherType grpc_weather.WeatherType
	switch req.Number {
	case 1:
		weatherType = grpc_weather.WeatherType_SUNNY
	case 2:
		weatherType = grpc_weather.WeatherType_CLOUDY
	case 3:
		weatherType = grpc_weather.WeatherType_RAIN
	default:
		weatherType = grpc_weather.WeatherType_UNKNOWN
	}
	return &grpc_weather.WeatherResponse{
		Weather: weatherType,
	}, nil
}

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{}
}
