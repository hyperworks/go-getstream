package getstream

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Slug struct {
	Slug, ID, Token string
}

func (s Slug) Valid() bool {
	return s.Slug != "" && s.ID != ""
}

// Signature syntax is slug+id _without_ the colon, then the token.
func (s Slug) Signature() string {
	return s.Slug + s.ID + " " + s.Token
}

func (s Slug) WithToken(token string) Slug {
	return Slug{s.Slug, s.ID, token}
}

func (s Slug) MarshalJSON() ([]byte, error) {
	str := s.Slug + ":" + s.ID
	if s.Token != "" {
		str += " " + s.Token
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
			s.Token = parts[1]
			if parts = strings.Split(parts[0], ":"); len(parts) != 2 {
				return false
			}

		default:
			return false
		}

		s.Slug, s.ID = parts[0], parts[1]
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
	result := s.Slug + ":" + s.ID
	if s.Token != "" {
		result += " " + s.Token
	}

	return result
}
