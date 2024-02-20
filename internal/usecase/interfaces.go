// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"github.com/evrone/go-clean-template/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	Weather interface {
		Get(ctx context.Context, city string) (entity.WeatherData, error)
		Update(ctx context.Context, city string) error
	}

	WeatherRepo interface {
		Add(ctx context.Context, data entity.WeatherData) (primitive.ObjectID, error)
		Get(ctx context.Context, city string) (entity.WeatherData, error)
		Update(ctx context.Context, city string, translation entity.WeatherData) error
	}

	WeatherWebAPI interface {
		Get(ctx context.Context, city string) (entity.WeatherData, error)
	}
)
