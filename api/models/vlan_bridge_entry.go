// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VlanBridgeEntry vlan bridge entry
//
// swagger:model VlanBridgeEntry
type VlanBridgeEntry struct {

	// Vlan ID
	Vid int64 `json:"Vid,omitempty"`
}

// Validate validates this vlan bridge entry
func (m *VlanBridgeEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this vlan bridge entry based on context it is used
func (m *VlanBridgeEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VlanBridgeEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VlanBridgeEntry) UnmarshalBinary(b []byte) error {
	var res VlanBridgeEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
