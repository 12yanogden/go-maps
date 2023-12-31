package maps

import (
	"testing"
)

func TestHowKey(t *testing.T) {
	mapp := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	expected := true
	actual := HasKey[string](mapp, "key2")

	if (expected != actual) {
		t.Fatalf("\nExpected:\t%t, key was found\nActual:\t\t%t, key was NOT found\n", expected, actual)
	}
}