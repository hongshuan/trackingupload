package common

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Controller handle all base methods
type Controller struct {
}

// SendJSON marshals v to a json struct and sends appropriate headers to w
func (c *Controller) SendJSON(w http.ResponseWriter, r *http.Request, v interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")

	b, err := json.Marshal(v)

	if err != nil {
		log.Print(fmt.Sprintf("Error while encoding JSON: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Internal server error"}`)
	} else {
		w.WriteHeader(code)
		io.WriteString(w, string(b))
	}
}

////////////////////////////////////////////////////////////

package kittiesbundle

import (
	"net/http"

	"github.com/laibulle/kitties/app/common"
)

// Kitty struct
type Kitty struct {
	Name      string `json:"name"`
	Breed     string `json:"breed"`
	BirthDate string `json:"birthDate"`
}

// NewKitty create a new Kitty
func NewKitty(name string, breed string, birthDate string) *Kitty {
	return &Kitty{
		Name:      name,
		Breed:     breed,
		BirthDate: birthDate,
	}
}

// KittiesController struct
type KittiesController struct {
	common.Controller
}

// Index func return all kitties in database
func (c *KittiesController) Index(w http.ResponseWriter, r *http.Request) {
	c.SendJSON(w, r,
		[]*Kitty{NewKitty("Gaspart", "British", "2016-07-05")},
		http.StatusOK,
	)
}
