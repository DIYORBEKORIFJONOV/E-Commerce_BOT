package connections

import (
	"context"
	"log"

	"github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/config"
	"github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/database/mongodb"
	interface17 "github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/interface"
	"github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/service"
	"github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/service/adjust"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mongo() *mongodb.Mongo {
	cfg := config.Configuration()
	opt, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	if err := opt.Ping(context.Background(), options.Client().ReadPreference); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	data := opt.Database(cfg.Database.DBname).Collection(cfg.Database.CollectionOrders)
	d2 := opt.Database(cfg.Database.DBname).Collection(cfg.Database.CollectionCart)
	return &mongodb.Mongo{D: data, B: d2, C: ctx}
}

func NewMongo() interface17.Order {
	m1 := Mongo()
	ctx := context.Background()

	return &mongodb.Mongo{D: m1.D, B: m1.B, C: ctx}
}

func NewCentralized() *adjust.Centralized {
	m := NewMongo()
	return &adjust.Centralized{M: m}
}

func NewAdjust() interface17.OrderService {
	c := NewCentralized()
	ctx := context.Background()
	return &adjust.Adjust{CA: c, C: ctx}
}

func NewService() *service.Service {
	a := NewAdjust()
	return &service.Service{A: a}
}
