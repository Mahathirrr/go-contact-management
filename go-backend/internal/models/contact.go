package models

type Contact struct {
	ID        int     `json:"id" db:"id"`
	FirstName string  `json:"first_name" db:"first_name"`
	LastName  *string `json:"last_name" db:"last_name"`
	Email     *string `json:"email" db:"email"`
	Phone     *string `json:"phone" db:"phone"`
	Username  string  `json:"username" db:"username"`
}

type ContactCreateRequest struct {
	FirstName string  `json:"first_name" validate:"required,max=100"`
	LastName  *string `json:"last_name,omitempty" validate:"omitempty,max=100"`
	Email     *string `json:"email,omitempty" validate:"omitempty,email,max=200"`
	Phone     *string `json:"phone,omitempty" validate:"omitempty,max=20"`
}

type ContactUpdateRequest struct {
	FirstName string  `json:"first_name" validate:"required,max=100"`
	LastName  *string `json:"last_name,omitempty" validate:"omitempty,max=100"`
	Email     *string `json:"email,omitempty" validate:"omitempty,email,max=200"`
	Phone     *string `json:"phone,omitempty" validate:"omitempty,max=20"`
}

type ContactResponse struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
}

type ContactSearchRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
	Page  int     `json:"page" validate:"min=1"`
	Size  int     `json:"size" validate:"min=1,max=100"`
}

type ContactSearchResponse struct {
	Data   []ContactResponse `json:"data"`
	Paging PagingResponse    `json:"paging"`
}

type PagingResponse struct {
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
	TotalItem int `json:"total_item"`
}