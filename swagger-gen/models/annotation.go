// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Annotation Annotation
//
// Associates an event that explains latency with a timestamp.
// Unlike log statements, annotations are often codes. Ex. "ws" for WireSend
//
// Zipkin v1 core annotations such as "cs" and "sr" have been replaced with
// Span.Kind, which interprets timestamp and duration.
//
//
// swagger:model Annotation
type Annotation struct {

	// Epoch **microseconds** of this event.
	//
	// For example, 1502787600000000 corresponds to 2017-08-15 09:00 UTC
	//
	// This value should be set directly by instrumentation, using the most precise
	// value possible. For example, gettimeofday or multiplying epoch millis by 1000.
	//
	Timestamp int64 `json:"timestamp,omitempty"`

	// Usually a short tag indicating an event, like "error"
	//
	// While possible to add larger data, such as garbage collection details, low
	// cardinality event names both keep the size of spans down and also are easy
	// to search against.
	//
	Value string `json:"value,omitempty"`
}

// Validate validates this annotation
func (m *Annotation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this annotation based on context it is used
func (m *Annotation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Annotation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Annotation) UnmarshalBinary(b []byte) error {
	var res Annotation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
