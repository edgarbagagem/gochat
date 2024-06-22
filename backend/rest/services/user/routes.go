package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/edgarbagagem/gochat/config"
	"github.com/edgarbagagem/gochat/services/auth"
	"github.com/edgarbagagem/gochat/types"
	"github.com/edgarbagagem/gochat/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods(http.MethodPost)
	router.HandleFunc("/register", h.handleRegister).Methods(http.MethodPost)
	router.HandleFunc("/users/{user}", h.handleGetUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{userID}", h.handleUpdateUser).Methods(http.MethodPut)
}

func (h *Handler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	// Handle editing user by ID logic here, using userID
	var payload types.UpdateUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	id, err := strconv.Atoi(userID)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID %d, error: %s", id, err.Error()))
		return
	}

	user, err := h.store.GetUserById(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user not found,  %s", err.Error()))
		return
	}

	//Update User
	if payload.Photo.Valid {
		user.Photo = payload.Photo
	}

	err = h.store.UpdateUser(user)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "")
}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]

	if _, err := strconv.Atoi(user); err == nil {
		// Handle get user by ID
		h.handleGetUserById(w, r, user)
	} else {
		// Handle get user by username
		h.handleGetUserByUsername(w, r, user)
	}
}

func (h *Handler) handleGetUserById(w http.ResponseWriter, _ *http.Request, id string) {

	userID, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID"))
		return
	}

	user, err := h.store.GetUserById(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) handleGetUserByUsername(w http.ResponseWriter, _ *http.Request, username string) {

	user, err := h.store.GetUserByUsername(username)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.store.GetUserByUsername(payload.Username)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user not found,  %s", err.Error()))
		return
	}

	if !auth.ComparePasswords(user.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token, "username": user.Username, "photoURL": user.Photo.String})

}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err := h.store.GetUserByUsername(payload.Username)

	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with username %s already exists", payload.Username))
		return
	}

	// hash password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		Username: payload.Username,
		Password: hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
