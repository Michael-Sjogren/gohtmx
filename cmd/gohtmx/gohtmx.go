package main

import (
	_ "embed"

	"gohtmx/internal/model"
	"gohtmx/internal/routes"
)

func main() {
	model.Setup()
	routes.SetupServerAndRun()
}
