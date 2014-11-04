/**
 * MediaType
 *
 * Copyright © 2014 Trevor N. Suarez (Rican7)
 */

// Package mutable provides a mutable implementation of the MediaType interface
package mutable

import (
	"github.com/rican7/mediatype"
)

// A mutable struct defining the components of a Media Type
type MediaType struct {
	Main   string
	Tree   []string
	Sub    string
	Suf    string
	Params map[string]string
}

// Return a New instance of a MediaType struct
func New() mediatype.MediaType {
	return &MediaType{}
}

// Get the "main" (top-level) type as a string
func (m *MediaType) MainType() string {
	return m.Main
}

// Get the "sub" type as a string
func (m *MediaType) SubType() string {
	return m.Sub
}

// Get the split "sub" type as an array of strings split by the namespace separator
func (m *MediaType) Trees() []string {
	return m.Tree
}

// Get the prefix of the type's trees
func (m *MediaType) Prefix() string {
	if 0 < len(m.Tree) {
		return m.Tree[0]
	}

	return ""
}

// Get the "suffix" of the type as a string
func (m *MediaType) Suffix() string {
	return m.Suf
}

// Get the defined parameters of the media type
func (m *MediaType) Parameters() map[string]string {
	return m.Params
}

// Get the normalized type and sub-type as a string
func (m *MediaType) FullType() string {
	return ""
	// return m.fullType
}

// Get a string representation conforming to RFC 2045 and RFC 2616
func (m *MediaType) String() string {
	return ""
	// return mime.FormatMediaType(m.fullType, m.Params)
}
