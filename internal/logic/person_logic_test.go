package logic

import (
	"context"
	"testing"

	"github.com/Arturchick1/person-test/internal/models"
	"github.com/Arturchick1/person-test/internal/repository"
	"github.com/Arturchick1/person-test/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPersonWhenOk(t *testing.T) {
	storage, err := database.New("../../storage/storage.db")
	assert.NoError(t, err)
	defer storage.DB.Close()

	personRepository := repository.New(storage.DB)
	personLogic := New(personRepository)
	ctx, _ := context.WithCancel(context.Background())
	var id int64 = 1

	person, err := personLogic.GetPerson(ctx, id)
	require.NoError(t, err)

	assert.Equal(t, id, person.ID)
	assert.NotEmpty(t, person.Email)
	assert.NotEmpty(t, person.Phone)
	assert.NotEmpty(t, person.FirstName)
	assert.NotEmpty(t, person.LastName)
}

func TestGetPersonWhenNoPerson(t *testing.T) {
	storage, err := database.New("../../storage/storage.db")
	assert.NoError(t, err)
	defer storage.DB.Close()

	personRepository := repository.New(storage.DB)
	personLogic := New(personRepository)
	ctx, _ := context.WithCancel(context.Background())
	var id int64 = -1

	person, err := personLogic.GetPerson(ctx, id)
	require.Error(t, err)

	assert.Empty(t, person.ID)
	assert.Empty(t, person.Email)
	assert.Empty(t, person.Phone)
	assert.Empty(t, person.FirstName)
	assert.Empty(t, person.LastName)
}

func TestInsertUpdateAndDeletePerson(t *testing.T) {
	storage, err := database.New("../../storage/storage.db")
	assert.NoError(t, err)
	defer storage.DB.Close()

	personRepository := repository.New(storage.DB)
	personLogic := New(personRepository)
	ctx, _ := context.WithCancel(context.Background())

	personDTO := models.PersonDTO{
		Email:     "test",
		Phone:     "test",
		FirstName: "test",
		LastName:  "test",
	}

	id, err := personLogic.CreatePerson(ctx, personDTO)
	require.NoError(t, err)

	person, err := personLogic.GetPerson(ctx, id)
	require.NoError(t, err)

	assert.Equal(t, id, person.ID)
	assert.NotEmpty(t, person.Email)
	assert.NotEmpty(t, person.Phone)
	assert.NotEmpty(t, person.FirstName)
	assert.NotEmpty(t, person.LastName)

	personUpdate := models.PersonDTO{
		Email:     personDTO.Email,
		Phone:     personDTO.Phone,
		FirstName: "test1",
		LastName:  "test1",
	}

	idUpdated, err := personLogic.UpdatePerson(ctx, personUpdate, id)
	require.NoError(t, err)

	personUpdated, err := personLogic.GetPerson(ctx, idUpdated)
	require.NoError(t, err)

	assert.Equal(t, id, personUpdated.ID)
	assert.NotEmpty(t, personUpdated.Email)
	assert.NotEmpty(t, personUpdated.Phone)
	assert.Equal(t, personUpdated.FirstName, personUpdate.FirstName)
	assert.Equal(t, personUpdated.LastName, personUpdate.LastName)

	err = personLogic.DeletePerson(ctx, id)
	require.NoError(t, err)

	p, err := personLogic.GetPerson(ctx, id)
	require.Error(t, err)

	assert.Empty(t, p.ID)
	assert.Empty(t, p.Email)
	assert.Empty(t, p.Phone)
	assert.Empty(t, p.FirstName)
	assert.Empty(t, p.LastName)
}
