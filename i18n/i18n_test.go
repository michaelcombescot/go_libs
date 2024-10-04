package i18n

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestAll(t *testing.T) {
	lang := LanguageOption{
		Name: "fr",
		Traductions: map[string]any{
			"firstLevel": "found",
			"withArg":    "%s",
			"nested": map[string]any{
				"secondLevel": "found",
				"nested": map[string]any{
					"thirdLevel": "found",
				},
			},
		},
	}

	manager, _ := InitI18nManager(lang)
	traduction := manager.Get(lang.Name)

	t.Run("Get trad from language", func(t *testing.T) {
		cases := []struct {
			key      string
			dflt     string
			args     []any
			expected string
		}{
			{"firstLevel", "", nil, "found"},
			{"nested.secondLevel", "", nil, "found"},
			{"nested.nested.thirdLevel", "", nil, "found"},
			{"withArg", "", []any{"something"}, "something"},
			{"InexistingKey", "%s", []any{"something"}, "something"},
		}

		for _, c := range cases {
			actual := traduction.Get(c.key, c.dflt, c.args...)
			if actual != c.expected {
				t.Errorf("Expected %s, got %s", c.expected, actual)
			}
		}
	})
}
