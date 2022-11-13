// +build generate
//go:generate go run ./astool -spec astool/activitystreams.jsonld -spec astool/security-v1.jsonld -spec astool/toot.jsonld -spec astool/forgefed.jsonld -path github.com/poast-social/activity ./streams

package activity
