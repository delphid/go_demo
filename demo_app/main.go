package main

import (
	"go.uber.org/fx"

	"demo_app/api"
)

func Start(_ *api.Server) {
}

func main() {
	app := fx.New(
		fx.Provide(api.NewServer),
		fx.Invoke(
			Start,
		),
	)
	app.Run()
}
