package app

import (
	"gateway/pkg/config/global_vars_config"
	"gateway/pkg/config/service_address_config"
	"gateway/pkg/service/setup"
)

func Run() {
	setup.AppSetup()
	err := global_vars_config.Router.Run(*service_address_config.Address)
	if err != nil {
		panic(err)
	}
}
