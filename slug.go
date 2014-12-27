package getstream

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Slug struct {
	slug, id, token string
}

func NewSlug(slug, id, token string) Slug {
	return Slug{slug, id, token}
}

func (s Slug) Slug() string  { return s.slug }
func (s Slug) ID() string    { return s.id }
func (s Slug) Token() string { return s.token }

func (s Slug) MarshalJSON() ([]byte, error) {
	str := s.slug + ":" + s.id
	if s.token != "" {
		str += " " + s.token
	}

	return json.Marshal(str)
}

func (s *Slug) UnmarshalJSON(bytes []byte) error {
	var raw interface{}
	if e := json.Unmarshal(bytes, &raw); e != nil {
		return e
	}

	parseArr := func(parts []string) bool {
		switch len(parts) {
		case 1:
			if parts = strings.Split(parts[0], ":"); len(parts) != 2 {
				return false
			}

		case 2:
			s.token = parts[1]
			if parts = strings.Split(parts[0], ":"); len(parts) != 2 {
				return false
			}

		default:
			return false
		}

		s.slug, s.id = parts[0], parts[1]
		return true
	}

Outer:
	switch r := raw.(type) {
	case string:
		if parseArr(strings.Split(r, " ")) {
			return nil
		}

	case []interface{}:
		parts := make([]string, len(r), cap(r))
		for i, iface := range r {
			if str, ok := iface.(string); ok {
				parts[i] = str
			} else {
				break Outer
			}
		}

		if parseArr(parts) {
			return nil
		}
	}

	return fmt.Errorf("cannot parse slug from: %#v", raw)
}

func (s Slug) String() string {
	result := s.slug + ":" + s.id
	if s.token != "" {
		result += " " + s.token
	}

	return result
}
