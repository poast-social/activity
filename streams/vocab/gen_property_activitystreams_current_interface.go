// Code generated by astool. DO NOT EDIT.

package vocab

import "net/url"

// In a paged Collection, indicates the page that contains the most recently
// updated member items.
//
// Example 74 (https://www.w3.org/TR/activitystreams-vocabulary/#ex71-jsonld):
//   {
//     "current": "http://example.org/collection",
//     "items": [
//       "http://example.org/posts/1",
//       "http://example.org/posts/2",
//       "http://example.org/posts/3"
//     ],
//     "summary": "Sally's blog posts",
//     "totalItems": 3,
//     "type": "Collection"
//   }
//
// Example 75 (https://www.w3.org/TR/activitystreams-vocabulary/#ex72-jsonld):
//   {
//     "current": {
//       "summary": "Most Recent Items",
//       "type": "owl:Class",
//       "url": "http://example.org/collection"
//     },
//     "items": [
//       "http://example.org/posts/1",
//       "http://example.org/posts/2",
//       "http://example.org/posts/3"
//     ],
//     "summary": "Sally's blog posts",
//     "totalItems": 3,
//     "type": "Collection"
//   }
type ActivityStreamsCurrentProperty interface {
	// Clear ensures no value of this property is set. Calling HasAny or any
	// of the 'Is' methods afterwards will return false.
	Clear()
	// GetActivityStreamsCollectionPage returns the value of this property.
	// When IsActivityStreamsCollectionPage returns false,
	// GetActivityStreamsCollectionPage will return an arbitrary value.
	GetActivityStreamsCollectionPage() ActivityStreamsCollectionPage
	// GetActivityStreamsLink returns the value of this property. When
	// IsActivityStreamsLink returns false, GetActivityStreamsLink will
	// return an arbitrary value.
	GetActivityStreamsLink() ActivityStreamsLink
	// GetActivityStreamsMention returns the value of this property. When
	// IsActivityStreamsMention returns false, GetActivityStreamsMention
	// will return an arbitrary value.
	GetActivityStreamsMention() ActivityStreamsMention
	// GetActivityStreamsOrderedCollectionPage returns the value of this
	// property. When IsActivityStreamsOrderedCollectionPage returns
	// false, GetActivityStreamsOrderedCollectionPage will return an
	// arbitrary value.
	GetActivityStreamsOrderedCollectionPage() ActivityStreamsOrderedCollectionPage
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetType returns the value in this property as a Type. Returns nil if
	// the value is not an ActivityStreams type, such as an IRI or another
	// value.
	GetType() Type
	// HasAny returns true if any of the different values is set.
	HasAny() bool
	// IsActivityStreamsCollectionPage returns true if this property has a
	// type of "CollectionPage". When true, use the
	// GetActivityStreamsCollectionPage and
	// SetActivityStreamsCollectionPage methods to access and set this
	// property.
	IsActivityStreamsCollectionPage() bool
	// IsActivityStreamsLink returns true if this property has a type of
	// "Link". When true, use the GetActivityStreamsLink and
	// SetActivityStreamsLink methods to access and set this property.
	IsActivityStreamsLink() bool
	// IsActivityStreamsMention returns true if this property has a type of
	// "Mention". When true, use the GetActivityStreamsMention and
	// SetActivityStreamsMention methods to access and set this property.
	IsActivityStreamsMention() bool
	// IsActivityStreamsOrderedCollectionPage returns true if this property
	// has a type of "OrderedCollectionPage". When true, use the
	// GetActivityStreamsOrderedCollectionPage and
	// SetActivityStreamsOrderedCollectionPage methods to access and set
	// this property.
	IsActivityStreamsOrderedCollectionPage() bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
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
	LessThan(o ActivityStreamsCurrentProperty) bool
	// Name returns the name of this property: "current".
	Name() string
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetActivityStreamsCollectionPage sets the value of this property.
	// Calling IsActivityStreamsCollectionPage afterwards returns true.
	SetActivityStreamsCollectionPage(v ActivityStreamsCollectionPage)
	// SetActivityStreamsLink sets the value of this property. Calling
	// IsActivityStreamsLink afterwards returns true.
	SetActivityStreamsLink(v ActivityStreamsLink)
	// SetActivityStreamsMention sets the value of this property. Calling
	// IsActivityStreamsMention afterwards returns true.
	SetActivityStreamsMention(v ActivityStreamsMention)
	// SetActivityStreamsOrderedCollectionPage sets the value of this
	// property. Calling IsActivityStreamsOrderedCollectionPage afterwards
	// returns true.
	SetActivityStreamsOrderedCollectionPage(v ActivityStreamsOrderedCollectionPage)
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetType attempts to set the property for the arbitrary type. Returns an
	// error if it is not a valid type to set on this property.
	SetType(t Type) error
}
