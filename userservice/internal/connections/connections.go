package connections

import (
	"account/internal/config"
	"account/internal/database"
	interfaceaccount "account/internal/interface"
	"account/internal/service"
	"account/internal/service/adjust"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mongo() interfaceaccount.InterfaceAccount {
	cfg := config.Configuration()
	opt, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal(err)
	}
	if err := opt.Ping(context.Background(), options.Client().ReadPreference); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	data := opt.Database(cfg.Database.DBname).Collection(cfg.Database.CollectionOrders)
	return &database.Database{M: data, C: ctx}
}

func NewAdjust() *adjust.AdjustService {
	return &adjust.AdjustService{
		I: Mongo(),
	}
}

func NewService() *service.Service {
	return &service.Service{
		A: NewAdjust(),
	}
}
