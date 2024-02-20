package repo

import (
	"context"
	"fmt"
	"github.com/evrone/go-clean-template/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WeatherRepo struct {
	db *mongo.Collection
}

func New1(db *mongo.Database) *WeatherRepo {
	return &WeatherRepo{
		db.Collection("weather_data"),
	}
}

func (wr *WeatherRepo) Add(ctx context.Context, data entity.WeatherData) (primitive.ObjectID, error) {
	res, err := wr.db.InsertOne(ctx, data)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to InsertOne err: %v", err)
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (wr *WeatherRepo) Get(ctx context.Context, city string) (entity.WeatherData, error) {
	var dest entity.WeatherData
	err := wr.db.FindOne(ctx, bson.D{{"city", city}}).Decode(&dest)
	if err != nil {
		return entity.WeatherData{}, fmt.Errorf("failed to FindOne err: %v", err)
	}

	return dest, nil
}

func (wr *WeatherRepo) Update(ctx context.Context, city string, data entity.WeatherData) error {
	args := prepareArgs(data)
	if len(args) > 0 {
		out, err := wr.db.UpdateOne(ctx, bson.M{"city": city}, bson.M{"$set": args})
		if err != nil {
			return fmt.Errorf("failed to UpdateOne err: %v", err)
		}

		if out.MatchedCount == 0 {
			return fmt.Errorf("failed to MatchedCount err: %v", err)
		}
	}

	return nil
}

func prepareArgs(data entity.WeatherData) (args bson.M) {
	args = make(bson.M)

	if data.Temp != 0.0 {
		args["temp"] = data.Temp
	}
	if data.TempMax != 0.0 {
		args["temp_max"] = data.TempMax
	}
	if data.TempMin != 0.0 {
		args["temp_min"] = data.TempMin
	}
	if data.Pressure != 0.0 {
		args["pressure"] = data.Pressure
	}
	if data.Humidity != 0.0 {
		args["humidity"] = data.Humidity
	}
	if data.WindSpeed != 0.0 {
		args["wind_speed"] = data.WindSpeed
	}
	return
}
