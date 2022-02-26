package main

import (
	"othala/app/deliveries"
	_ "othala/app/products"
	_ "othala/app/products/adapters"
	_ "othala/app/products/repositories"
	_ "othala/app/products/repositories/implementations"
	_ "othala/app/products/usecases"
	_ "othala/app/products/usecases/implementations"
)

func main() {
	deliveries.Run()
}
