// Code generated by astool. DO NOT EDIT.

package propertypreferredusername

import (
	"fmt"
	langstring "github.com/poast-social/activity/streams/values/langString"
	string1 "github.com/poast-social/activity/streams/values/string"
	vocab "github.com/poast-social/activity/streams/vocab"
	"net/url"
)

// ActivityStreamsPreferredUsernameProperty is the functional property
// "preferredUsername". It is permitted to be one of multiple value types. At
// most, one type of value can be present, or none at all. Setting a value
// will clear the other types of values so that only one of the 'Is' methods
// will return true. It is possible to clear all values, so that this property
// is empty.
type ActivityStreamsPreferredUsernameProperty struct {
	xmlschemaStringMember string
	hasStringMember       bool
	rdfLangStringMember   map[string]string
	unknown               interface{}
	iri                   *url.URL
	alias                 string
}

// DeserializePreferredUsernameProperty creates a "preferredUsername" property
// from an interface representation that has been unmarshalled from a text or
// binary format.
func DeserializePreferredUsernameProperty(m map[string]interface{}, aliasMap map[string]string) (*ActivityStreamsPreferredUsernameProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	propName := "preferredUsername"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "preferredUsername")
	}
	i, ok := m[propName]
	if !ok {
		// Attempt to find the map instead.
		i, ok = m[propName+"Map"]
	}
	if ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
			if err == nil && len(u.Scheme) > 0 {
				this := &ActivityStreamsPreferredUsernameProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if v, err := string1.DeserializeString(i); err == nil {
			this := &ActivityStreamsPreferredUsernameProperty{
				alias:                 alias,
				hasStringMember:       true,
				xmlschemaStringMember: v,
			}
			return this, nil
		} else if v, err := langstring.DeserializeLangString(i); err == nil {
			this := &ActivityStreamsPreferredUsernameProperty{
				alias:               alias,
				rdfLangStringMember: v,
			}
			return this, nil
		}
		this := &ActivityStreamsPreferredUsernameProperty{
			alias:   alias,
			unknown: i,
		}
		return this, nil
	}
	return nil, nil
}

// NewActivityStreamsPreferredUsernameProperty creates a new preferredUsername
// property.
func NewActivityStreamsPreferredUsernameProperty() *ActivityStreamsPreferredUsernameProperty {
	return &ActivityStreamsPreferredUsernameProperty{alias: ""}
}

// Clear ensures no value and no language map for this property is set. Calling
// HasAny or any of the 'Is' methods afterwards will return false.
func (this *ActivityStreamsPreferredUsernameProperty) Clear() {
	this.hasStringMember = false
	this.rdfLangStringMember = nil
	this.unknown = nil
	this.iri = nil
	this.rdfLangStringMember = nil
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this ActivityStreamsPreferredUsernameProperty) GetIRI() *url.URL {
	return this.iri
}

// GetLanguage returns the value for the specified BCP47 language code, or an
// empty string if it is either not a language map or no value is present.
func (this ActivityStreamsPreferredUsernameProperty) GetLanguage(bcp47 string) string {
	if this.rdfLangStringMember == nil {
		return ""
	} else if v, ok := this.rdfLangStringMember[bcp47]; ok {
		return v
	} else {
		return ""
	}
}

// GetRDFLangString returns the value of this property. When IsRDFLangString
// returns false, GetRDFLangString will return an arbitrary value.
func (this ActivityStreamsPreferredUsernameProperty) GetRDFLangString() map[string]string {
	return this.rdfLangStringMember
}

// GetXMLSchemaString returns the value of this property. When IsXMLSchemaString
// returns false, GetXMLSchemaString will return an arbitrary value.
func (this ActivityStreamsPreferredUsernameProperty) GetXMLSchemaString() string {
	return this.xmlschemaStringMember
}

// HasAny returns true if any of the values are set, except for the natural
// language map. When true, the specific has, getter, and setter methods may
// be used to determine what kind of value there is to access and set this
// property. To determine if the property was set as a natural language map,
// use the IsRDFLangString method instead.
func (this ActivityStreamsPreferredUsernameProperty) HasAny() bool {
	return this.IsXMLSchemaString() ||
		this.IsRDFLangString() ||
		this.iri != nil
}

// HasLanguage returns true if the natural language map has an entry for the
// specified BCP47 language code.
func (this ActivityStreamsPreferredUsernameProperty) HasLanguage(bcp47 string) bool {
	if this.rdfLangStringMember == nil {
		return false
	} else {
		_, ok := this.rdfLangStringMember[bcp47]
		return ok
	}
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this ActivityStreamsPreferredUsernameProperty) IsIRI() bool {
	return this.iri != nil
}

// IsRDFLangString returns true if this property has a type of "langString". When
// true, use the GetRDFLangString and SetRDFLangString methods to access and
// set this property.. To determine if the property was set as a natural
// language map, use the IsRDFLangString method instead.
func (this ActivityStreamsPreferredUsernameProperty) IsRDFLangString() bool {
	return this.rdfLangStringMember != nil
}

// IsXMLSchemaString returns true if this property has a type of "string". When
// true, use the GetXMLSchemaString and SetXMLSchemaString methods to access
// and set this property.. To determine if the property was set as a natural
// language map, use the IsRDFLangString method instead.
func (this ActivityStreamsPreferredUsernameProperty) IsXMLSchemaString() bool {
	return this.hasStringMember
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsPreferredUsernameProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/ns/activitystreams": this.alias}
	var child map[string]string

	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this ActivityStreamsPreferredUsernameProperty) KindIndex() int {
	if this.IsXMLSchemaString() {
		return 0
	}
	if this.IsRDFLangString() {
		return 1
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ActivityStreamsPreferredUsernameProperty) LessThan(o vocab.ActivityStreamsPreferredUsernameProperty) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsXMLSchemaString() {
		return string1.LessString(this.GetXMLSchemaString(), o.GetXMLSchemaString())
	} else if this.IsRDFLangString() {
		return langstring.LessLangString(this.GetRDFLangString(), o.GetRDFLangString())
	} else if this.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	}
	return false
}

// Name returns the name of this property: "preferredUsername".
func (this ActivityStreamsPreferredUsernameProperty) Name() string {
	if this.IsRDFLangString() {
		return "preferredUsernameMap"
	} else {
		return "preferredUsername"
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsPreferredUsernameProperty) Serialize() (interface{}, error) {
	if this.IsXMLSchemaString() {
		return string1.SerializeString(this.GetXMLSchemaString())
	} else if this.IsRDFLangString() {
		return langstring.SerializeLangString(this.GetRDFLangString())
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *ActivityStreamsPreferredUsernameProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}

// SetLanguage sets the value for the specified BCP47 language code.
func (this *ActivityStreamsPreferredUsernameProperty) SetLanguage(bcp47, value string) {
	this.hasStringMember = false
	this.rdfLangStringMember = nil
	this.unknown = nil
	this.iri = nil
	if this.rdfLangStringMember == nil {
		this.rdfLangStringMember = make(map[string]string)
	}
	this.rdfLangStringMember[bcp47] = value
}

// SetRDFLangString sets the value of this property and clears the natural
// language map. Calling IsRDFLangString afterwards will return true. Calling
// IsRDFLangString afterwards returns false.
func (this *ActivityStreamsPreferredUsernameProperty) SetRDFLangString(v map[string]string) {
	this.Clear()
	this.rdfLangStringMember = v
}

// SetXMLSchemaString sets the value of this property and clears the natural
// language map. Calling IsXMLSchemaString afterwards will return true.
// Calling IsRDFLangString afterwards returns false.
func (this *ActivityStreamsPreferredUsernameProperty) SetXMLSchemaString(v string) {
	this.Clear()
	this.xmlschemaStringMember = v
	this.hasStringMember = true
}
