package handler

import (
	"fmt"
	"github.com/KhanbalaRashidov/Go-Ecommerce/configs"
	"github.com/KhanbalaRashidov/Go-Ecommerce/helper"
	"github.com/KhanbalaRashidov/Go-Ecommerce/models"
	"github.com/KhanbalaRashidov/Go-Ecommerce/models/dto"
	"github.com/KhanbalaRashidov/Go-Ecommerce/service/auth"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) LoginHandle(w http.ResponseWriter, r *http.Request) {
	var user dto.LoginUserDto
	if err := helper.ParseJSON(r, &user); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := helper.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	u, err := h.userStore.GetUserByEmail(user.Email)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !helper.ComparePasswords(u.Password, []byte(user.Password)) {
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	secret := []byte(configs.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.Id)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) RegisterHandle(w http.ResponseWriter, r *http.Request) {
	var user dto.RegisterUserDto
	if err := helper.ParseJSON(r, &user); err != nil {
		helper.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := helper.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// check if user exists
	_, err := h.userStore.GetUserByEmail(user.Email)
	if err == nil {
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	// hash password
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.userStore.CreateUser(models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	helper.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) GetUserHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["userID"]
	if !ok {
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user ID"))
		return
	}

	userID, err := strconv.Atoi(str)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID"))
		return
	}

	user, err := h.userStore.GetUserByID(userID)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	helper.WriteJSON(w, http.StatusOK, user)
}
