package models

type Address struct {
	ID         int     `json:"id" db:"id"`
	Street     *string `json:"street" db:"street"`
	City       *string `json:"city" db:"city"`
	Province   *string `json:"province" db:"province"`
	Country    string  `json:"country" db:"country"`
	PostalCode string  `json:"postal_code" db:"postal_code"`
	ContactID  int     `json:"contact_id" db:"contact_id"`
}

type AddressCreateRequest struct {
	Street     *string `json:"street,omitempty" validate:"omitempty,max=255"`
	City       *string `json:"city,omitempty" validate:"omitempty,max=100"`
	Province   *string `json:"province,omitempty" validate:"omitempty,max=100"`
	Country    string  `json:"country" validate:"required,max=100"`
	PostalCode string  `json:"postal_code" validate:"required,max=10"`
}

type AddressUpdateRequest struct {
	Street     *string `json:"street,omitempty" validate:"omitempty,max=255"`
	City       *string `json:"city,omitempty" validate:"omitempty,max=100"`
	Province   *string `json:"province,omitempty" validate:"omitempty,max=100"`
	Country    string  `json:"country" validate:"required,max=100"`
	PostalCode string  `json:"postal_code" validate:"required,max=10"`
}

type AddressResponse struct {
	ID         int     `json:"id"`
	Street     *string `json:"street"`
	City       *string `json:"city"`
	Province   *string `json:"province"`
	Country    string  `json:"country"`
	PostalCode string  `json:"postal_code"`
}