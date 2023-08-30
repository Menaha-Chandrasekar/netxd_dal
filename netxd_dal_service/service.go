// package netxddalservice

// import (
// 	"context"

// 	// "log"
// 	netxddalinterface "module/netxd_dal/netxd_dal_interface"
// 	netxddalmodels "module/netxd_dal/netxd_dal_models"

// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
	
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type CustomerService struct {
// 	CustomerCollection *mongo.Collection
// 	ctx               context.Context
// }

// func InitCustomerService(collection *mongo.Collection, ctx context.Context) netxddalinterface.ICustomer {
// 	return &CustomerService{ collection , ctx}
// }

// func (c *CustomerService) CreateCustomer(user *netxddalmodels.CustomerRequest) (*netxddalmodels.CustomerResponse, error) {
	
// 	// indexModel := mongo.IndexModel{
// 	// 	Keys: bson.M{"customer_id": 1}, // 1 for ascending, -1 for descending
// 	// 	Options: options.Index().SetUnique(true),
// 	// 	}
// 	// 	_, err := c.CustomerCollection.Indexes().CreateOne(c.ctx, indexModel)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 		date := time.Now()
// 		user.CreatedAt = date.Format("2006-01-02 12.50.00.000000000")
		
// 		res,err := c.CustomerCollection.InsertOne(c.ctx, &user)
// 		if err!=nil{
// 		return nil,err
// 		}
		
// 		var newUser *netxddalmodels.CustomerResponse
// 		query := bson.M{"_id": res.InsertedID}
	
// 		err = c.CustomerCollection.FindOne(c.ctx, query).Decode(&newUser)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return newUser, nil
// 	}

package netxddalservices

import (
"context"
netxddalinterfaces "module/netxd_dal/netxd_dal_interface"
netxddal   "module/netxd_dal/netxd_dal_models"
"time"

"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/mongo"
//netxddal "dal/netxd_dal/netxd_dal_models"
)

type CustomerService struct{
CustomerCollection *mongo.Collection
ctx context.Context
}

func InitialiseCustomerService(collection *mongo.Collection,ctx context.Context)netxddalinterfaces.ICustomer{
return &CustomerService{ collection,ctx}
}

func(c * CustomerService) CreateCustomer(user *netxddal.CustomerRequest)(*netxddal.CustomerResponse,error){

	date := time.Now()
	user.CreatedAt = date.Format("2006-01-02 12.50.00.000000000")
user.UpdatedAt = user.CreatedAt
user.IsActive = false
res, err := c.CustomerCollection.InsertOne(c.ctx, user)
if err != nil {
return nil,err
}

var newUser *netxddal.CustomerResponse
query := bson.M{"_id": res.InsertedID}

err = c.CustomerCollection.FindOne(c.ctx, query).Decode(&newUser)
if err != nil {
return nil, err
}
return newUser, nil
}