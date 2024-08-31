package handler

import (
	"github.com/KhanbalaRashidov/Go-Ecommerce/models"
	"github.com/KhanbalaRashidov/Go-Ecommerce/service/cart"
)

func (h *Handler) CreateOrder(products []models.Product, cartItems []models.CartCheckoutItem, userID int) (int, float64, error) {
	// create a map of products for easier access
	productsMap := make(map[int]models.Product)
	for _, product := range products {
		productsMap[product.Id] = product
	}

	// check if all products are available
	if err := cart.CheckIfCartIsInStock(cartItems, productsMap); err != nil {
		return 0, 0, err
	}

	// calculate total price
	totalPrice := cart.CalculateTotalPrice(cartItems, productsMap)

	// reduce the quantity of products in the store
	for _, item := range cartItems {
		product := productsMap[item.ProductID]
		product.Quantity -= item.Quantity
		h.productStore.UpdateProduct(product)
	}

	// create order record
	orderID, err := h.orderStore.CreateOrder(models.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "some address", // could fetch address from a user addresses table
	})
	if err != nil {
		return 0, 0, err
	}

	// create order the items records
	for _, item := range cartItems {
		h.orderStore.CreateOrderItem(models.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productsMap[item.ProductID].Price,
		})
	}

	return orderID, totalPrice, nil
}
