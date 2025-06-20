// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SessionUlClEntry session ul cl entry
//
// swagger:model SessionUlClEntry
type SessionUlClEntry struct {

	// ulcl argument
	UlclArgument *SessionUlClEntryUlclArgument `json:"ulclArgument,omitempty"`

	// IP address and netmask
	// Required: true
	UlclIdent *string `json:"ulclIdent"`
}

// Validate validates this session ul cl entry
func (m *SessionUlClEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUlclArgument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUlclIdent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SessionUlClEntry) validateUlclArgument(formats strfmt.Registry) error {
	if swag.IsZero(m.UlclArgument) { // not required
		return nil
	}

	if m.UlclArgument != nil {
		if err := m.UlclArgument.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ulclArgument")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ulclArgument")
			}
			return err
		}
	}

	return nil
}

func (m *SessionUlClEntry) validateUlclIdent(formats strfmt.Registry) error {

	if err := validate.Required("ulclIdent", "body", m.UlclIdent); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this session ul cl entry based on the context it is used
func (m *SessionUlClEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUlclArgument(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SessionUlClEntry) contextValidateUlclArgument(ctx context.Context, formats strfmt.Registry) error {

	if m.UlclArgument != nil {
		if err := m.UlclArgument.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ulclArgument")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ulclArgument")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SessionUlClEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SessionUlClEntry) UnmarshalBinary(b []byte) error {
	var res SessionUlClEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SessionUlClEntryUlclArgument session ul cl entry ulcl argument
//
// swagger:model SessionUlClEntryUlclArgument
type SessionUlClEntryUlclArgument struct {

	// QFI number
	Qfi int64 `json:"qfi,omitempty"`

	// Access network IP address
	UlclIP string `json:"ulclIP,omitempty"`
}

// Validate validates this session ul cl entry ulcl argument
func (m *SessionUlClEntryUlclArgument) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this session ul cl entry ulcl argument based on context it is used
func (m *SessionUlClEntryUlclArgument) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SessionUlClEntryUlclArgument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SessionUlClEntryUlclArgument) UnmarshalBinary(b []byte) error {
	var res SessionUlClEntryUlclArgument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
