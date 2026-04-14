package main

import (
	"gin-test/router"
)

func main() {

	r := router.Setup()

	r.Run(":6063")
}
