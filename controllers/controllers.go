package controllers

import (
	"context"
	"fmt"
	"github.com/codagott/ecommerce/database"
	"github.com/codagott/ecommerce/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var ProductionCollection *mongo.Collection = database.ProductData(database.Client, "Products")
var Validate = validator.New()

func HashPassword(password string){

}

func VerifyPassword(userPassword string, givenPassword string) (bool, string){

}

func SignUp() gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx , cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := Validate.Struct(user)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : validationErr})
			return
		}
		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error" : "user already exists"})
		}
		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone" : user.Phone})

		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error" : "phone number in use already"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format())
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format())
		user.ID = primitive.NewObjectID()
		user.UserId = user.ID.Hex()
		token, refreshToken, _ := generate.TokenGenerator(*user.Email, *user.FirstName, *user.LastName, user.UserId)
		user.Token = &token
		user.RefreshToken = &refreshToken
		user.UserCart = make([]models.ProductUser, 0)
		user.AddressDetails = make([]models.Address, 0)
		user.OrderStatus = make([]models.Order, 0)

		_, inserter := UserCollection.InsertOne(ctx, user)

		if inserter != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "the user did not get created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusCreated, "Account created successfully.")

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		var foundUser models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		err := UserCollection.FindOne(ctx, bson.M{"email" : user.Email}).Decode(&foundUser)
		defer cancel()

		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "password or email not correct"})
			return
		}

		PasswordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)

		defer cancel()

		if !PasswordIsValid{
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}
		token, refreshToken, _ := generate.TokenGenerator(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, foundUser.UserId)
		defer cancel()

		generate.UpdateAllToken(token, refreshToken, foundUser.UserId)

		c.JSON(http.StatusFound, foundUser)
	}
}

func ProductViewerAdmin() gin.HandlerFunc{

}

func SearchProduct() gin.HandlerFunc{

}

func SearchProductByQuery() gin.HandlerFunc{

}
