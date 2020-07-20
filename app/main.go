package main

import (
	"othala/app/deliveries"
	_ "othala/app/products"
	_ "othala/app/products/adapters"
)

func main() {
	deliveries.Run()
}
