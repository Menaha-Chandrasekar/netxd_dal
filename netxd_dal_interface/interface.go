package netxddalinterface

import netxddalmodels "module/netxd_dal/netxd_dal_models"

type ICustomer interface{
    CreateCustomer(customer *netxddalmodels.CustomerRequest)(*netxddalmodels.CustomerResponse,error)
}