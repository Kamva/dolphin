package dolphin

import "fmt"

// Dictionary is a map of string to string that represent a key to its translation
type Dictionary = map[string]string

// Translator is a struct, containing a dictionary
type Translator struct {
	dictionary map[string]string
}

// Translate convert given key
func (t Translator) Translate(key string, params ...interface{}) string {
	if translation, ok := t.dictionary[key]; ok {
		return fmt.Sprintf(translation, params...)
	}

	return key
}

// NewTranslator generates and returns Translator object
func NewTranslator(dictionary ...Dictionary) Translator {
	return Translator{
		dictionary: mergeDictionary(dictionary...),
	}
}

func mergeDictionary(dictionaries ...Dictionary) Dictionary {
	var dictionary = make(Dictionary)
	for _, dict := range dictionaries {
		for key, value := range dict {
			dictionary[key] = value
		}
	}

	return dictionary
}
