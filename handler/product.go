package handler

import (
	"fmt"
	"github.com/KhanbalaRashidov/Go-Ecommerce/helper"
	"github.com/KhanbalaRashidov/Go-Ecommerce/models/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductsHandle(w http.ResponseWriter, r *http.Request) {
	products, err := h.productStore.GetProducts()
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	helper.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) GetProductHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["productId"]
	if !ok {
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing product ID"))
		return
	}

	productId, err := strconv.Atoi(str)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID"))
		return
	}

	product, err := h.productStore.GetProductByID(productId)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	helper.WriteJSON(w, http.StatusOK, product)
}

func (h *Handler) CreateProductHandle(w http.ResponseWriter, r *http.Request) {

	var product dto.CreateProductDto
	if err := helper.ParseJSON(r, &product); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := helper.Validate.Struct(product); err != nil {
		errors := err.(validator.ValidationErrors)
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.productStore.CreateProduct(product)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	helper.WriteJSON(w, http.StatusCreated, product)
}
