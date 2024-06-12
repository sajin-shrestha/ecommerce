package cart

import (
	"fmt"

	"github.com/sajin-shrestha/ecommerce/types"
)

func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	productIDs := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductID)
		}

		productIDs[i] = item.ProductID
	}

	return productIDs, nil
}

func (h *Handler) createOrder(ps []types.Product, item []types.CartItem, userID int) (int, float64, error) {
	// we need to loop thorugh products many times so we create a product-map for ease of use and optimization
	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}

	// check if all products are actually in stock
	// calculate the total price
	// reduce the quantity of products in our database
	// create the order
	// create order items
}
