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

// MediaTypeImmutable is an immutable encapsulation of a Media Type
type MediaTypeImmutable struct {
	contained MediaTypeMutable
}

/**
 * Functions
 */

// NewImmutable returns a new empty instance of a MediaTypeImmutable struct
func NewImmutable() MediaType {
	return &MediaTypeImmutable{}
}

// NewImmutableAsContainer returns a new instance of a MediaTypeImmutable struct
// with the values matching the values of the passed MediaTypeMutable struct
func NewImmutableAsContainer(mutable MediaTypeMutable) MediaType {
	return &MediaTypeImmutable{mutable}
}

// MainType returns the "main" (top-level) type as a string
func (m *MediaTypeImmutable) MainType() string {
	return m.contained.MainType()
}

// SubType returns the "sub" type as a string
func (m *MediaTypeImmutable) SubType() string {
	return m.contained.SubType()
}

// Trees returns the split "sub" type as an array of strings split by the namespace separator
func (m *MediaTypeImmutable) Trees() []string {
	return m.contained.Trees()
}

// Prefix returns the prefix of the type's trees
func (m *MediaTypeImmutable) Prefix() string {
	return m.contained.Prefix()
}

// Suffix returns the "suffix" of the type as a string
func (m *MediaTypeImmutable) Suffix() string {
	return m.contained.Suffix()
}

// Parameters returns the defined parameters of the media type
func (m *MediaTypeImmutable) Parameters() map[string]string {
	return m.contained.Parameters()
}

// FullType returns the normalized full type as a string
func (m *MediaTypeImmutable) FullType() string {
	return m.contained.FullType()
}

// Get a string representation conforming to RFC 2045 and RFC 2616
func (m *MediaTypeImmutable) String() string {
	return m.contained.String()
}

// Mutable returns a mutable version of this structure
func (m *MediaTypeImmutable) Mutable() MediaTypeMutable {
	return m.contained
}
