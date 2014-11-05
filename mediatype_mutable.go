/**
 * MediaType
 *
 * Copyright © 2014 Trevor N. Suarez (Rican7)
 */

// This file provides provides a mutable implementation of the MediaType interface

package mediatype

import (
	"mime"
	"strings"
)

/**
 * Types
 */

// Mutable is a struct defining the components of a Media Type
type Mutable struct {
	Main   string
	Tree   []string
	Sub    string
	Suf    string
	Params map[string]string
}

// NewMutable returns a new empty instance of a Mutable struct
func NewMutable() MediaType {
	return &Mutable{}
}

// MainType returns the "main" (top-level) type as a string
func (m *Mutable) MainType() string {
	return m.Main
}

// SubType returns the "sub" type as a string
func (m *Mutable) SubType() string {
	return m.Sub
}

// Trees returns the split "sub" type as an array of strings split by the namespace separator
func (m *Mutable) Trees() []string {
	return m.Tree
}

// Prefix returns the prefix of the type's trees
func (m *Mutable) Prefix() string {
	if 0 < len(m.Tree) {
		return m.Tree[0]
	}

	return ""
}

// Suffix returns the "suffix" of the type as a string
func (m *Mutable) Suffix() string {
	return m.Suf
}

// Parameters returns the defined parameters of the media type
func (m *Mutable) Parameters() map[string]string {
	return m.Params
}

// FullType returns the normalized full type as a string
func (m *Mutable) FullType() string {
	var fullType string

	fullType += m.Main

	// If there are other pieces
	if 0 < len(m.Tree) || "" != m.Sub || "" != m.Suf {
		fullType += MainSubSplitCharacter
	}

	if 0 < len(m.Tree) {
		fullType += strings.Join(m.Tree, TreeSeparatorCharacter)
		fullType += TreeSeparatorCharacter
	}

	fullType += m.Sub

	if "" != m.Suf {
		fullType += SuffixCharacter + m.Suf
	}

	return fullType
}

// Get a string representation conforming to RFC 2045 and RFC 2616
func (m *Mutable) String() string {
	return mime.FormatMediaType(m.FullType(), m.Params)
}

// Immutable returns an immutable version of this structure
func (m Mutable) Immutable() *Immutable {
	return NewImmutableAsContainer(m).(*Immutable)
}

// Split the full type string into parts and assign those values to our struct
func splitTypes(fullType string) *Mutable {
	var mt Mutable

	// Split the main/sub types
	mainSubSplit := strings.Split(fullType, MainSubSplitCharacter)

	mt.Main = mainSubSplit[0]

	// If we got more than one part, we must have a sub-type
	if 1 < len(mainSubSplit) {
		// Split the remaining main/sub split from a possible suffix
		subSuffixSplit := strings.Split(mainSubSplit[1], SuffixCharacter)

		// If we got more than one part, we must have a suffix
		if 1 < len(subSuffixSplit) {
			mt.Suf = subSuffixSplit[1]
		}

		// Split the sub-type split into the possibly different trees
		treeSubSplit := strings.Split(subSuffixSplit[0], TreeSeparatorCharacter)
		treeSubSplitLength := len(treeSubSplit)

		mt.Sub = treeSubSplit[treeSubSplitLength-1]

		// If we got more than one part, we must have tree definitions
		if 1 < treeSubSplitLength {
			mt.Tree = treeSubSplit[0 : treeSubSplitLength-1]
		}
	}

	// Build from the raw
	return &mt
}
