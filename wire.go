//+build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/alanwade2001/spa-customer-api/repositories"
	"github.com/alanwade2001/spa-customer-api/routers"
	"github.com/alanwade2001/spa-customer-api/services"
	"github.com/alanwade2001/spa-customer-api/types"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitialiseServerAPI() types.ServerAPI {
	wire.Build(
		gin.Default,
		repositories.NewMongoService,
		services.NewConfigService,
		services.NewService,
		routers.NewCustomerRouter,
		routers.NewCustomerSearchRouter,
		routers.NewRegisterService,
		NewServer,
	)

	return &Server{}
}
