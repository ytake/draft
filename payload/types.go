package payload

// Link
// For Link attributes follow the same semantics as defined in the HAL specification section 5. For completeness, the required elements are all represented below.
type Link struct {
	// The "href" property is REQUIRED.
	//
	// Required: true
	// Example:  http://path.to/user/resource/1
	Href string `json:"href"`
}

// Ping for health check
type Ping struct {
	// server status
	//
	// Required: true
	// Example: ok
	Status string `json:"status"`
}
