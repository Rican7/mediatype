/**
 * MediaType
 *
 * Copyright © 2014 Trevor N. Suarez (Rican7)
 */

// Package mediatype provides an [Internet] Media Type model and MIME type
// string parser/formatter
package mediatype

import (
	"mime"
	"strings"
)

/**
 * Types
 */

// An immutable struct of a Media Type
type mediaType struct {
	main   string
	tree   []string
	sub    string
	suf    string
	params map[string]string
}

/**
 * Functions
 */

// Parse a raw media type string into a MediaType interface compatible struct
func Parse(raw string) (MediaType, error) {
	normalizedFullType, params, err := mime.ParseMediaType(raw)

	if nil != err {
		return nil, err
	}

	mediaType := splitTypes(normalizedFullType)
	mediaType.params = params

	return mediaType, nil
}

// Get the "main" (top-level) type as a string
func (m *mediaType) MainType() string {
	return m.main
}

// Get the "sub" type as a string
func (m *mediaType) SubType() string {
	return m.sub
}

// Get the split "sub" type as an array of strings split by the namespace separator
func (m *mediaType) Trees() []string {
	return m.tree
}

// Get the prefix of the type's trees
func (m *mediaType) Prefix() string {
	if 0 < len(m.tree) {
		return m.tree[0]
	}

	return ""
}

// Get the "suffix" of the type as a string
func (m *mediaType) Suffix() string {
	return m.suf
}

// Get the defined parameters of the media type
func (m *mediaType) Parameters() map[string]string {
	return m.params
}

// Get the normalized type and sub-type as a string
func (m *mediaType) FullType() string {
	return ""
	// return m.fullType
}

// Get a string representation conforming to RFC 2045 and RFC 2616
func (m *mediaType) String() string {
	return ""
	// return mime.FormatMediaType(m.fullType, m.params)
}

// Split the full type string into parts and assign those values to our struct
func splitTypes(fullType string) *mediaType {
	var mt *mediaType

	// Split the main/sub types
	mainSubSplit := strings.Split(fullType, MainSubSplitCharacter)

	mt.main = mainSubSplit[0]

	// If we got more than one part, we must have a sub-type
	if 1 < len(mainSubSplit) {
		// Split the remaining main/sub split from a possible suffix
		subSuffixSplit := strings.Split(mainSubSplit[1], SuffixCharacter)

		// If we got more than one part, we must have a suffix
		if 1 < len(subSuffixSplit) {
			mt.suf = subSuffixSplit[1]
		}

		// Split the sub-type split into the possibly different trees
		treeSubSplit := strings.Split(subSuffixSplit[0], TreeSeparatorCharacter)
		treeSubSplitLength := len(treeSubSplit)

		mt.sub = treeSubSplit[treeSubSplitLength-1]

		// If we got more than one part, we must have tree definitions
		if 1 < treeSubSplitLength {
			mt.tree = treeSubSplit[0 : treeSubSplitLength-1]
		}
	}

	// Build from the raw
	return mt
}
