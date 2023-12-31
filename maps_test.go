package maps

import (
	"testing"
)

func TestHasKey(t *testing.T) {
	mapp := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	expected := true
	actual := HasKey[string](mapp, "key2")

	if expected != actual {
		t.Fatalf("\nExpected:\t%t, key was found\nActual:\t\t%t, key was NOT found\n", expected, actual)
	}
}

func TestNotHasKey(t *testing.T) {
	mapp := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	expected := false
	actual := HasKey[string](mapp, "key4")

	if expected != actual {
		t.Fatalf("\nExpected:\t%t, key was NOT found\nActual:\t\t%t, key was found\n", expected, actual)
	}
}
