/**
 * MediaType
 *
 * Copyright © 2014 Trevor N. Suarez (Rican7)
 */

// Package mediatype provides an [Internet] Media Type model and MIME type
// string parser/formatter
package mediatype

import (
	"fmt"
)

/**
 * Constants
 */

const (
	// MainSubSplitCharacter - The character used to split the main and sub-types from a full type string
	MainSubSplitCharacter = "/"

	// SuffixCharacter - The character used to denote a suffix declaration
	SuffixCharacter = "+"

	// TreeSeparatorCharacter - The character used to separate trees
	TreeSeparatorCharacter = "."
)

/**
 * Types
 */

// MediaType is the interface for an [Internet] Media Type model
//
// Each method is a getter for each piece of a Media Type's structure
//
// The `String()` method implements the fmt.Stringer interface to easily
// print out the reconstructed parts of the MediaType model in a standards
// compliant manner, following the `mime.FormatMediaType()` output
type MediaType interface {
	fmt.Stringer

	MainType() string
	SubType() string
	Trees() []string
	Prefix() string
	Suffix() string
	Parameters() map[string]string

	FullType() string
}
