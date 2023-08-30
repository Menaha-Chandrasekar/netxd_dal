package netxddalmodels

import "time"


type CustomerRequest struct {
	CustomerId         int32    `json:"customerid" bson:"customerid"`
	FirstName          string    `json:"firstname" bson:"firstname"`
	LastName           string    `json:"lastname" bson:"lastname"`
	BankId             int32    `json:"bankid" bson:"bankid"`
	Balance            int64    `json:"balance" bson:"balance"`
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" bson:"updated_at"`
	IsActive           bool      `json:"isactive" bson:"isactive"`
}


type  CustomerResponse struct {
	CustomerId         int32     `json:"customerid" bson:"customerid"`
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
}