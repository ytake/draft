package payload

/* see https://github.com/blongden/vnd.error
 * media type application/hal+json
 */

// VndRequestURI for expressing a JSON Pointer (RFC6901) to a field in related resource (contained in the 'about' link relation) that this error is relevant for.
type VndRequestURI string
// VndAboutURI the "about" link relation
type VndAboutURI string
