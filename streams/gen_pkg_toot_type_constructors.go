// Code generated by astool. DO NOT EDIT.

package streams

import (
	typeemoji "github.com/poast-social/activity/streams/impl/toot/type_emoji"
	typeidentityproof "github.com/poast-social/activity/streams/impl/toot/type_identityproof"
	vocab "github.com/poast-social/activity/streams/vocab"
)

// NewTootEmoji creates a new TootEmoji
func NewTootEmoji() vocab.TootEmoji {
	return typeemoji.NewTootEmoji()
}

// NewTootIdentityProof creates a new TootIdentityProof
func NewTootIdentityProof() vocab.TootIdentityProof {
	return typeidentityproof.NewTootIdentityProof()
}
