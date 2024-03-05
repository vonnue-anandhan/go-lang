package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the structure of a user.
type User struct {
	ID             primitive.ObjectID `json:"id" bson:"id"`
	FirstName      *string            `json:"firstName" bson:"firstName" validate:"required, min=2, max=30"`
	LastName       *string            `json:"lastName" bson:"lastName" validate:"required, min=2, max=30"`
	Password       *string            `json:"password" bson:"password" validate:"required, min=30"`
	Email          *string            `json:"email" bson:"email" validate:"email, required"`
	Phone          *string            `json:"phone" bson:"phone" validate:"required"`
	Token          *string            `json:"token" bson:"token"`
	RefreshToken   *string            `json:"refreshToken" bson:"refreshToken"`
	CreatedAt      time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt" bson:"updatedAt"`
	Cart           Cart               `json:"cart" bson:"cart"`
	AddressDetails []AddressDetails   `json:"addressDetails" bson:"addressDetails"`
	Orders         []Order            `json:"orders" bson:"orders"`
}

// Cart represents the structure of a shopping cart.
type Cart struct {
	ID           primitive.ObjectID `json:"id" bson:"id"`
	Products     []Product          `json:"products" bson:"products"`
	UserID       primitive.ObjectID `json:"userId" bson:"userId"`
	IsCheckedOut bool               `json:"isCheckedOut" bson:"isCheckedOut"`
}

// Product represents a product in the shopping cart.
type Product struct {
	ID       primitive.ObjectID `json:"id" bson:"id"`
	Name     *string            `json:"name" bson:"name"`
	Price    *uint64            `json:"price" bson:"price"`
	Rating   *uint8             `json:"rating" bson:"rating"`
	Image    *string            `json:"image" bson:"image"`
	Quantity int                `json:"quantity" bson:"quantity"`
}

// AddressDetails represents an address.
type AddressDetails struct {
	House   *string `json:"house" bson:"house"`
	City    *string `json:"city" bson:"city"`
	Pincode *string `json:"pincode" bson:"pincode"`
}

// Order represents an order.
type Order struct {
	ID            primitive.ObjectID `json:"id" bson:"id"`
	CartID        primitive.ObjectID `json:"cartId" bson:"cartId"`
	OrderedAt     time.Time          `json:"orderedAt" bson:"orderedAt"`
	Price         int                `json:"price" bson:"price"`
	Discount      *int               `json:"discount" bson:"discount"`
	PaymentMethod PaymentMethod      `json:"paymentMethod" bson:"paymentMethod"`
	DeliveryAddr  AddressDetails     `json:"deliveryAddr" bson:"deliveryAddr"`
}

// PaymentMethod represents a payment method.
type PaymentMethod struct {
	IsDigital bool `json:"isDigital" bson:"isDigital"`
	IsCOD     bool `json:"isCod" bson:"isCod"`
}
