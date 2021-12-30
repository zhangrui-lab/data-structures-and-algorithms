package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"sort"
	"testing"
)

func TestRadix_Insert(t *testing.T) {
	radix := NewRadix()
	n := radix.Insert("foo", 1)
	assert.Nilf(t, n, "Expected nil, got: %v", n)

	n = radix.Insert("foo", 2)
	assert.Equalf(t, n.(int), 1, "Expected 1, got: %v", n)

	n = radix.Insert("foobar", 3)
	assert.Nilf(t, n, "Expected nil, got: %v", n)
	n = radix.Insert("fotball", 4)
	assert.Nilf(t, n, "Expected nil, got: %v", n)
}

func TestRadix_Find(t *testing.T) {
	radix := NewRadix()
	maps := map[string]int{
		"foo":      1,
		"foobar":   2,
		"football": 3,
		"foojar":   4,
		"fomang":   5,
	}
	for key, val := range maps {
		n := radix.Insert(key, val)
		assert.Nilf(t, n, "Expected nil, got: %v", n)
	}

	for key, expect := range maps {
		actual := radix.Find(key)
		assert.Equalf(t, actual.(int), expect, "Expected: %s, got: %s", expect, actual)
	}

	// not exists
	actual := radix.Find("askdakhd")
	assert.Nilf(t, actual, "Expected: nil, got: %v", actual)
}

func TestRadix_Remove(t *testing.T) {
	radix := NewRadix()
	maps := map[string]int{
		"foo":      1,
		"foobar":   2,
		"football": 3,
		"foojar":   4,
		"fomang":   5,
		"liusi":    32,
	}
	remove := map[string]interface{}{"foo": nil, "foojar": nil, "liusi": nil}
	for key, val := range maps {
		radix.Insert(key, val)
	}

	assert.Equalf(t, radix.Size(), len(maps), "Expected: %s, got: %s", len(maps), radix.Size())

	for key, _ := range remove {
		expect := maps[key]
		actual := radix.Remove(key)
		assert.Equalf(t, actual.(int), expect, "Expected: %s, got: %s", expect, actual)
	}
	assert.Equalf(t, radix.Size(), len(maps)-len(remove), "Expected: %s, got: %s", len(maps)-len(remove), radix.Size())

	for key, expect := range maps {
		actual := radix.Find(key)
		if _, ok := remove[key]; ok {
			assert.Nilf(t, actual, "Expected: nil, got: %v", actual)
		} else {
			assert.Equalf(t, actual.(int), expect, "Expected: %s, got: %s", expect, actual)
		}
	}
}

func TestRadix_ToMap(t *testing.T) {
	radix := NewRadix()
	maps := map[string]int{
		"foo":    1,
		"foobar": 2,
		"foball": 3,
		"foojar": 4,
		"fomang": 5,
		"liusi":  32,
	}
	for key, val := range maps {
		radix.Insert(key, val)
	}
	actual := radix.ToMap()
	for key, val := range maps {
		k, ok := actual[key]
		if !ok {
			t.Errorf("Expect:%v=%v, actual: not found", key, val)
		}
		if k.(int) != val {
			t.Errorf("Expect:%v, actual: %v", val, k)
		}
	}
}

func TestRadix_Prefix(t *testing.T) {
	radix := NewRadix()
	maps := map[string]int{
		"foo":    1,
		"foobar": 2,
		"foball": 3,
		"foojar": 4,
		"fomang": 5,
		"liusi":  32,
	}
	for key, val := range maps {
		radix.Insert(key, val)
	}

	prefix := "fo"
	actual := radix.Prefix(prefix)
	expect := []string{"foo", "foobar", "foball", "foojar", "fomang"}
	sort.Strings(expect)
	assert.Truef(t, reflect.DeepEqual(expect, actual), "Expect:%v, actual:%v", expect, actual)

	prefix = "foo"
	actual = radix.Prefix(prefix)
	expect = []string{"foo", "foobar", "foojar"}
	sort.Strings(expect)
	fmt.Println()
	assert.Truef(t, reflect.DeepEqual(expect, actual), "Expect:%v, actual:%v", expect, actual)

	prefix = "l"
	actual = radix.Prefix(prefix)
	expect = []string{"liusi"}
	assert.Truef(t, reflect.DeepEqual(expect, actual), "Expect:%v, actual:%v", expect, actual)

	prefix = "sdf"
	actual = radix.Prefix(prefix)
	assert.Nilf(t, actual, "Expect:nil, actual:%v", actual)

	prefix = ""
	actual = radix.Prefix(prefix)
	expect = []string{"foo", "foobar", "foball", "foojar", "fomang", "liusi"}
	sort.Strings(expect)
	assert.Truef(t, reflect.DeepEqual(expect, actual), "Expect:%v, actual:%v", expect, actual)
}
