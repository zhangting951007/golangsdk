package flavors

import (
	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/pagination"
)

// GetResult temporarily holds the response from a Get call.
type GetResult struct {
	golangsdk.Result
}

// Extract provides access to the individual Flavor returned by the Get function.
func (r GetResult) Extract() (*Flavor, error) {
	var s struct {
		Flavor *Flavor `json:"flavor"`
	}
	err := r.ExtractInto(&s)
	return s.Flavor, err
}

// Flavor records represent (virtual) hardware configurations for server resources in a region.
type Flavor struct {
	// The flavor's unique identifier.
	// Contains 0 if the ID is not an integer.
	ID int `json:"id"`

	// The RAM capacity for the flavor.
	RAM int `json:"ram"`

	// The Name field provides a human-readable moniker for the flavor.
	Name string `json:"name"`

	// Links to access the flavor.
	Links []golangsdk.Link

	// The flavor's unique identifier as a string
	StrID string `json:"str_id"`
}

// FlavorPage contains a single page of the response from a List call.
type FlavorPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines if a page contains any results.
func (page FlavorPage) IsEmpty() (bool, error) {
	flavors, err := ExtractFlavors(page)
	return len(flavors) == 0, err
}

// NextPageURL uses the response's embedded link reference to navigate to the next page of results.
func (page FlavorPage) NextPageURL() (string, error) {
	var s struct {
		Links []golangsdk.Link `json:"flavors_links"`
	}
	err := page.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(s.Links)
}

// ExtractFlavors provides access to the list of flavors in a page acquired from the List operation.
func ExtractFlavors(r pagination.Page) ([]Flavor, error) {
	var s struct {
		Flavors []Flavor `json:"flavors"`
	}
	err := (r.(FlavorPage)).ExtractInto(&s)
	return s.Flavors, err
}