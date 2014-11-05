/**
 * MediaType
 *
 * Copyright © 2014 Trevor N. Suarez (Rican7)
 */

// This file provides provides an immutable implementation of the MediaType interface

package mediatype

/**
 * Types
 */

// Immutable is an immutable encapsulation of a Media Type
type Immutable struct {
	contained Mutable
}

/**
 * Functions
 */

// NewImmutable returns a new empty instance of a Immutable struct
func NewImmutable() MediaType {
	return &Immutable{}
}

// NewImmutableAsContainer returns a new instance of a Immutable struct
// with the values matching the values of the passed Mutable struct
func NewImmutableAsContainer(mutable Mutable) MediaType {
	return &Immutable{mutable}
}

// MainType returns the "main" (top-level) type as a string
func (m *Immutable) MainType() string {
	return m.contained.MainType()
}

// SubType returns the "sub" type as a string
func (m *Immutable) SubType() string {
	return m.contained.SubType()
}

// Trees returns the split "sub" type as an array of strings split by the namespace separator
func (m *Immutable) Trees() []string {
	return m.contained.Trees()
}

// Prefix returns the prefix of the type's trees
func (m *Immutable) Prefix() string {
	return m.contained.Prefix()
}

// Suffix returns the "suffix" of the type as a string
func (m *Immutable) Suffix() string {
	return m.contained.Suffix()
}

// Parameters returns the defined parameters of the media type
func (m *Immutable) Parameters() map[string]string {
	return m.contained.Parameters()
}

// FullType returns the normalized full type as a string
func (m *Immutable) FullType() string {
	return m.contained.FullType()
}

// Get a string representation conforming to RFC 2045 and RFC 2616
func (m *Immutable) String() string {
	return m.contained.String()
}

// Mutable returns a mutable version of this structure
func (m *Immutable) Mutable() Mutable {
	return m.contained
}
