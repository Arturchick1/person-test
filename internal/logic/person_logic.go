package logic

import (
	"context"
	"fmt"

	"github.com/Arturchick1/person-test/internal/models"
	"github.com/Arturchick1/person-test/internal/repository"
)

type PersonLogic struct {
	personRepository *repository.PersonRepository
}

func New(personRepository *repository.PersonRepository) *PersonLogic {
	return &PersonLogic{
		personRepository: personRepository,
	}
}

func (p *PersonLogic) GetPerson(ctx context.Context, id int64) (models.Person, error) {
	const fn = "person_logic.GetPerson"

	person, err := p.personRepository.GetPerson(ctx, id)
	if err != nil {
		return models.Person{}, fmt.Errorf("%s: %w", fn, err)
	}

	return models.Person{
		ID:        person.ID,
		Email:     person.Email,
		Phone:     person.Phone,
		FirstName: person.FirstName,
		LastName:  person.LastName,
	}, nil
}

func (p *PersonLogic) GetPersons(ctx context.Context, limit int64, offset int64) ([]models.Person, error) {
	const fn = "person_logic.GetPersons"

	persons, err := p.personRepository.GetPersons(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	personsLogic := make([]models.Person, len(persons), cap(persons))

	for i, v := range persons {
		personsLogic[i] = models.Person{
			ID:        v.ID,
			Email:     v.Email,
			Phone:     v.Phone,
			FirstName: v.FirstName,
			LastName:  v.LastName}
	}

	return personsLogic, nil
}

func (p *PersonLogic) UpdatePerson(ctx context.Context, person models.PersonDTO, id int64) (int64, error) {
	const fn = "person_logic.UpdatePerson"

	idLastUpdated, err := p.personRepository.UpdatePerson(ctx, person, id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", fn, err)
	}

	return idLastUpdated, nil
}

func (p *PersonLogic) CreatePerson(ctx context.Context, person models.PersonDTO) (int64, error) {
	const fn = "person_logic.CreatePerson"

	idLastCreated, err := p.personRepository.CreatePerson(ctx, person)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", fn, err)
	}

	return idLastCreated, nil
}

func (p *PersonLogic) DeletePerson(ctx context.Context, id int64) error {
	const fn = "person_logic.DeletePerson"

	err := p.personRepository.DeletePerson(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
