package repositories

import (
	"othala/app/products/repositories"
	"time"
)

type ProductLists struct {
	ID    int
	Items []Item
}

type Item struct {
	Product   repositories.Product
	Value     float64
	AddedDate time.Time
}
