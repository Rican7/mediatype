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

// An immutable struct of a Media Type
type MediaTypeImmutable struct {
	contained MediaTypeMutable
}

/**
 * Functions
 */

// Return a New instance of a MediaTypeImmutable struct
func NewImmutable() MediaType {
	return &MediaTypeImmutable{}
}

// Return a New instance of a MediaTypeImmutable struct
func NewImmutableAsContainer(mutable MediaTypeMutable) MediaType {
	return &MediaTypeImmutable{mutable}
}

// Get the "main" (top-level) type as a string
func (m *MediaTypeImmutable) MainType() string {
	return m.contained.MainType()
}

// Get the "sub" type as a string
func (m *MediaTypeImmutable) SubType() string {
	return m.contained.SubType()
}

// Get the split "sub" type as an array of strings split by the namespace separator
func (m *MediaTypeImmutable) Trees() []string {
	return m.contained.Trees()
}

// Get the prefix of the type's trees
func (m *MediaTypeImmutable) Prefix() string {
	return m.contained.Prefix()
}

// Get the "suffix" of the type as a string
func (m *MediaTypeImmutable) Suffix() string {
	return m.contained.Suffix()
}

// Get the defined parameters of the media type
func (m *MediaTypeImmutable) Parameters() map[string]string {
	return m.contained.Parameters()
}

// Get the normalized full type as a string
func (m *MediaTypeImmutable) FullType() string {
	return m.contained.FullType()
}

// Get a string representation conforming to RFC 2045 and RFC 2616
func (m *MediaTypeImmutable) String() string {
	return m.contained.String()
}

// Get a mutable version of this structure
func (m *MediaTypeImmutable) Mutable() MediaTypeMutable {
	return m.contained
}
