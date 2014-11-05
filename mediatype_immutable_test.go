/**
 * MediaType
 *
 * Copyright © 2014 Trevor N. Suarez (Rican7)
 */

// Tests for the mediatype immutable structure

package mediatype

import (
	"testing"
)

/**
 * Helper functions
 */

func parseAndConvertToImmutable(raw string) (*MediaTypeImmutable, error) {
	mt, err := Parse(validComplexMediaType)

	if nil != err {
		return nil, err
	}

	return mt.(*MediaTypeMutable).Immutable(), nil
}

/**
 * Tests functions
 */

func TestNewImmutable(t *testing.T) {
	mt := NewImmutable()

	if _, ok := mt.(MediaType); !ok {
		t.Errorf("%+v doesn't satisfy the MediaType interface", mt)
	}

	if _, ok := mt.(*MediaTypeImmutable); !ok {
		t.Errorf("MediaType %+v isn't an immutable", mt)
	}
}

func TestNewImmutableAsContainer(t *testing.T) {
	mutable := &MediaTypeMutable{Main: "application", Sub: "doge"}

	mt := NewImmutableAsContainer(*mutable)

	if _, ok := mt.(MediaType); !ok {
		t.Errorf("%+v doesn't satisfy the MediaType interface", mt)
	}

	if _, ok := mt.(*MediaTypeImmutable); !ok {
		t.Errorf("MediaType %+v isn't an immutable", mt)
	}

	mutable.Sub = "nope"

	if mt.String() == mutable.String() {
		t.Errorf("Immutable was mutated for %+v", mt)
	}
}

func TestMutable(t *testing.T) {
	mt, err := parseAndConvertToImmutable(validComplexMediaType)

	if nil != err {
		t.Errorf("Parsing failed for valid '%s'", validComplexMediaType)
	} else {
		mutable := mt.Mutable()

		if mt.String() != mutable.String() {
			t.Errorf("Mutable doesn't match for %+v", mt)
		}

		mutable.Main = "asd"

		if mt.String() == mutable.String() {
			t.Errorf("Immutable was mutated for %+v", mt)
		}
	}
}
