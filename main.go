package main

import (
	"github.com/ade-iqbal/fga-assignment-3/router"
)

func main() {
	router.StartApp().Run(":8080")
}