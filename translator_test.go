package dolphin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dictionary1 = Dictionary{
	"key1": "This is translated value for key1",
	"key2": "This is translated value for key2 with string param [%s]",
	"key3": "This should be overwritten",
	"key4": "This also should be overwritten",
	"key5": "This is translated value for key5 with param [%s] and [%v]",
}

var dictionary2 = Dictionary{
	"key3": "This is translated value for key3 with numeric param [%n]",
	"key4": "This is translated value for key4 with generic param [%v]",
}

var dictionary = Dictionary{
	"key1": "This is translated value for key1",
	"key2": "This is translated value for key2 with string param [%s]",
	"key3": "This is translated value for key3 with numeric param [%n]",
	"key4": "This is translated value for key4 with generic param [%v]",
	"key5": "This is translated value for key5 with param [%s] and [%v]",
}

var strParam = "string param"
var numParam = 677
var genericParam = false

var values = []string{
	"This is translated value for key1",
	"This is translated value for key2 with string param [string param]",
	"This is translated value for key3 with numeric param [677]",
	"This is translated value for key4 with generic param [false]",
	"This is translated value for key5 with param [string param] and [false]",
}

func TestNewTranslator(t *testing.T) {
	translator := NewTranslator(dictionary1, dictionary2)

	assert.Equal(t, translator.dictionary, dictionary)
}

func TestTranslator_Translate(t *testing.T) {
	t.Run("withoutParamsSuccess", func(t *testing.T) {
		translator := NewTranslator(dictionary1, dictionary2)
		t1 := translator.Translate("key1")

		assert.Equal(t, values[0], t1)
	})

	t.Run("withoutParamsFailed", func(t *testing.T) {
		translator := NewTranslator(dictionary1, dictionary2)
		t1 := translator.Translate("key1", strParam)

		assert.NotEqual(t, values[0], t1)
	})

	t.Run("withParamsSuccess", func(t *testing.T) {
		translator := NewTranslator(dictionary1, dictionary2)
		t2 := translator.Translate("key2")
		t3 := translator.Translate("key3")
		t4 := translator.Translate("key4")
		t5 := translator.Translate("key5")

		assert.NotEqual(t, values[0], t2)
		assert.NotEqual(t, values[0], t3)
		assert.NotEqual(t, values[0], t4)
		assert.NotEqual(t, values[0], t5)
	})

	t.Run("withParamsFailed", func(t *testing.T) {
		translator := NewTranslator(dictionary1, dictionary2)
		t2 := translator.Translate("key2", strParam)
		t3 := translator.Translate("key3", numParam)
		t4 := translator.Translate("key4", genericParam)
		t5 := translator.Translate("key5", strParam, genericParam)

		assert.NotEqual(t, values[0], t2)
		assert.NotEqual(t, values[0], t3)
		assert.NotEqual(t, values[0], t4)
		assert.NotEqual(t, values[0], t5)
	})
}
