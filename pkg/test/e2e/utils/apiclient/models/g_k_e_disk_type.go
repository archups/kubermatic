// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GKEDiskType GKEDiskType represents a object of GKE disk type.
//
// swagger:model GKEDiskType
type GKEDiskType struct {

	// DefaultDiskSizeGb: Server-defined default disk size in GB.
	DefaultDiskSizeGb int64 `json:"defaultDiskSizeGb,omitempty"`

	// Description: An optional description of this resource.
	Description string `json:"description,omitempty"`

	// Kind: Type of the resource. Always compute#diskType for
	// disk types.
	Kind string `json:"kind,omitempty"`

	// Name of the resource.
	Name string `json:"name,omitempty"`
}

// Validate validates this g k e disk type
func (m *GKEDiskType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this g k e disk type based on context it is used
func (m *GKEDiskType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GKEDiskType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GKEDiskType) UnmarshalBinary(b []byte) error {
	var res GKEDiskType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
