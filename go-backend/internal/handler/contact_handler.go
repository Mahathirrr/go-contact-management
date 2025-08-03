package handler

import (
	"encoding/json"
	"go-backend/internal/models"
	"go-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ContactHandler struct {
	contactService service.ContactService
}

func NewContactHandler(contactService service.ContactService) *ContactHandler {
	return &ContactHandler{
		contactService: contactService,
	}
}

func (h *ContactHandler) Create(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("X-User-Username")

	var req models.ContactCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Errors: "Invalid request body",
		})
		return
	}

	result, err := h.contactService.Create(username, &req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Data: result,
	})
}

func (h *ContactHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("X-User-Username")
	vars := mux.Vars(r)
	
	contactID, err := strconv.Atoi(vars["contactId"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Errors: "Invalid contact ID",
		})
		return
	}

	result, err := h.contactService.GetByID(contactID, username)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Data: result,
	})
}

func (h *ContactHandler) Update(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("X-User-Username")
	vars := mux.Vars(r)
	
	contactID, err := strconv.Atoi(vars["contactId"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Errors: "Invalid contact ID",
		})
		return
	}

	var req models.ContactUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Errors: "Invalid request body",
		})
		return
	}

	result, err := h.contactService.Update(contactID, username, &req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Data: result,
	})
}

func (h *ContactHandler) Delete(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("X-User-Username")
	vars := mux.Vars(r)
	
	contactID, err := strconv.Atoi(vars["contactId"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Errors: "Invalid contact ID",
		})
		return
	}

	err = h.contactService.Delete(contactID, username)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Data: "OK",
	})
}

func (h *ContactHandler) Search(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("X-User-Username")

	req := &models.ContactSearchRequest{
		Page: 1,
		Size: 10,
	}

	// Parse query parameters
	if name := r.URL.Query().Get("name"); name != "" {
		req.Name = &name
	}
	if email := r.URL.Query().Get("email"); email != "" {
		req.Email = &email
	}
	if phone := r.URL.Query().Get("phone"); phone != "" {
		req.Phone = &phone
	}
	if page := r.URL.Query().Get("page"); page != "" {
		if p, err := strconv.Atoi(page); err == nil {
			req.Page = p
		}
	}
	if size := r.URL.Query().Get("size"); size != "" {
		if s, err := strconv.Atoi(size); err == nil {
			req.Size = s
		}
	}

	result, err := h.contactService.Search(username, req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":   result.Data,
		"paging": result.Paging,
	})
}