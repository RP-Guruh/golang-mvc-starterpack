package repositories

import (
	"database/sql"
	"golang_mvc_starterpack/models"
)

type PeopleRepository interface {
	Store(p *models.People) error
	Index() ([]models.People, error)
	Show(id string) (models.People, error)
	Update(p models.People) error
	Delete(id string) error
}

type peopleRepository struct {
	db *sql.DB
}

func NewPeopleRepository(db *sql.DB) PeopleRepository {
	return &peopleRepository{db}
}

func (r *peopleRepository) Index() ([]models.People, error) {
	query := `
		SELECT * FROM people
	`

	result, err := r.db.Query(query)
	if err != nil {

		return nil, err
	}
	defer result.Close()

	var people []models.People
	for result.Next() {
		var p models.People
		if err := result.Scan(&p.ID, &p.FirstName, &p.LastName, &p.PlaceOfBirth, &p.DateOfBirth, &p.Address, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		people = append(people, p)
	}
	return people, nil
}

func (r *peopleRepository) Show(id string) (models.People, error) {
	var p models.People

	query := "SELECT id, first_name, last_name, place_of_birth, date_of_birth, address, created_at, updated_at FROM people WHERE id = ?"
	row := r.db.QueryRow(query, id)

	err := row.Scan(
		&p.ID,
		&p.FirstName,
		&p.LastName,
		&p.PlaceOfBirth,
		&p.DateOfBirth,
		&p.Address,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		return models.People{}, err
	}

	return p, nil
}

func (r *peopleRepository) Store(p *models.People) error {
	query := `
        INSERT INTO people (first_name, last_name, place_of_birth, date_of_birth, address, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, NOW(), NOW())
    `
	result, err := r.db.Exec(query,
		p.FirstName,
		p.LastName,
		p.PlaceOfBirth,
		p.DateOfBirth,
		p.Address,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = uint(id)

	return nil
}

func (r *peopleRepository) Update(p models.People) error {
	query := `
        UPDATE people 
        SET first_name = ?, 
            last_name = ?, 
            place_of_birth = ?, 
            date_of_birth = ?, 
            address = ?, 
            updated_at = NOW()
        WHERE id = ?
    `
	result, err := r.db.Exec(query,
		p.FirstName,
		p.LastName,
		p.PlaceOfBirth,
		p.DateOfBirth,
		p.Address,
		p.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *peopleRepository) Delete(id string) error {
	query := `
        DELETE FROM people WHERE id = ?
    `
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
