package usecase

import (
	"context"
	"fmt"
	"github.com/evrone/go-clean-template/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WeatherUseCase struct {
	repo   WeatherRepo
	webApi WeatherWebAPI
}

func New1(r WeatherRepo, w WeatherWebAPI) *WeatherUseCase {
	return &WeatherUseCase{
		r,
		w,
	}
}

func (wu *WeatherUseCase) Get(ctx context.Context, city string) (entity.WeatherData, error) {
	res, err := wu.repo.Get(ctx, city)
	if err != nil {
		data, err := wu.webApi.Get(ctx, city)
		if err != nil {
			return entity.WeatherData{}, fmt.Errorf("failed to Get err: %v", err)
		}

		data.ID = primitive.NewObjectID()

		_, err = wu.repo.Add(ctx, data)
		if err != nil {
			return entity.WeatherData{}, fmt.Errorf("failed to Add err: %v", err)
		}

		res = data
	}

	return res, nil
}

func (wu *WeatherUseCase) Update(ctx context.Context, city string) error {
	data, err := wu.webApi.Get(ctx, city)
	if err != nil {
		return fmt.Errorf("failed to Get err: %v", err)
	}

	data.ID = primitive.NewObjectID()

	err = wu.repo.Update(ctx, city, data)
	if err != nil {
		return fmt.Errorf("failed to Update err: %v", err)
	}

	return nil
}
