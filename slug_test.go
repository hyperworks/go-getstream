package getstream_test

import (
	"encoding/json"
	"fmt"
	. "github.com/hyperworks/go-getstream"
	a "github.com/stretchr/testify/assert"
	"testing"
)

var _ json.Marshaler = Slug{}
var _ json.Unmarshaler = &Slug{}
var _ fmt.Stringer = Slug{}

func TestSlug(t *testing.T) {
	slug := NewSlug("test", "123", "token")
	a.Equal(t, "test", slug.Slug())
	a.Equal(t, "123", slug.ID())
	a.Equal(t, "token", slug.Token())
	a.Equal(t, "test:123 token", slug.String())

	slug = NewSlug("test", "456", "")
	a.Equal(t, "test:456", slug.String())
}

func TestSlug_JSON(t *testing.T) {
	marshals := map[string]Slug{
		`"slug:123"`:           NewSlug("slug", "123", ""),
		`"slug:123 signature"`: NewSlug("slug", "123", "signature"),
	}

	for str, slug := range marshals {
		bytes, e := json.Marshal(slug)
		a.NoError(t, e, "failed to marshal slug: "+slug.String())
		a.Equal(t, str, string(bytes))
	}

	unmarshals := map[string]Slug{
		`"slug:123"`:                NewSlug("slug", "123", ""),
		`"slug:123 signature"`:      NewSlug("slug", "123", "signature"),
		`["slug:123", "signature"]`: NewSlug("slug", "123", "signature"),
	}

	for str, slug := range unmarshals {
		result := Slug{}
		e := json.Unmarshal([]byte(str), &result)
		a.NoError(t, e, "failed to unmarshal json: "+str)
		a.Equal(t, slug, result)
	}
}
