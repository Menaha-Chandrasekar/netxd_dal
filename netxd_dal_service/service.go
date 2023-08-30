
package netxddalservices

import (
	"context"
	"log"
	netxddalinterfaces "module/netxd_dal/netxd_dal_interface"
	
	netxddalmodels "module/netxd_dal/netxd_dal_models"
	"time"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

type CustomerService struct{
CustomerCollection *mongo.Collection
ctx context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) netxddalinterfaces.ICustomer {
	return &CustomerService{collection, ctx}
}

func (c * CustomerService)CreateCustomer(detail * netxddalmodels.CustomerRequest)(*netxddalmodels.CustomerResponse,error){
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"customerid" : 1}, // 1 for ascending, -1 for descending
		Options: options.Index().SetUnique(true),
	}
	_, err := c.CustomerCollection.Indexes().CreateOne(c.ctx, indexModel)
	if err != nil {
		log.Fatal(err)
	}
	detail.IsActive = true
	detail.CreatedAt = time.Now()
	detail.UpdatedAt = detail.CreatedAt
	fmt.Println(detail.CustomerId)

	res,err:=c.CustomerCollection.InsertOne(c.ctx,&detail)
	if err!=nil{
		return nil,err
	}
	var newUser *netxddalmodels.CustomerResponse
	query:=bson.M{"_id":res.InsertedID}

	err = c.CustomerCollection.FindOne(c.ctx,query).Decode(&newUser)
	if err != nil{
		return nil,err
	}
	return newUser,nil
}
