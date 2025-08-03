package service

import (
	"errors"
	"go-backend/internal/models"
	"go-backend/internal/repository"
	"go-backend/internal/utils"
	"math"
)

type ContactService interface {
	Create(username string, req *models.ContactCreateRequest) (*models.ContactResponse, error)
	GetByID(id int, username string) (*models.ContactResponse, error)
	Update(id int, username string, req *models.ContactUpdateRequest) (*models.ContactResponse, error)
	Delete(id int, username string) error
	Search(username string, req *models.ContactSearchRequest) (*models.ContactSearchResponse, error)
}

type contactService struct {
	contactRepo repository.ContactRepository
}

func NewContactService(contactRepo repository.ContactRepository) ContactService {
	return &contactService{
		contactRepo: contactRepo,
	}
}

func (s *contactService) Create(username string, req *models.ContactCreateRequest) (*models.ContactResponse, error) {
	if err := utils.ValidateStruct(req); err != nil {
		return nil, err
	}

	contact := &models.Contact{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Username:  username,
	}

	createdContact, err := s.contactRepo.Create(contact)
	if err != nil {
		return nil, err
	}

	return &models.ContactResponse{
		ID:        createdContact.ID,
		FirstName: createdContact.FirstName,
		LastName:  createdContact.LastName,
		Email:     createdContact.Email,
		Phone:     createdContact.Phone,
	}, nil
}

func (s *contactService) GetByID(id int, username string) (*models.ContactResponse, error) {
	contact, err := s.contactRepo.FindByID(id, username)
	if err != nil {
		return nil, err
	}
	if contact == nil {
		return nil, errors.New("contact is not found")
	}

	return &models.ContactResponse{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Email:     contact.Email,
		Phone:     contact.Phone,
	}, nil
}

func (s *contactService) Update(id int, username string, req *models.ContactUpdateRequest) (*models.ContactResponse, error) {
	if err := utils.ValidateStruct(req); err != nil {
		return nil, err
	}

	count, err := s.contactRepo.CountByID(id, username)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("contact is not found")
	}

	contact := &models.Contact{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Username:  username,
	}

	if err := s.contactRepo.Update(contact); err != nil {
		return nil, err
	}

	return &models.ContactResponse{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Email:     contact.Email,
		Phone:     contact.Phone,
	}, nil
}

func (s *contactService) Delete(id int, username string) error {
	count, err := s.contactRepo.CountByID(id, username)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("contact is not found")
	}

	return s.contactRepo.Delete(id, username)
}

func (s *contactService) Search(username string, req *models.ContactSearchRequest) (*models.ContactSearchResponse, error) {
	// Set defaults
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	if req.Size > 100 {
		req.Size = 100
	}

	if err := utils.ValidateStruct(req); err != nil {
		return nil, err
	}

	contacts, totalItems, err := s.contactRepo.Search(req, username)
	if err != nil {
		return nil, err
	}

	var contactResponses []models.ContactResponse
	for _, contact := range contacts {
		contactResponses = append(contactResponses, models.ContactResponse{
			ID:        contact.ID,
			FirstName: contact.FirstName,
			LastName:  contact.LastName,
			Email:     contact.Email,
			Phone:     contact.Phone,
		})
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(req.Size)))

	return &models.ContactSearchResponse{
		Data: contactResponses,
		Paging: models.PagingResponse{
			Page:      req.Page,
			TotalPage: totalPages,
			TotalItem: totalItems,
		},
	}, nil
}