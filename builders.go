/**
 * MediaType
 *
 * Copyright © 2014 Trevor N. Suarez (Rican7)
 */

package mediatype

import (
	"mime/multipart"
	"net/http"
	"net/textproto"
)

/**
 * Constants
 */

const ContentTypeHeader = "Content-Type"

/**
 * Functions
 */

// Build a MediaType from the detected content type of a byte array
func DetectFromBytes(data []byte) (MediaType, error) {
	return Parse(http.DetectContentType(data))
}

// Build a MediaType from a MIMEHeader map
func DetectFromHeader(header textproto.MIMEHeader) (MediaType, error) {
	contentType := header.Get(ContentTypeHeader)

	return Parse(contentType)
}

// Build a MediaType from a FileHeader
func DetectFromFileHeader(fileHeader multipart.FileHeader) (MediaType, error) {
	return DetectFromHeader(fileHeader.Header)
}
