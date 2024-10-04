package i18n

import (
	"fmt"
	"log/slog"
)

type ctxKey string

const (
	I18nKey = ctxKey("i18n")
)

// this type is what is effectively used by the consumer of the library once the LanguageOption is flattened
// all retrieved data on the side of the consumer will use this type
type Language map[string]string

func (l Language) Get(key string, dflt string, args ...any) string {
	val, ok := l[key]
	if !ok {
		val = dflt
	}

	if len(args) != 0 {
		return fmt.Sprintf(val, args...)
	} else {
		return val
	}
}

// this type is here for initializing the i18nManager
// trad is done from a map[string]any which is then flattened into a Language, which is a map[string]string with a special access method Get
type LanguageOption struct {
	Name        string
	Traductions map[string]any
}

// the manager is an object which will memoise all possible traductions given by the consumer as LanguageOptions
// consumer should initialize it with InitI18nManager
type I18nManager map[string]Language

func InitI18nManager(languages ...LanguageOption) (I18nManager, error) {
	if len(languages) == 0 {
		return nil, fmt.Errorf("[i18n.InitI18nManager] no languages provided, please provide at least one")
	}

	manager := make(map[string]Language)

	for _, l := range languages {
		if l.Traductions == nil {
			return nil, fmt.Errorf("[i18n.InitI18nManager] traduction has not been initialized for for LanguageOption with name '%s'", l.Name)
		}

		manager[l.Name] = flatten(l.Traductions, ".")
	}

	return manager, nil
}

// If no language is found with the key, a random language is chosen to prevent a panic/error.
func (i I18nManager) Get(lang string) Language {
	language, ok := i[lang]
	if !ok {
		slog.Warn("[i18n.Get] language not found, defaulting to empty language", slog.String("missing language key", lang))
		return Language{}
	}
	return language
}
