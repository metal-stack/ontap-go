// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AwsKmsState Indicates whether or not the Amazon Web Services Key Management Service (AWS KMS) key protection is available cluster-wide.
//
// swagger:model aws_kms_state
type AwsKmsState struct {

	// Set to true when AWS KMS key protection is available on all nodes of the cluster.
	ClusterState *bool `json:"cluster_state,omitempty" yaml:"cluster_state,omitempty"`

	// Code corresponding to the message. Returns a 0 if AWS KMS key protection is available on all nodes of the cluster.
	// Example: 346758
	Code *string `json:"code,omitempty" yaml:"code,omitempty"`

	// Error message set when cluster_state is false.
	// Example: AWS KMS key protection is unavailable on the following nodes: node1, node2.
	Message *string `json:"message,omitempty" yaml:"message,omitempty"`
}

// Validate validates this aws kms state
func (m *AwsKmsState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aws kms state based on context it is used
func (m *AwsKmsState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsKmsState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsKmsState) UnmarshalBinary(b []byte) error {
	var res AwsKmsState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
