// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RVXW/rNgz9K4K2Rzd2spuXvHVFH4o79Aboug8EQaHITK1OljSRSpcV/u8DZSdOm2zD",
	"NuyhjS1Q5OHh4fGb1L4N3oEjlIs3ibqBVuXHpX+FeOMdpjaQ8Y7PQvQBIhnIETqkp9ZYa/ilBtTRDJHy",
	"+waEan1yJPxW3CwfUUT4NZkItdjsBSjdiFcff7Fe1RNZSPhNtcFy2lVVzKuqmFVVtS4k7QPIhTSO4Bmi",
	"7ArpUvt0uIl/V/gYKBq1A0FebEAoa71WBOeFp8V0frHoqyJ6Mk5HUAiXiwKSaTmrOMQJ40RgEoVORxbF",
	"1se/bH+1LlZz/iumVTGdFzP+v14X0hC0ueFzeMOJilHtZdcV8sC1XKxOpvSRvLFVv3kBTfkqgk7R0P6B",
	"ldDPWQXzGfbXiRp+M9xxA6qGyBlVywl+urpe3l19vv1ZjmjyLdlxUuO2PivGO1Ka+DFFy3mIAi7K8tlQ",
	"kzYT7dvyHmofop9Vs1n5qvzVQKyPpUFMgNywNRpcP4kBwHVQuoGr2aSSxb/JvbF+U7bKuPK7u5vb+4fb",
	"zCvEFr9sHyDujIZ/mJKZMGT52o/XX8TtyfkOIvbKqSbTScWlfACngpEL+U0+KmRQ1GT6S24Rg9KA5ZvD",
	"rjyW4HfVQlfulE2AZVab/rCyKqoWCCKL66Nw7w+ZBTWKBDUw4hQR0KeoQRgUNQTr9/3KZAEwunH8DuWp",
	"6CgmKAYvYQyDvOVC1rBVydKoEqRo3DMr7xI23uHLoP4MCP/8VyjrQgaPWaVsd4oB3dVyIZce6UgZ3uMR",
	"FvLpD3kIy48z6NEA0re+3h+WABz1mxWs0blA+YL9wEasX0fYyoX8qhw9uhwMujxz5+793nPf+QCDd9jv",
	"8ayq/uf67yc40MOuxxkssDti0hoQt8na/YSV/6lH9f7qndspa2rRlzmxz4FLgYmhsyBziul5ikenEjU+",
	"mt+PUZ/Oo0ZlOU9i65Prg+eXUD06tbH5M9LrAibvLDMv2KlZrtasJUxtq+JeLuQDuFqow0dhJO/48Ti0",
	"x+om9cwrOyKU666vFneHbe6NTnKVIfxoiOO1bt39EQAA//8KWBfG5wcAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
