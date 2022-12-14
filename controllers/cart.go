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

func (app *Application)RemoveItem() gin.HandlerFunc {
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

		err = database.RemoveCartItem(ctx, app.productCollection, app.userCollection, productId, userQueryId)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, "Successfully removed item from cart")
	}
}

func GetItemFromCart() gin.HandlerFunc {
	
}

func (app *Application) BuyFromCart () gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")

		if userQueryID == ""{
			log.Panicln("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		defer cancel()
		
		err := database.BuyItemFromCart(ctx, app.userCollection, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(http.StatusOK, "successfully place the order")
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
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

		err = database.InstantBuyer(ctx, app.productCollection, app.userCollection, productId, userQueryId)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(http.StatusOK, "successfully placed order")
	}

}