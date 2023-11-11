package main

import (
	"github.com/Michael-Sjogren/gohtmx/internal/model"
	"github.com/Michael-Sjogren/gohtmx/internal/routes"
)

func main() {
	model.Setup()
	routes.SetupServerAndRun()
}
