package vocab

import "net/url"

// Indicates the accuracy of position coordinates on a Place objects. Expressed in
// properties of percentage. e.g. "94.0" means "94.0% accurate".
//
// Example 112 (https://www.w3.org/TR/activitystreams-vocabulary/#ex127-jsonld):
//   {
//     "accuracy": 94.5,
//     "latitude": 36.75,
//     "longitude": 119.7667,
//     "name": "Liu Gu Lu Cun, Pingdu, Qingdao, Shandong, China",
//     "type": "Place"
//   }
type AccuracyPropertyInterface interface {
	// Clear ensures no value of this property is set. Calling IsFloat
	// afterwards will return false.
	Clear()
	// Get returns the value of this property. When IsFloat returns false, Get
	// will return any arbitrary value.
	Get() float64
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return any arbitrary value.
	GetIRI() *url.URL
	// HasAny returns true if the value or IRI is set.
	HasAny() bool
	// IsFloat returns true if this property is set and not an IRI.
	IsFloat() bool
	// IsIRI returns true if this property is an IRI.
	IsIRI() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o AccuracyPropertyInterface) bool
	// Name returns the name of this property: "accuracy".
	Name() string
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// Set sets the value of this property. Calling IsFloat afterwards will
	// return true.
	Set(v float64)
	// SetIRI sets the value of this property. Calling IsIRI afterwards will
	// return true.
	SetIRI(v *url.URL)
}