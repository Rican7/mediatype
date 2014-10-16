/**
 * MediaType
 *
 * Copyright © 2014 Trevor N. Suarez (Rican7)
 */

package mediatype

import (
	"mime"
	"strings"
)

/**
 * Constants
 */

const (
	// The character used to split the main and sub-types from a full type string
	MainSubSplitCharacter = "/"

	// The character used to denote a suffix declaration
	SuffixCharacter = "+"

	// The character used to separate trees
	TreeSeparatorCharacter = "."
)

/**
 * Types
 */

type MediaType interface {
	FullType() string
	Parameters() map[string]string
	MainType() string
	SubType() string
	Trees() []string
	Prefix() string
	Suffix() string
	String() string
}

type mediaType struct {
	fullType string
	params   map[string]string

	mainType string
	trees    []string
	subType  string
	suffix   string
}

/**
 * Functions
 */

// Parse a raw media type string into a MediaType interface compatible struct
func Parse(raw string) (MediaType, error) {
	normalized, params, err := mime.ParseMediaType(raw)

	if nil != err {
		return nil, err
	}

	// Build from the raw
	mediaType := &mediaType{
		fullType: normalized,
		params:   params,
	}

	mediaType.splitTypes()

	return mediaType, nil
}

// Get the normalized type and sub-type as a string
func (m *mediaType) FullType() string {
	return m.fullType
}

// Get the defined parameters of the media type
func (m *mediaType) Parameters() map[string]string {
	return m.params
}

// Get the "main" (top-level) type as a string
func (m *mediaType) MainType() string {
	return m.mainType
}

// Get the "sub" type as a string
func (m *mediaType) SubType() string {
	return m.subType
}

// Get the split "sub" type as an array of strings split by the namespace separator
func (m *mediaType) Trees() []string {
	return m.trees
}

// Get the prefix of the type's trees
func (m *mediaType) Prefix() string {
	if 0 < len(m.trees) {
		return m.trees[0]
	}

	return ""
}

// Get the "suffix" of the type as a string
func (m *mediaType) Suffix() string {
	return m.suffix
}

// Get a string representation conforming to RFC 2045 and RFC 2616
func (m *mediaType) String() string {
	return mime.FormatMediaType(m.fullType, m.params)
}

// Split the full type string into parts and assign those values to our struct
func (m *mediaType) splitTypes() {
	// Split the main/sub types
	mainSubSplit := strings.Split(m.fullType, MainSubSplitCharacter)

	m.mainType = mainSubSplit[0]

	// If we got more than one part, we must have a sub-type
	if 1 < len(mainSubSplit) {
		// Split the remaining main/sub split from a possible suffix
		subSuffixSplit := strings.Split(mainSubSplit[1], SuffixCharacter)

		// If we got more than one part, we must have a suffix
		if 1 < len(subSuffixSplit) {
			m.suffix = subSuffixSplit[1]
		}

		// Split the sub-type split into the possibly different trees
		treeSubSplit := strings.Split(subSuffixSplit[0], TreeSeparatorCharacter)
		treeSubSplitLength := len(treeSubSplit)

		m.subType = treeSubSplit[treeSubSplitLength-1]

		// If we got more than one part, we must have tree definitions
		if 1 < treeSubSplitLength {
			m.trees = treeSubSplit[0 : treeSubSplitLength-1]
		}
	}
}
