package handler

import (
	"fmt"
	"github.com/KhanbalaRashidov/Go-Ecommerce/helper"
	"github.com/KhanbalaRashidov/Go-Ecommerce/models/dto"
	"github.com/KhanbalaRashidov/Go-Ecommerce/service/cart"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func (h *Handler) CheckoutHandle(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())

	var cartDto dto.CartCheckoutDto
	if err := helper.ParseJSON(r, &cartDto); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := helper.Validate.Struct(cartDto); err != nil {
		errors := err.(validator.ValidationErrors)
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	productIds, err := cart.GetCartItemsIDs(cartDto.Items)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// get products
	products, err := h.productStore.GetProductsByID(productIds)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(products, cart.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"total_price": totalPrice,
		"order_id":    orderID,
	})
}
