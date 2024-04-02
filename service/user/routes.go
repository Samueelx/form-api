package user

import (
	"fmt"
	"net/http"

	"github.com/Samueelx/form-api/service/auth"
	"github.com/Samueelx/form-api/types"
	"github.com/Samueelx/form-api/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.HandleRegister).Methods("POST")
	router.HandleFunc("/login", h.HandleLogin).Methods("POST")
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	/**Get JSON Payload*/
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	/**Check if User exists*/
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		/**User exists*/
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with emal: %s already exists", payload.Email))
		return
	}

	/**Hash the password*/
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	/**Create the new user*/
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
		Age:       payload.Age,
		Town:      payload.Town,
		Gender:    payload.Gender,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
