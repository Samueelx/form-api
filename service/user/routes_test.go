package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Samueelx/form-api/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserSore{}
	handler := NewHandler(userStore)

	t.Run("Should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Testing",
			LastName:  "Mic",
			Email:     "invalid",
			Town:      "Kiambu",
			Age:       fmt.Sprint(19),
			Gender:    "Male",
			Password:  "asd",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		/**Make the request against the handler*/
		rr := httptest.NewRecorder() //Response recorder
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.HandleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected error code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("Should correctly register the user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Testing",
			LastName:  "Mic",
			Email:     "valid@mail.com",
			Town:      "Kiambu",
			Age:       fmt.Sprint(19),
			Gender:    "Male",
			Password:  "asd",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		/**Make the request against the handler*/
		rr := httptest.NewRecorder() //Response recorder
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.HandleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected error code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserSore struct{}

/**Ensure that mockUserStore implements the UserStore interface*/
func (m *mockUserSore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserSore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserSore) CreateUser(types.User) error {
	return nil
}
