package main

import (
	"github.com/dyhalmeida/golang-students-api/api"
)

func main() {
	serviceApi := api.New()
	serviceApi.Start(":3333")
}
