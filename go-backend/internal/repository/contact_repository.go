package repository

import (
	"database/sql"
	"fmt"
	"go-backend/internal/database"
	"go-backend/internal/models"
	"strings"
)

type ContactRepository interface {
	Create(contact *models.Contact) (*models.Contact, error)
	FindByID(id int, username string) (*models.Contact, error)
	Update(contact *models.Contact) error
	Delete(id int, username string) error
	Search(req *models.ContactSearchRequest, username string) ([]models.Contact, int, error)
	CountByID(id int, username string) (int, error)
}

type contactRepository struct {
	db *sql.DB
}

func NewContactRepository() ContactRepository {
	return &contactRepository{
		db: database.DB,
	}
}

func (r *contactRepository) Create(contact *models.Contact) (*models.Contact, error) {
	query := `INSERT INTO contacts (first_name, last_name, email, phone, username) VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, contact.FirstName, contact.LastName, contact.Email, contact.Phone, contact.Username)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	contact.ID = int(id)
	return contact, nil
}

func (r *contactRepository) FindByID(id int, username string) (*models.Contact, error) {
	query := `SELECT id, first_name, last_name, email, phone, username FROM contacts WHERE id = ? AND username = ?`
	row := r.db.QueryRow(query, id, username)

	var contact models.Contact
	err := row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone, &contact.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &contact, nil
}

func (r *contactRepository) Update(contact *models.Contact) error {
	query := `UPDATE contacts SET first_name = ?, last_name = ?, email = ?, phone = ? WHERE id = ? AND username = ?`
	_, err := r.db.Exec(query, contact.FirstName, contact.LastName, contact.Email, contact.Phone, contact.ID, contact.Username)
	return err
}

func (r *contactRepository) Delete(id int, username string) error {
	query := `DELETE FROM contacts WHERE id = ? AND username = ?`
	_, err := r.db.Exec(query, id, username)
	return err
}

func (r *contactRepository) Search(req *models.ContactSearchRequest, username string) ([]models.Contact, int, error) {
	// Build WHERE clause
	var conditions []string
	var args []interface{}

	conditions = append(conditions, "username = ?")
	args = append(args, username)

	if req.Name != nil && *req.Name != "" {
		conditions = append(conditions, "(first_name LIKE ? OR last_name LIKE ?)")
		namePattern := "%" + *req.Name + "%"
		args = append(args, namePattern, namePattern)
	}

	if req.Email != nil && *req.Email != "" {
		conditions = append(conditions, "email LIKE ?")
		args = append(args, "%"+*req.Email+"%")
	}

	if req.Phone != nil && *req.Phone != "" {
		conditions = append(conditions, "phone LIKE ?")
		args = append(args, "%"+*req.Phone+"%")
	}

	whereClause := strings.Join(conditions, " AND ")

	// Count total items
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM contacts WHERE %s", whereClause)
	var totalItems int
	err := r.db.QueryRow(countQuery, args...).Scan(&totalItems)
	if err != nil {
		return nil, 0, err
	}

	// Get contacts with pagination
	offset := (req.Page - 1) * req.Size
	query := fmt.Sprintf("SELECT id, first_name, last_name, email, phone, username FROM contacts WHERE %s LIMIT ? OFFSET ?", whereClause)
	args = append(args, req.Size, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var contact models.Contact
		err := rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone, &contact.Username)
		if err != nil {
			return nil, 0, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, totalItems, nil
}

func (r *contactRepository) CountByID(id int, username string) (int, error) {
	query := `SELECT COUNT(*) FROM contacts WHERE id = ? AND username = ?`
	row := r.db.QueryRow(query, id, username)

	var count int
	err := row.Scan(&count)
	return count, err
}