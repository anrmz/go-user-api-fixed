package user

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Svc *service
}

func (h Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	// Check method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var u User

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}

	// ✅ Validation
	if u.Name == "" || u.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{
	"error": "username and password are required",
}

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusBadRequest)
json.NewEncoder(w).Encode(response)
return
	}

	if len(u.Password) < 6 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("password must be at least 6 characters"))
		return
	}

	// Call service
	message, err := h.Svc.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{
	"error": "Failed to add user",
}

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusInternalServerError)
json.NewEncoder(w).Encode(response)
return	}

	// Success response
	response := map[string]string{
	"id":     message,
	"status": "success",
}

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(response)
}