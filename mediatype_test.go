/**
 * MediaType
 *
 * Copyright © 2014 Trevor N. Suarez (Rican7)
 */

// Tests for the mediatype package

package mediatype

import (
	"testing"
)

/**
 * Variables
 */

var validComplexMediaType = "application/vnd.google-earth.kml+xml; charset=utf-8"

var validMediaTypeStrings = []string{
	"example",
	"application/json",
	"application/xhtml+xml",
	"audio/vnd.rn-realaudio",
	"image/vnd.djvu",
}

var invalidMediaTypeStrings = []string{
	"application/vnd/json",
	"invalid/+json.really/sick",
}

/**
 * Helper functions
 */

// Parse our valid media types
func parseValidMediaTypes() (map[string]MediaType, map[string]error) {
	mediaTypes := make(map[string]MediaType)
	errors := make(map[string]error)

	for _, val := range validMediaTypeStrings {
		mt, err := Parse(val)

		if nil != mt {
			mediaTypes[val] = mt
		}

		if nil != err {
			errors[val] = err
		}
	}

	return mediaTypes, errors
}

// Parse our invalid media types
func parseInvalidMediaTypes() (map[string]MediaType, map[string]error) {
	mediaTypes := make(map[string]MediaType)
	errors := make(map[string]error)

	for _, val := range invalidMediaTypeStrings {
		mt, err := Parse(val)

		if nil != mt {
			mediaTypes[val] = mt
		}

		if nil != err {
			errors[val] = err
		}
	}

	return mediaTypes, errors
}

/**
 * Tests functions
 */

func TestParse(t *testing.T) {
	_, validErrs := parseValidMediaTypes()
	invalid, _ := parseInvalidMediaTypes()

	for key, err := range validErrs {
		t.Errorf("Parsing failed for valid '%s' with error '%s'", key, err)
	}

	for key := range invalid {
		t.Errorf("Parsing succeeded for invalid '%s'", key)
	}
}
