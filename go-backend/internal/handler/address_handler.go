package handler

import (
	"encoding/json"
	"go-backend/internal/models"
	"go-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AddressHandler struct {
	addressService service.AddressService
}

func NewAddressHandler(addressService service.AddressService) *AddressHandler {
	return &AddressHandler{
		addressService: addressService,
	}
}

func (h *AddressHandler) Create(w http.ResponseWriter, r *http.Request) {
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

	var req models.AddressCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Errors: "Invalid request body",
		})
		return
	}

	result, err := h.addressService.Create(contactID, username, &req)
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

func (h *AddressHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	addressID, err := strconv.Atoi(vars["addressId"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Errors: "Invalid address ID",
		})
		return
	}

	result, err := h.addressService.GetByID(addressID, contactID, username)
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

func (h *AddressHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	addressID, err := strconv.Atoi(vars["addressId"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Errors: "Invalid address ID",
		})
		return
	}

	var req models.AddressUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Errors: "Invalid request body",
		})
		return
	}

	result, err := h.addressService.Update(addressID, contactID, username, &req)
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

func (h *AddressHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	addressID, err := strconv.Atoi(vars["addressId"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Errors: "Invalid address ID",
		})
		return
	}

	err = h.addressService.Delete(addressID, contactID, username)
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

func (h *AddressHandler) GetByContactID(w http.ResponseWriter, r *http.Request) {
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

	result, err := h.addressService.GetByContactID(contactID, username)
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