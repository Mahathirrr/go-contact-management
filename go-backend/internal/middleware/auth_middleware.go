package middleware

import (
	"database/sql"
	"encoding/json"
	"go-backend/internal/database"
	"go-backend/internal/models"
	"net/http"
)

type AuthMiddleware struct {
	db *sql.DB
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		db: database.DB,
	}
}

func (m *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Errors: "Unauthorized",
			})
			return
		}

		user := m.findUserByToken(token)
		if user == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Errors: "Unauthorized",
			})
			return
		}

		// Add user to context
		r = r.WithContext(r.Context())
		r.Header.Set("X-User-Username", user.Username)

		next.ServeHTTP(w, r)
	})
}

func (m *AuthMiddleware) findUserByToken(token string) *models.User {
	query := `SELECT username, password, name, token FROM users WHERE token = ?`
	row := m.db.QueryRow(query, token)

	var user models.User
	err := row.Scan(&user.Username, &user.Password, &user.Name, &user.Token)
	if err != nil {
		return nil
	}

	return &user
}