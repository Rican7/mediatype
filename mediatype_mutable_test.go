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

func TestFullType(t *testing.T) {
	mt, err := Parse(validComplexMediaType)

	if nil != err {
		t.Errorf("Parsing failed for valid '%s'", validComplexMediaType)
	} else {

		if mt.FullType() != "application/vnd.google-earth.kml+xml" {
			t.Errorf("Incorrect full type for %+v", mt)
		}
	}
}

func TestFullTypeWithSimpleMain(t *testing.T) {
	mt, err := Parse("application")

	if nil != err {
		t.Errorf("Parsing failed for valid '%s'", validComplexMediaType)
	} else {

		if mt.FullType() != "application" {
			t.Errorf("Incorrect full type for %+v", mt)
		}
	}
}

func TestFullTypeWithSimpleMainAndSub(t *testing.T) {
	mt, err := Parse("application/json")

	if nil != err {
		t.Errorf("Parsing failed for valid '%s'", validComplexMediaType)
	} else {

		if mt.FullType() != "application/json" {
			t.Errorf("Incorrect full type for %+v", mt)
		}
	}
}

func TestFullTypeWithSimpleMainAndSubAndTrees(t *testing.T) {
	mt, err := Parse("image/vnd.djvu")

	if nil != err {
		t.Errorf("Parsing failed for valid '%s'", validComplexMediaType)
	} else {

		if mt.FullType() != "image/vnd.djvu" {
			t.Errorf("Incorrect full type for %+v", mt)
		}
	}
}

func TestFullTypeWithSimpleMainAndSubAndPrefix(t *testing.T) {
	mt, err := Parse("application/xhtml+xml")

	if nil != err {
		t.Errorf("Parsing failed for valid '%s'", validComplexMediaType)
	} else {

		if mt.FullType() != "application/xhtml+xml" {
			t.Errorf("Incorrect full type for %+v", mt)
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

func TestImmutable(t *testing.T) {
	mt, err := Parse(validComplexMediaType)

	if nil != err {
		t.Errorf("Parsing failed for valid '%s'", validComplexMediaType)
	} else {
		mutable := mt.(*MediaTypeMutable)
		immutable := mutable.Immutable()

		if mutable.String() != immutable.String() {
			t.Errorf("Immutable doesn't match for %+v", mutable)
		}

		mutable.Main = "asd"

		if mutable.String() == immutable.String() {
			t.Errorf("Immutable was mutated for %+v", mutable)
		}
	}
}
