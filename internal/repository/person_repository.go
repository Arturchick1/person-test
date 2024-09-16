package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Arturchick1/person-test/internal/models"
)

type PersonRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *PersonRepository {
	return &PersonRepository{
		db: db,
	}
}

func (p *PersonRepository) GetPerson(ctx context.Context, id int64) (models.PersonDB, error) {
	const fn = "person_repository.GetPerson"

	person := models.PersonDB{}

	row := p.db.QueryRowContext(ctx, "SELECT id, email, phone, firstName, lastName FROM person WHERE id = :id", sql.Named("id", id))
	err := row.Scan(
		&person.ID,
		&person.Email,
		&person.Phone,
		&person.FirstName,
		&person.LastName,
	)
	if err != nil {
		return models.PersonDB{}, fmt.Errorf("%s: %w", fn, err)
	}

	return person, nil
}

func (p *PersonRepository) GetPersons(ctx context.Context, limit int64, offset int64) ([]models.PersonDB, error) {
	const fn = "person_repository.GetPersons"

	persons := make([]models.PersonDB, 0)

	rows, err := p.db.QueryContext(ctx, "SELECT id, email, phone, firstName, lastName FROM person ORDER BY id LIMIT :limit OFFSET :offset",
		sql.Named("limit", limit),
		sql.Named("offset", offset))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer rows.Close()

	for rows.Next() {
		p := models.PersonDB{}

		err := rows.Scan(
			&p.ID,
			&p.Email,
			&p.Phone,
			&p.FirstName,
			&p.LastName,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}

		persons = append(persons, p)
	}

	return persons, nil
}

func (p *PersonRepository) UpdatePerson(ctx context.Context, person models.PersonDTO, id int64) (int64, error) {
	const fn = "person_repository.UpdatePerson"

	res, err := p.db.ExecContext(ctx, "UPDATE person SET email = :email, phone = :phone, firstName = :firstName, lastName = :lastName WHERE id = :id",
		sql.Named("email", person.Email),
		sql.Named("phone", person.Phone),
		sql.Named("firstName", person.FirstName),
		sql.Named("lastName", person.LastName),
		sql.Named("id", id))
	if err != nil {
		return 0, fmt.Errorf("%s: %w", fn, err)
	}

	idLastUpdated, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", fn, err)
	}

	return idLastUpdated, nil
}

func (p *PersonRepository) CreatePerson(ctx context.Context, person models.PersonDTO) (int64, error) {
	const fn = "person_repository.CreatePerson"

	res, err := p.db.ExecContext(ctx, "INSERT INTO person(email, phone, firstName, lastName) VALUES (:email, :phone, :firstName, :lastName)",
		sql.Named("email", person.Email),
		sql.Named("phone", person.Phone),
		sql.Named("firstName", person.FirstName),
		sql.Named("lastName", person.LastName))
	if err != nil {
		return 0, fmt.Errorf("%s: %w", fn, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", fn, err)
	}

	return id, nil
}

func (p *PersonRepository) DeletePerson(ctx context.Context, id int64) error {
	const fn = "person_repository.DeletePerson"

	_, err := p.db.Exec("DELETE FROM person WHERE id = :id", sql.Named("id", id))
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
