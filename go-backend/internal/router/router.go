package router

import (
	"go-backend/internal/handler"
	"go-backend/internal/middleware"
	"go-backend/internal/repository"
	"go-backend/internal/service"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Initialize repositories
	userRepo := repository.NewUserRepository()
	contactRepo := repository.NewContactRepository()
	addressRepo := repository.NewAddressRepository()

	// Initialize services
	userService := service.NewUserService(userRepo)
	contactService := service.NewContactService(contactRepo)
	addressService := service.NewAddressService(addressRepo, contactRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	contactHandler := handler.NewContactHandler(contactService)
	addressHandler := handler.NewAddressHandler(addressService)
	healthHandler := handler.NewHealthHandler()

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware()

	// Public routes
	r.HandleFunc("/api/users", userHandler.Register).Methods("POST")
	r.HandleFunc("/api/users/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/ping", healthHandler.Ping).Methods("GET")

	// Protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(authMiddleware.RequireAuth)

	// User routes
	protected.HandleFunc("/users/current", userHandler.GetCurrent).Methods("GET")
	protected.HandleFunc("/users/current", userHandler.Update).Methods("PATCH")
	protected.HandleFunc("/users/logout", userHandler.Logout).Methods("DELETE")

	// Contact routes
	protected.HandleFunc("/contacts", contactHandler.Create).Methods("POST")
	protected.HandleFunc("/contacts/{contactId:[0-9]+}", contactHandler.GetByID).Methods("GET")
	protected.HandleFunc("/contacts/{contactId:[0-9]+}", contactHandler.Update).Methods("PUT")
	protected.HandleFunc("/contacts/{contactId:[0-9]+}", contactHandler.Delete).Methods("DELETE")
	protected.HandleFunc("/contacts", contactHandler.Search).Methods("GET")

	// Address routes
	protected.HandleFunc("/contacts/{contactId:[0-9]+}/addresses", addressHandler.Create).Methods("POST")
	protected.HandleFunc("/contacts/{contactId:[0-9]+}/addresses/{addressId:[0-9]+}", addressHandler.GetByID).Methods("GET")
	protected.HandleFunc("/contacts/{contactId:[0-9]+}/addresses/{addressId:[0-9]+}", addressHandler.Update).Methods("PUT")
	protected.HandleFunc("/contacts/{contactId:[0-9]+}/addresses/{addressId:[0-9]+}", addressHandler.Delete).Methods("DELETE")
	protected.HandleFunc("/contacts/{contactId:[0-9]+}/addresses", addressHandler.GetByContactID).Methods("GET")

	return r
}