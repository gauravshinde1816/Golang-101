package controllers

import (
	"context"
	"crud_101/configs"
	"crud_101/models"
	"crud_101/responses"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = configs.GetCollection(configs.DB, "product")

func CreateProduct(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// null object
	var proudct models.Product
	err := c.Bind(&proudct)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "invalid body schema", Data: &echo.Map{"data": err}})
	}

	newUser := models.Product{
		Id:    primitive.NewObjectID(),
		Name:  proudct.Name,
		Price: proudct.Price,
	}

	result, err := productCollection.InsertOne(ctx, newUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "Insert Error", Data: &echo.Map{"data": err}})
	}

	return c.JSON(http.StatusCreated, responses.ProductResponse{Status: http.StatusCreated, Message: "Status ok", Data: &echo.Map{"data": result}})

}

func GetProducts(c echo.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var products []models.Product

	results, err := productCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "ok", Data: &echo.Map{"data": err}})
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var product models.Product
		err = results.Decode(&product)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err}})
		}
		products = append(products, product)
	}

	return c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "ok", Data: &echo.Map{"data": products}})
}

func GetProduct(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	ID := c.Param("productID")

	productID, _ := primitive.ObjectIDFromHex(ID)

	var product models.Product

	err := productCollection.FindOne(ctx, bson.M{"id": productID}).Decode(&product)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ProductResponse{Status: http.StatusInternalServerError, Message: "collection find error", Data: &echo.Map{"data": err}})
	}

	return c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "ok", Data: &echo.Map{"data": product}})

}
