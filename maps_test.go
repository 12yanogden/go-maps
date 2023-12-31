package maps

import (
	"strconv"
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

func TestMapEquality(t *testing.T) {
	m1 := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	m2 := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	expected := true
	actual := Equals(m1, m2)

	if !actual {
		t.Fatalf("\nExpected: '" + strconv.FormatBool(expected) + "', they are equal\nActual: '" + strconv.FormatBool(actual) + "', they are not equal")
	}
}

func TestMapInequalityByLength(t *testing.T) {
	m1 := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	m2 := map[string]string{"key1": "value1", "key2": "value2"}
	expected := false
	actual := Equals(m1, m2)

	if actual {
		t.Fatalf("\nExpected: '" + strconv.FormatBool(expected) + "', they are inequal\nActual: '" + strconv.FormatBool(actual) + "', they are equal")
	}
}

func TestMapInequalityByKeys(t *testing.T) {
	m1 := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	m2 := map[string]string{"key1": "value1", "key2": "value2", "key": "value3"}
	expected := false
	actual := Equals(m1, m2)

	if actual {
		t.Fatalf("\nExpected: '" + strconv.FormatBool(expected) + "', they are inequal\nActual: '" + strconv.FormatBool(actual) + "', they are equal")
	}
}

func TestMapInequalityByValues(t *testing.T) {
	m1 := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	m2 := map[string]string{"key1": "value1", "key2": "value2", "key3": "value"}
	expected := false
	actual := Equals(m1, m2)

	if actual {
		t.Fatalf("\nExpected: '" + strconv.FormatBool(expected) + "', they are inequal\nActual: '" + strconv.FormatBool(actual) + "', they are equal")
	}
}
