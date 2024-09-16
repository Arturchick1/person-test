package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Arturchick1/person-test/internal/logic"
	"github.com/Arturchick1/person-test/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type PersonHandlers struct {
	personLogic *logic.PersonLogic
	log         *logrus.Logger
}

func New(personLogic *logic.PersonLogic, log *logrus.Logger) *PersonHandlers {
	return &PersonHandlers{
		personLogic: personLogic,
		log:         log,
	}
}

func (p *PersonHandlers) GetOne(c echo.Context) error {
	const fn = "person_handlers.GetOne"

	idStr := c.Param("id")
	if idStr == "" {
		p.log.Error(fmt.Errorf("%s: param ID is empty", fn))
		return c.String(http.StatusBadRequest, "incorrect value id")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		p.log.Error(fmt.Errorf("%s: %w", fn, err))
		return c.NoContent(http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	person, err := p.personLogic.GetPerson(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	c.Response().Header().Set("Content-type", "application/json")
	return c.JSON(http.StatusOK, person)
}

func (p *PersonHandlers) Get(c echo.Context) error {
	const fn = "person_handlers.Get"

	limitStr := c.QueryParam("limit")
	if limitStr == "" {
		limitStr = "3"
	}

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		p.log.Error(fmt.Errorf("%s: %w", fn, err))
		return c.String(http.StatusBadRequest, "incorrect value limit")
	}

	offsetStr := c.QueryParam("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}

	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		p.log.Error(fmt.Errorf("%s: %w", fn, err))
		return c.String(http.StatusBadRequest, "incorrect value limit")
	}

	ctx := c.Request().Context()

	persons, err := p.personLogic.GetPersons(ctx, limit, offset)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	c.Response().Header().Set("Content-type", "application/json")
	return c.JSON(http.StatusOK, persons)
}

func (p *PersonHandlers) Update(c echo.Context) error {
	const fn = "person_handlers.Update"

	idStr := c.Param("id")
	if idStr == "" {
		p.log.Error(fmt.Errorf("%s: param ID is empty", fn))
		return c.NoContent(http.StatusBadRequest)
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		p.log.Error(fmt.Errorf("%s: %w", fn, err))
		return c.String(http.StatusBadRequest, "incorrect value id")
	}

	personDTO := models.PersonDTO{}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(c.Request().Body)
	if err != nil {
		p.log.Error(fmt.Errorf("%s: %w", fn, err))
		return c.NoContent(http.StatusBadRequest)
	}

	if err = json.Unmarshal(buf.Bytes(), &personDTO); err != nil {
		p.log.Error(fmt.Errorf("%s: %w", fn, err))
		return c.NoContent(http.StatusBadRequest)
	}

	person := models.PersonDTO{
		ID:        nil,
		Email:     personDTO.Email,
		Phone:     personDTO.Phone,
		FirstName: personDTO.FirstName,
		LastName:  personDTO.LastName,
	}

	ctx := c.Request().Context()

	personUpdated, err := p.personLogic.UpdatePerson(ctx, person, id)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	c.Response().Header().Set("Content-type", "application/json")
	return c.JSON(http.StatusOK, personUpdated)
}

func (p *PersonHandlers) Create(c echo.Context) error {
	const fn = "person_handlers.Create"

	personDTO := models.PersonDTO{}

	var buf bytes.Buffer
	_, err := buf.ReadFrom(c.Request().Body)
	if err != nil {
		p.log.Error(fmt.Errorf("%s: %w", fn, err))
		return c.NoContent(http.StatusBadRequest)
	}

	if err = json.Unmarshal(buf.Bytes(), &personDTO); err != nil {
		p.log.Error(fmt.Errorf("%s: %w", fn, err))
		return c.NoContent(http.StatusBadRequest)
	}

	person := models.PersonDTO{
		ID:        nil,
		Email:     personDTO.Email,
		Phone:     personDTO.Phone,
		FirstName: personDTO.FirstName,
		LastName:  personDTO.LastName,
	}

	ctx := c.Request().Context()

	idLastCreated, err := p.personLogic.CreatePerson(ctx, person)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	personCreated, err := p.personLogic.GetPerson(ctx, idLastCreated)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	c.Response().Header().Set("Content-type", "application/json")
	return c.JSON(http.StatusCreated, personCreated)
}

func (p *PersonHandlers) Delete(c echo.Context) error {
	const fn = "person_handlers.Delete"

	idStr := c.Param("id")
	if idStr == "" {
		p.log.Error(fmt.Errorf("%s: param ID is empty", fn))
		return c.NoContent(http.StatusBadRequest)
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		p.log.Error(fmt.Errorf("%s: %w", fn, err))
		return c.String(http.StatusBadRequest, "incorrect value id")
	}

	ctx := c.Request().Context()

	err = p.personLogic.DeletePerson(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	c.Response().Header().Set("Content-type", "application/json")
	return c.NoContent(http.StatusOK)
}
