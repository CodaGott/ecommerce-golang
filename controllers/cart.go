package controllers

import (
	"context"
	"errors"
	"github.com/codagott/ecommerce/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
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
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id empty"))
			return
		}

		productId, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 5 * time.Second)

		defer cancel()

		err = database.AddProductToCart(ctx, app.productCollection, app.userCollection, productId, userQueryId)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(http.StatusOK, "Successfully added to the cart")
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