package main

import (
	fxapp "github.com/krakatoa/go-project-start-template/internal"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fxapp.Module,
	).Run()
}
