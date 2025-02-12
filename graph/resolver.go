package graph

import (
	"fmt"
	"sync"

	"github.com/graphQL_easy/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductAdded chan *model.Product
}

var productStore = make(map[string]*model.Product)
var mu sync.Mutex
var productIDCounter = 1

func nextID() string {
	mu.Lock()
	defer mu.Unlock()
	id := fmt.Sprintf("%d", productIDCounter)
	productIDCounter++
	return id
}
