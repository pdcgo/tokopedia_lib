package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TopedAggregateCategory interface {
	IterCategory(namespace string, handler func(tokopediaID int, count int, name []string) error) error
}

type ProductAggregateIpml struct {
	Collection *mongo.Collection
}

type aggCategItem struct {
	ID       int      `bson:"_id" json:"_id"`
	PriceMin int      `bson:"price_min" json:"price_min"`
	PriceMax int      `bson:"price_max" json:"price_max"`
	Count    int      `bson:"count" json:"count"`
	Name     []string `bson:"name" json:"name"`
}

func (agg *ProductAggregateIpml) IterCategory(namespace string, handler func(tokopediaID int, count int, name []string) error) error {
	hasil := []*aggCategItem{}
	query := bson.M{
		"marketplace": "tokopedia",
		"namespace":   namespace,
	}
	matchStage := bson.D{{Key: "$match", Value: query}}

	aggregateStage := bson.D{{
		Key: "$group",
		Value: bson.M{
			"_id":       "$category_id",
			"price_min": bson.M{"$min": "$price_after_discount"},
			"price_max": bson.M{"$max": "$price_after_discount"},
			"count":     bson.M{"$sum": 1},
			"name":      bson.M{"$first": "$categories.display_name"},
		},
	}}

	pipeline := mongo.Pipeline{
		matchStage,
		aggregateStage,
	}

	cursor, err := agg.Collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return err
	}

	cursor.All(context.TODO(), &hasil)

	if err != nil {
		return err
	}

	for _, data := range hasil {
		handler(data.ID, data.Count, data.Name)
	}

	return err
}
