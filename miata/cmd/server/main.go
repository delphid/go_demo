package main

import (
	"go.uber.org/fx"
	"miata/model"

	"miata/api"
	config "miata/init"
)

func Start(_ *api.Server) {
}

func main() {
	app := fx.New(
		config.Module,
		model.Module,
		fx.Provide(api.NewServer),
		fx.Invoke(
			Start,
		),
	)
	app.Run()
}
