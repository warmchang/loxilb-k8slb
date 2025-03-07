// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Node node
//
// swagger:model Node
type Node struct {

	// color
	Color string `json:"color,omitempty"`

	// icon
	Icon string `json:"icon,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// mainstat
	Mainstat float64 `json:"mainstat,omitempty"`

	// node radius
	NodeRadius int64 `json:"nodeRadius,omitempty"`

	// secondarystat
	Secondarystat float64 `json:"secondarystat,omitempty"`

	// subtitle
	Subtitle string `json:"subtitle,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this node
func (m *Node) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this node based on context it is used
func (m *Node) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Node) UnmarshalBinary(b []byte) error {
	var res Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
