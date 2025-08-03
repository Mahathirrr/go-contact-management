package service

import (
	"errors"
	"go-backend/internal/models"
	"go-backend/internal/repository"
	"go-backend/internal/utils"
)

type AddressService interface {
	Create(contactID int, username string, req *models.AddressCreateRequest) (*models.AddressResponse, error)
	GetByID(id int, contactID int, username string) (*models.AddressResponse, error)
	Update(id int, contactID int, username string, req *models.AddressUpdateRequest) (*models.AddressResponse, error)
	Delete(id int, contactID int, username string) error
	GetByContactID(contactID int, username string) ([]models.AddressResponse, error)
}

type addressService struct {
	addressRepo repository.AddressRepository
	contactRepo repository.ContactRepository
}

func NewAddressService(addressRepo repository.AddressRepository, contactRepo repository.ContactRepository) AddressService {
	return &addressService{
		addressRepo: addressRepo,
		contactRepo: contactRepo,
	}
}

func (s *addressService) checkContactExists(contactID int, username string) error {
	count, err := s.contactRepo.CountByID(contactID, username)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("contact is not found")
	}
	return nil
}

func (s *addressService) Create(contactID int, username string, req *models.AddressCreateRequest) (*models.AddressResponse, error) {
	if err := s.checkContactExists(contactID, username); err != nil {
		return nil, err
	}

	if err := utils.ValidateStruct(req); err != nil {
		return nil, err
	}

	address := &models.Address{
		Street:     req.Street,
		City:       req.City,
		Province:   req.Province,
		Country:    req.Country,
		PostalCode: req.PostalCode,
		ContactID:  contactID,
	}

	createdAddress, err := s.addressRepo.Create(address)
	if err != nil {
		return nil, err
	}

	return &models.AddressResponse{
		ID:         createdAddress.ID,
		Street:     createdAddress.Street,
		City:       createdAddress.City,
		Province:   createdAddress.Province,
		Country:    createdAddress.Country,
		PostalCode: createdAddress.PostalCode,
	}, nil
}

func (s *addressService) GetByID(id int, contactID int, username string) (*models.AddressResponse, error) {
	if err := s.checkContactExists(contactID, username); err != nil {
		return nil, err
	}

	address, err := s.addressRepo.FindByID(id, contactID)
	if err != nil {
		return nil, err
	}
	if address == nil {
		return nil, errors.New("address is not found")
	}

	return &models.AddressResponse{
		ID:         address.ID,
		Street:     address.Street,
		City:       address.City,
		Province:   address.Province,
		Country:    address.Country,
		PostalCode: address.PostalCode,
	}, nil
}

func (s *addressService) Update(id int, contactID int, username string, req *models.AddressUpdateRequest) (*models.AddressResponse, error) {
	if err := s.checkContactExists(contactID, username); err != nil {
		return nil, err
	}

	if err := utils.ValidateStruct(req); err != nil {
		return nil, err
	}

	count, err := s.addressRepo.CountByID(id, contactID)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("address is not found")
	}

	address := &models.Address{
		ID:         id,
		Street:     req.Street,
		City:       req.City,
		Province:   req.Province,
		Country:    req.Country,
		PostalCode: req.PostalCode,
		ContactID:  contactID,
	}

	if err := s.addressRepo.Update(address); err != nil {
		return nil, err
	}

	return &models.AddressResponse{
		ID:         address.ID,
		Street:     address.Street,
		City:       address.City,
		Province:   address.Province,
		Country:    address.Country,
		PostalCode: address.PostalCode,
	}, nil
}

func (s *addressService) Delete(id int, contactID int, username string) error {
	if err := s.checkContactExists(contactID, username); err != nil {
		return err
	}

	count, err := s.addressRepo.CountByID(id, contactID)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("address is not found")
	}

	return s.addressRepo.Delete(id, contactID)
}

func (s *addressService) GetByContactID(contactID int, username string) ([]models.AddressResponse, error) {
	if err := s.checkContactExists(contactID, username); err != nil {
		return nil, err
	}

	addresses, err := s.addressRepo.FindByContactID(contactID)
	if err != nil {
		return nil, err
	}

	var addressResponses []models.AddressResponse
	for _, address := range addresses {
		addressResponses = append(addressResponses, models.AddressResponse{
			ID:         address.ID,
			Street:     address.Street,
			City:       address.City,
			Province:   address.Province,
			Country:    address.Country,
			PostalCode: address.PostalCode,
		})
	}

	return addressResponses, nil
}