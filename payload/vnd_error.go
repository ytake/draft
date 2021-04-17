package payload

import "net/http"

// ErrorResponse for vnd error
type ErrorResponse struct {
	// Code Error Code
	Code      int
	AboutPath interface{}
}

// VndError for error response in JSON
type VndError struct {
	// For expressing a human readable message related to the current error which may be displayed to the user of the api.
	Message string `json:"message"`
	// For expressing a JSON Pointer (RFC6901) to a field in related resource (contained in the 'about' link relation) that this error is relevant for.
	//
	// Required: false
	// Example: /v1/resources
	Path VndRequestURI `json:"path"`
	// For expressing a (numeric/alpha/alphanumeric) identifier to refer to the specific error on the server side for logging purposes (i.e. a request number).
	//
	// Required: false
	// Example: 404
	LogRef int `json:"logref"`
	// For error link relations
	//
	// Required: true
	// Example: {}
	Links VndLinks `json:"_links"`
}

// VndLinks Link attributes follow the same semantics as defined in the HAL specification section 5.
type VndLinks struct {
	// The "about" link relation is OPTIONAL.
	//
	// Required: false
	About Link `json:"about"`
}

// NewError for Payload vnd+error response
func NewError() *ErrorResponse {
	return &ErrorResponse{Code: http.StatusNotFound}
}

// VngErrorFactory make vnd error
func (e *ErrorResponse) VngErrorFactory(message string, path VndRequestURI, aboutPath interface{}) *VndError {
	ap, ok := aboutPath.(VndAboutURI)
	if !ok {
		ap = "/"
	}
	return &VndError{
		Message: message,
		Path:    path,
		LogRef:  e.Code,
		Links: VndLinks{
			About: Link{Href: string(ap)},
		},
	}
}
