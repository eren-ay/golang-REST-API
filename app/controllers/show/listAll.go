package show

import (
	"context"

	"github.com/eren-ay/golang-REST-API/app/models"
	"github.com/eren-ay/golang-REST-API/pkg/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AllId(ctx *fiber.Ctx) error {
	showCollection := database.GetCollection(database.DB, "Movie")
	//filter := bson.D{{Key: "title", Value: "eren"}}
	//bson.D{} = no filter
	filter := bson.D{}
	array, err := getShows(showCollection, filter)
	if err != nil {
		return err
	}
	return ctx.Status(200).JSON(array)
}

func getShows(coll *mongo.Collection, filter bson.D) ([]models.Show, error) {
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []models.Show
	if err = cursor.All(context.TODO(), &results); err != nil {
		return results, err
	}
	return results, nil
}
