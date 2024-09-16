package models

type PersonDTO struct {
	ID        *int64 `json:"id"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Person struct {
	ID        int64  `json:"id"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type PersonDB struct {
	ID        int64  `json:"id"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
