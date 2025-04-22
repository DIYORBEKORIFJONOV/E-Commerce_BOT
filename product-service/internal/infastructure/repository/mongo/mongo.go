package mongodb

import (
	"context"
	"fmt"
	"log"
	"productservice/internal/config"
	productentity "productservice/internal/entity/product"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	mongoClient *mongo.Client
	db          *mongo.Database
	collectionCategory  *mongo.Collection
	collecectionProductName *mongo.Collection
	collectionModel *mongo.Collection
}


func NewMongoDB(cfg *config.Config) (*MongoDB, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", cfg.DB.Host, cfg.DB.Port)
	clientOptions := options.Client().ApplyURI(uri).
		SetServerSelectionTimeout(5 * time.Second)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("mongo.Connect error: %w", err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("mongo.Ping error: %w", err)
	}

	db := client.Database(cfg.DB.Name)

	return &MongoDB{
		mongoClient:            client,
		db:                     db,
		collectionCategory:     db.Collection(cfg.DB.CollectionCategory),
		collecectionProductName: db.Collection(cfg.DB.CollectionProductName),
		collectionModel:        db.Collection(cfg.DB.CollectionModel),
	}, nil
}




func (m *MongoDB) AddProduct(ctx context.Context, product *productentity.ProductMain) error {
	_, err := m.collecectionProductName.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB)AddModelToProduct(ctx context.Context, product *productentity.Product) (err error) {
	product.IsDeleted = false
	_,err = m.collectionModel.InsertOne(ctx,product)
	if err != nil {
		return err
	}

	return nil
}


func (m *MongoDB) GetAllProduct(ctx context.Context, req *productentity.GetProductsReq) ([]*productentity.Product, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"isdeleted": bson.M{"$exists": false}},
			{"isdeleted": false},
		},
	}

	if req.Field != "" && req.Value != "" {
		filter[req.Field] = req.Value
	}

	findOptions := options.Find()
	if req.Page > 0 && req.Limit > 0 {
		skip := (req.Page - 1) * req.Limit
		findOptions.SetSkip(skip)
		findOptions.SetLimit(req.Limit)
	}

	cursor, err := m.collectionModel.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("mongo GetAllProduct: %w", err)
	}
	defer cursor.Close(ctx)

	var products []*productentity.Product
	for cursor.Next(ctx) {
		var product productentity.Product
		if err := cursor.Decode(&product); err != nil {
			continue
		}
		products = append(products, &product)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("mongo cursor error: %w", err)
	}
	log.Println(products,filter)

	return products, nil
}



func (m *MongoDB) UpdateProduct(ctx context.Context, req *productentity.UpdateProductReq) (*productentity.Product, error) {

	update := bson.M{}
	if req.Description != "" {
		update["description"] = req.Description
	}
	if req.Colour != "" {
		update["colour"] = req.Colour
	}
	if req.Size != 0 {
		update["size"] = req.Size
	}
	if req.Price != 0 {
		update["price"] = req.Price
	}
	if req.Quantity != 0 {
		update["quantity"] = req.Quantity
	}
	if req.PhotoURL != "" {
		update["photourl"] = req.PhotoURL
	}

	if len(update) == 0 {
		return nil, fmt.Errorf("nothing to update")
	}

	filter := bson.M{"productid": req.ProductID}
	updateDoc := bson.M{"$set": update}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updated productentity.Product
	err := m.collectionModel.FindOneAndUpdate(ctx, filter, updateDoc, opts).Decode(&updated)
	if err != nil {
		return nil, fmt.Errorf("mongo UpdateProduct: %w", err)
	}

	return &updated, nil
}


func (m *MongoDB) DeleteProduct(ctx context.Context, productId string, hardDelete bool) error {
	filter := bson.M{"productid": productId}

	if hardDelete {
		_, err := m.collectionModel.DeleteOne(ctx, filter)
		if err != nil {
			return fmt.Errorf("mongo hard delete: %w", err)
		}
		fmt.Println("Товар полностью удалён из базы данных.")
	} else {
		update := bson.M{"$set": bson.M{"isdeleted": true}}
		_, err := m.collectionModel.UpdateOne(ctx, filter, update)
		if err != nil {
			return fmt.Errorf("mongo soft delete: %w", err)
		}
		fmt.Println("Товар помечен как удалённый (isdeleted = true).")
	}

	return nil
}


func (m *MongoDB) CreateCategory(ctx context.Context, nameCategory string) error {
	doc := bson.M{"category": nameCategory}
	_, err := m.collectionCategory.InsertOne(ctx, doc)
	if err != nil {
		return fmt.Errorf("mongo CreateCategory: %w", err)
	}
	return nil
}


func (m *MongoDB) GetAllCategory(ctx context.Context) ([]*productentity.CreateCategoryReq, error) {
	cursor, err := m.collectionCategory.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("mongo GetAllCategory: %w", err)
	}
	defer cursor.Close(ctx)

	var categories []*productentity.CreateCategoryReq
	for cursor.Next(ctx) {
		var category productentity.CreateCategoryReq
		if err := cursor.Decode(&category); err != nil {
			continue
		}
		categories = append(categories, &category)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("mongo cursor error: %w", err)
	}

	return categories, nil
}


func (m *MongoDB) UpdateCategory(ctx context.Context, category, newCategory string) error {
	filter := bson.M{"category_name": category}
	update := bson.M{"$set": bson.M{"category_name": newCategory}}

	result, err := m.collectionCategory.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("mongo UpdateCategory: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("category '%s' not found", category)
	}

	return nil
}


func (m *MongoDB) DeleteCategory(ctx context.Context, nameCategory string) error {
	filter := bson.M{"category_name": nameCategory}

	result, err := m.collectionCategory.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("mongo DeleteCategory: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("category '%s' not found", nameCategory)
	}

	return nil
}



func (m *MongoDB) Update(ctx context.Context, name, newName string) error {
	filter := bson.M{"name": name}
	update := bson.M{"$set": bson.M{"name": newName}}

	result, err := m.collecectionProductName.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("mongo Update: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("product with name '%s' not found", name)
	}

	return nil
}


func (m *MongoDB) GetMainProduct(ctx context.Context,field,value string)([]*productentity.ProductMain, error) {
	filter := bson.M{field: value}

	cursor, err := m.collecectionProductName.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("mongo Find GetMainProduct: %w", err)
	}
	defer cursor.Close(ctx)

	var products []*productentity.ProductMain
	if err := cursor.All(ctx, &products); err != nil {
		return nil, fmt.Errorf("decoding products: %w", err)
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("products not found")
	}

	return products, nil
}
