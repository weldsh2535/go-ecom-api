package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/weldsh2535/go-rest-api/types"
)

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("Should succeed if the user payload is valid", func(t *testing.T) {
		// A valid payload that should succeed
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "validemail@example.com",
			Password:  "securepassword",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})

	t.Run("Should fail if the user payload is invalid", func(t *testing.T) {
		// An invalid payload with an improperly formatted email
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "invalidemail",
			Password:  "password",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

// mockUserStore is a mock implementation of the UserStore interface
type mockUserStore struct {
}

// Simulate no user existing for the email lookup
func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	if email == "validemail@example.com" {
		return &types.User{}, nil // Simulate existing user
	}
	return nil, errors.New("user not found")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

// Simulate successful user creation by returning a mock user ID
func (m *mockUserStore) CreateUser(user types.User) (int64, error) {
	return 1, nil
}
