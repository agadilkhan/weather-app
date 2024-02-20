package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type WeatherData struct {
	ID        primitive.ObjectID `bson:"_id"`
	City      string             `bson:"city"`
	Temp      float64            `bson:"temp"`
	TempMax   float64            `bson:"temp_max"`
	TempMin   float64            `bson:"temp_min"`
	FeelsLike float64            `bson:"feels_like"`
	Humidity  float64            `bson:"humidity"`
	Pressure  float64            `bson:"pressure"`
	WindSpeed float64            `bson:"wind_speed"`
}
