package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID 						primitive.ObjectID			`bson:"_id"`
	FirstName 				*string						`json:"first_name" bson:"first_name" validate:"required, min=2, max=30"`
	LastName 				*string						`json:"last_name" validate:"required, min=2, max=30"`
	Email 					*string						`json:"email" validate:"email, required"`
	Password 				*string						`json:"password" validate:"required"`
	Token 					*string						`json:"token"`
	RefreshToken 			*string						`json:"refresh_token"`
	CreatedAt 				time.Time					`json:"created_at"`
	UpdatedAt 				time.Time					`json:"updated_at"`
	UserId 					*string						`json:"user_id"`
	UserCart 				[]ProductUser				`json:"user_cart" bson:"user_cart"`
	AddressDetails 			[]Address					`json:"address_details" bson:"address_details"`
	OrderStatus 			[]Order						`json:"order_status" bson:"order_status"`
}

type Product struct {
	ProductId 				primitive.ObjectID			`bson:"_id"`
	ProductName 			*string						`json:"product_name" bson:"product_name"`
	Price 					*uint64						`json:"price" bson:"price"`
	Rating 					*uint8						`json:"rating" bson:"rating"`
	Image 					*string						`json:"image" bson:"image"`
}

type ProductUser struct {
	ProductId 				primitive.ObjectID			`bson:"_id"`
	Name 					*string						`json:"name"`
	Price 					*uint64						`json:"price"`
	Rating 					*uint8						`json:"rating"`
	Image 					*string						`json:"image"`
}

type Address struct {
	AddressId 				primitive.ObjectID			`bson:"_id"`
	House 					*string						`json:"house" bson:"house"`
	Street 					*string						`json:"street" bson:"street"`
	City 					*string						`json:"city" bson:"city"`
	PostalCode 				*string						`json:"postal_code" bson:"postal_code"`
}

type Order struct {
	OrderId 				primitive.ObjectID			`bson:"_id"`
	OrderCart 				[]Product					`json:"order_cart" bson:"order_cart"`
	OrderedAt 				time.Time					`json:"ordered_at" bson:"ordered_at"`
	Price 					*uint64						`json:"price" bson:"price"`
	Discount 				*int						`json:"discount" bson:"discount"`
	PaymentMethod 			Payment						`json:"payment_method" bson:"payment_method"`

}

type Payment struct {
	Digital 				bool 						`json:"digital" bson:"digital"`// Online Payment
	COD 					bool						`json:"cod" bson:"cod"` // Cash On Delivery
}