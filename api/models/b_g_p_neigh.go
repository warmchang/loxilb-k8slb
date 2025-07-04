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

// BGPNeigh b g p neigh
//
// swagger:model BGPNeigh
type BGPNeigh struct {

	// BGP Neighbor IP address
	// Required: true
	IPAddress *string `json:"ipAddress"`

	// Remote AS number
	// Required: true
	RemoteAs *int64 `json:"remoteAs"`

	// Remote Connect Port (default 179)
	RemotePort int64 `json:"remotePort,omitempty"`

	// Enable multi-hop peering (if needed)
	SetMultiHop bool `json:"setMultiHop,omitempty"`
}

// Validate validates this b g p neigh
func (m *BGPNeigh) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteAs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BGPNeigh) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ipAddress", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

func (m *BGPNeigh) validateRemoteAs(formats strfmt.Registry) error {

	if err := validate.Required("remoteAs", "body", m.RemoteAs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this b g p neigh based on context it is used
func (m *BGPNeigh) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BGPNeigh) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BGPNeigh) UnmarshalBinary(b []byte) error {
	var res BGPNeigh
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
