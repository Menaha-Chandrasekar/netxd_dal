package netxddalinterface

import netxddalmodels "github.com/Menaha-Chandrasekar/netxd_dal/netxd_dal_models"


type ICustomer interface{
    CreateCustomer(customer *netxddalmodels.CustomerRequest)(*netxddalmodels.CustomerResponse,error)
}