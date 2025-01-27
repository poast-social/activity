// Code generated by astool. DO NOT EDIT.

package streams

import (
	propertyblurhash "github.com/poast-social/activity/streams/impl/toot/property_blurhash"
	propertydiscoverable "github.com/poast-social/activity/streams/impl/toot/property_discoverable"
	propertyfeatured "github.com/poast-social/activity/streams/impl/toot/property_featured"
	propertysignaturealgorithm "github.com/poast-social/activity/streams/impl/toot/property_signaturealgorithm"
	propertysignaturevalue "github.com/poast-social/activity/streams/impl/toot/property_signaturevalue"
	propertyvoterscount "github.com/poast-social/activity/streams/impl/toot/property_voterscount"
	vocab "github.com/poast-social/activity/streams/vocab"
)

// NewTootTootBlurhashProperty creates a new TootBlurhashProperty
func NewTootBlurhashProperty() vocab.TootBlurhashProperty {
	return propertyblurhash.NewTootBlurhashProperty()
}

// NewTootTootDiscoverableProperty creates a new TootDiscoverableProperty
func NewTootDiscoverableProperty() vocab.TootDiscoverableProperty {
	return propertydiscoverable.NewTootDiscoverableProperty()
}

// NewTootTootFeaturedProperty creates a new TootFeaturedProperty
func NewTootFeaturedProperty() vocab.TootFeaturedProperty {
	return propertyfeatured.NewTootFeaturedProperty()
}

// NewTootTootSignatureAlgorithmProperty creates a new
// TootSignatureAlgorithmProperty
func NewTootSignatureAlgorithmProperty() vocab.TootSignatureAlgorithmProperty {
	return propertysignaturealgorithm.NewTootSignatureAlgorithmProperty()
}

// NewTootTootSignatureValueProperty creates a new TootSignatureValueProperty
func NewTootSignatureValueProperty() vocab.TootSignatureValueProperty {
	return propertysignaturevalue.NewTootSignatureValueProperty()
}

// NewTootTootVotersCountProperty creates a new TootVotersCountProperty
func NewTootVotersCountProperty() vocab.TootVotersCountProperty {
	return propertyvoterscount.NewTootVotersCountProperty()
}
