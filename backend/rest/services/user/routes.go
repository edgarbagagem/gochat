package user

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
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
	router.HandleFunc("/profile-photo/{user}", auth.WithJWTAuth(h.handleGetUserPhoto, h.store)).Methods(http.MethodGet)
	router.HandleFunc("/users/{userID}", auth.WithJWTAuth(h.handleUpdateUser, h.store)).Methods(http.MethodPut)
	router.HandleFunc("/upload-photo", auth.WithJWTAuth(h.handleUploadPhoto, h.store)).Methods(http.MethodPost)
}

func (h *Handler) handleUploadPhoto(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Could not parse multipart form", http.StatusBadRequest)
		return
	}

	// Get the user ID or username from the form
	username := r.FormValue("username")
	if username == "" {
		http.Error(w, "Missing username", http.StatusBadRequest)
		return
	}

	user, err := h.store.GetUserByUsername(username)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Get the file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Could not get uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Upload the file to Google Cloud Storage
	url, err := uploadFileToGCS(file, handler)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not upload file to GCS %s", err.Error()), http.StatusInternalServerError)
		return
	}

	user.Photo.String = url

	h.store.UpdateUser(user)

	utils.WriteJSON(w, http.StatusOK, url)
}

var bucketName = "gochat-images"

func uploadFileToGCS(file multipart.File, handler *multipart.FileHeader) (string, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create client: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	objectName := fmt.Sprintf("%d-%s", time.Now().Unix(), handler.Filename)
	wc := bucket.Object(objectName).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("failed to copy file to bucket: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("failed to close bucket writer: %v", err)
	}

	// // Make the object publicly accessible
	// if err := bucket.Object(objectName).ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
	// 	return "", fmt.Errorf("failed to set bucket object ACL: %v", err)
	// }

	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName), nil
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

func (h *Handler) handleGetUserPhoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]

	if _, err := strconv.Atoi(user); err == nil {
		// Handle get user by ID
		h.handleGetUserPhotoById(w, r, user)
	} else {
		// Handle get user by username
		h.handleGetUserPhotoByUsername(w, r, user)
	}
}

func (h *Handler) handleGetUserPhotoById(w http.ResponseWriter, _ *http.Request, id string) {

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

	utils.WriteJSON(w, http.StatusOK, map[string]string{"photo": user.Photo.String})
}

func (h *Handler) handleGetUserPhotoByUsername(w http.ResponseWriter, _ *http.Request, username string) {

	user, err := h.store.GetUserByUsername(username)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"photo": user.Photo.String})
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

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})

}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate username
	if !isValidUsername(payload.Username) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("username must be at least 3 characters long and contain at least one letter"))
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

func isValidUsername(username string) bool {
	// Check if the username is at least 3 characters long and contains at least one letter
	if len(username) < 3 {
		return false
	}
	match, _ := regexp.MatchString(`[a-zA-Z]`, username)
	return match
}
