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

// ContentTypeHeader - The standard key of a MIMEHeader's content type
const ContentTypeHeader = "Content-Type"

/**
 * Functions
 */

// DetectFromBytes - Build a MediaType from the detected content type of a byte array
func DetectFromBytes(data []byte) (MediaType, error) {
	return Parse(http.DetectContentType(data))
}

// DetectFromHeader - Build a MediaType from a MIMEHeader map
func DetectFromHeader(header textproto.MIMEHeader) (MediaType, error) {
	contentType := header.Get(ContentTypeHeader)

	return Parse(contentType)
}

// DetectFromFileHeader - Build a MediaType from a FileHeader
func DetectFromFileHeader(fileHeader multipart.FileHeader) (MediaType, error) {
	return DetectFromHeader(fileHeader.Header)
}
