# mediatype

[![GoDoc](https://godoc.org/github.com/Rican7/mediatype?status.png)](https://godoc.org/github.com/Rican7/mediatype) [![Build Status](https://travis-ci.org/Rican7/mediatype.svg?branch=master)](https://travis-ci.org/Rican7/mediatype)

An [Internet] Media Type and MIME string parser and modeler for Go

This library expands on the builtin "mime" package to provide a model for an ["Internet Media Type"][wiki-internet-media-type] normally specified by `Content-Type` HTTP and MIME header fields, and as specified in [RFC 2045][ietf-rfc-2045].


## Install

`go get github.com/Rican7/mediatype`


## Example

```go
contentTypeString := "application/vnd.google-earth.kml+xml; charset=utf-8"

mediaType, _ := mediatype.Parse(contentTypeString)

mediaType.FullType()   // "application/vnd.google-earth.kml+xml"
mediaType.Parameters() // ["charset": "utf-8"]
mediaType.MainType()   // "application"
mediaType.SubType()    // "kml"
mediaType.Trees()      // ["vnd", "google-earth"]
mediaType.Prefix()     // "vnd"
mediaType.Suffix()     // "xml"
mediaType.String()     // "application/vnd.google-earth.kml+xml; charset=utf-8"
```




[wiki-internet-media-type]: http://en.wikipedia.org/wiki/Internet_media_type
[ietf-rfc-2045]: https://www.ietf.org/rfc/rfc2045.txt
