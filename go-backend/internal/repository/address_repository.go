package repository

import (
	"database/sql"
	"go-backend/internal/database"
	"go-backend/internal/models"
)

type AddressRepository interface {
	Create(address *models.Address) (*models.Address, error)
	FindByID(id int, contactID int) (*models.Address, error)
	Update(address *models.Address) error
	Delete(id int, contactID int) error
	FindByContactID(contactID int) ([]models.Address, error)
	CountByID(id int, contactID int) (int, error)
}

type addressRepository struct {
	db *sql.DB
}

func NewAddressRepository() AddressRepository {
	return &addressRepository{
		db: database.DB,
	}
}

func (r *addressRepository) Create(address *models.Address) (*models.Address, error) {
	query := `INSERT INTO addresses (street, city, province, country, postal_code, contact_id) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, address.Street, address.City, address.Province, address.Country, address.PostalCode, address.ContactID)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	address.ID = int(id)
	return address, nil
}

func (r *addressRepository) FindByID(id int, contactID int) (*models.Address, error) {
	query := `SELECT id, street, city, province, country, postal_code, contact_id FROM addresses WHERE id = ? AND contact_id = ?`
	row := r.db.QueryRow(query, id, contactID)

	var address models.Address
	err := row.Scan(&address.ID, &address.Street, &address.City, &address.Province, &address.Country, &address.PostalCode, &address.ContactID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &address, nil
}

func (r *addressRepository) Update(address *models.Address) error {
	query := `UPDATE addresses SET street = ?, city = ?, province = ?, country = ?, postal_code = ? WHERE id = ? AND contact_id = ?`
	_, err := r.db.Exec(query, address.Street, address.City, address.Province, address.Country, address.PostalCode, address.ID, address.ContactID)
	return err
}

func (r *addressRepository) Delete(id int, contactID int) error {
	query := `DELETE FROM addresses WHERE id = ? AND contact_id = ?`
	_, err := r.db.Exec(query, id, contactID)
	return err
}

func (r *addressRepository) FindByContactID(contactID int) ([]models.Address, error) {
	query := `SELECT id, street, city, province, country, postal_code, contact_id FROM addresses WHERE contact_id = ?`
	rows, err := r.db.Query(query, contactID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []models.Address
	for rows.Next() {
		var address models.Address
		err := rows.Scan(&address.ID, &address.Street, &address.City, &address.Province, &address.Country, &address.PostalCode, &address.ContactID)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (r *addressRepository) CountByID(id int, contactID int) (int, error) {
	query := `SELECT COUNT(*) FROM addresses WHERE id = ? AND contact_id = ?`
	row := r.db.QueryRow(query, id, contactID)

	var count int
	err := row.Scan(&count)
	return count, err
}