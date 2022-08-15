package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type Application struct {
	productCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(productCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		productCollection: productCollection,
		userCollection: userCollection,
	}
}

func (app *Application)AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")

		if productQueryID  == ""{
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}
		userQueryId := c.Query("userID")
		if userQueryId == ""{

		}
	}
}

func RemoveItem() gin.HandlerFunc {
	
}

func GetItemFromCart() gin.HandlerFunc {
	
}

func BuyFromCart () gin.HandlerFunc {
	
}

func InstantBuy() gin.HandlerFunc {

}