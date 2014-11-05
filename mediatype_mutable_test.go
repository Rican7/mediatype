/**
 * MediaType
 *
 * Copyright © 2014 Trevor N. Suarez (Rican7)
 */

// Tests for the mediatype mutable structure

package mediatype

import (
	"testing"
)

/**
 * Tests functions
 */

func TestFullType(t *testing.T) {
	mt, err := Parse(validComplexMediaType)

	if nil != err {
		t.Errorf("Parsing failed for valid '%s'", validComplexMediaType)
	} else {

		if mt.FullType() != "application/vnd.google-earth.kml+xml" {
			t.Errorf(mt.FullType())
			t.Errorf("Incorrect full type for %+v", mt)
		}
	}
}

func TestParameters(t *testing.T) {
	mt, err := Parse(validComplexMediaType)

	if nil != err {
		t.Errorf("Parsing failed for valid '%s'", validComplexMediaType)
	} else {
		correctParameters := map[string]string{"charset": "utf-8"}

		for i, tree := range mt.Parameters() {
			if tree != correctParameters[i] {
				t.Errorf("Incorrect parameters for %+v", mt)
			}
		}
	}
}

func TestMainType(t *testing.T) {
	mt, err := Parse(validComplexMediaType)

	if nil != err {
		t.Errorf("Parsing failed for valid '%s'", validComplexMediaType)
	} else {

		if mt.MainType() != "application" {
			t.Errorf("Incorrect main type for %+v", mt)
		}
	}
}

func TestSubType(t *testing.T) {
	mt, err := Parse(validComplexMediaType)

	if nil != err {
		t.Errorf("Parsing failed for valid '%s'", validComplexMediaType)
	} else {

		if mt.SubType() != "kml" {
			t.Errorf("Incorrect sub type for %+v", mt)
		}
	}
}

func TestTrees(t *testing.T) {
	mt, err := Parse(validComplexMediaType)

	if nil != err {
		t.Errorf("Parsing failed for valid '%s'", validComplexMediaType)
	} else {
		correctTrees := []string{"vnd", "google-earth"}

		for i, tree := range mt.Trees() {
			if tree != correctTrees[i] {
				t.Errorf("Incorrect trees for %+v", mt)
			}
		}
	}
}

func TestPrefix(t *testing.T) {
	mt, err := Parse(validComplexMediaType)

	if nil != err {
		t.Errorf("Parsing failed for valid '%s'", validComplexMediaType)
	} else {

		if mt.Prefix() != "vnd" {
			t.Errorf("Incorrect prefix for %+v", mt)
		}
	}
}

func TestSuffix(t *testing.T) {
	mt, err := Parse(validComplexMediaType)

	if nil != err {
		t.Errorf("Parsing failed for valid '%s'", validComplexMediaType)
	} else {

		if mt.Suffix() != "xml" {
			t.Errorf("Incorrect suffix for %+v", mt)
		}
	}
}

func TestString(t *testing.T) {
	mt, err := Parse(validComplexMediaType)

	if nil != err {
		t.Errorf("Parsing failed for valid '%s'", validComplexMediaType)
	} else {

		if mt.String() != validComplexMediaType {
			t.Errorf("Incorrect string for %+v", mt)
		}
	}
}
