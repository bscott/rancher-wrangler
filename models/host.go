package models

import (
	"time"

	"encoding/json"

	"github.com/markbates/pop"

	"github.com/markbates/validate"

	"github.com/satori/go.uuid"
)

type Host struct {
	ID uuid.UUID `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Description string `json:"description" db:"desc"`
	Url string `json:"url" db:"url"`
	ApiVersion string `json:"api_version" db:"api_version"`
	AccessKey string `json:"access_key" db:"access_key"`
	SecretKey string `json:"secret_key" db:"secret_key"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (h Host) String() string {
	b, _ := json.Marshal(h)
	return string(b)
}

// Hosts is not required by pop and may be deleted
type Hosts []Host

// String is not required by pop and may be deleted
func (h Hosts) String() string {
	b, _ := json.Marshal(h)
	return string(b)
}

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (h *Host) Validate(tx *pop.Connection) (*validate.Errors, error) {

	return validate.NewErrors(), nil

}

// ValidateSave gets run everytime you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (h *Host) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run everytime you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (h *Host) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
