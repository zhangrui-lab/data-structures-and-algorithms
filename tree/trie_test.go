package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"sort"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := NewTrie()
	n := trie.Insert("foo", 1)
	assert.Nilf(t, n, "Expected nil, got: %v", n)

	n = trie.Insert("foo", 2)
	assert.Equalf(t, n.(int), 1, "Expected 1, got: %v", n)

	n = trie.Insert("foobar", 3)
	assert.Nilf(t, n, "Expected nil, got: %v", n)
	n = trie.Insert("football", 4)
	assert.Nilf(t, n, "Expected nil, got: %v", n)
}

func TestTrie_Find(t *testing.T) {
	trie := NewTrie()
	maps := map[string]int{
		"foo":      1,
		"foobar":   2,
		"football": 3,
		"foojar":   4,
		"fomang":   5,
	}
	for key, val := range maps {
		n := trie.Insert(key, val)
		assert.Nilf(t, n, "Expected nil, got: %v", n)
	}

	for key, expect := range maps {
		actual := trie.Find(key)
		assert.Equalf(t, actual.(int), expect, "Expected: %s, got: %s", expect, actual)
	}

	// not exists
	actual := trie.Find("askdakhd")
	assert.Nilf(t, actual, "Expected: nil, got: %v", actual)
}

func TestTrie_Remove(t *testing.T) {
	trie := NewTrie()
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
		trie.Insert(key, val)
	}

	for key, _ := range remove {
		expect := maps[key]
		actual := trie.Remove(key)
		assert.Equalf(t, actual.(int), expect, "Expected: %s, got: %s", expect, actual)
	}

	for key, expect := range maps {
		actual := trie.Find(key)
		if _, ok := remove[key]; ok {
			assert.Nilf(t, actual, "Expected: nil, got: %v", actual)
		} else {
			assert.Equalf(t, actual.(int), expect, "Expected: %s, got: %s", expect, actual)
		}
	}
}

func TestTrie_ToMap(t *testing.T) {
	trie := NewTrie()
	maps := map[string]int{
		"foo":    1,
		"foobar": 2,
		"foball": 3,
		"foojar": 4,
		"fomang": 5,
		"liusi":  32,
	}
	for key, val := range maps {
		trie.Insert(key, val)
	}
	actual := trie.ToMap()
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

func TestTrie_Prefix(t *testing.T) {
	trie := NewTrie()
	maps := map[string]int{
		"foo":    1,
		"foobar": 2,
		"foball": 3,
		"foojar": 4,
		"fomang": 5,
		"liusi":  32,
	}
	for key, val := range maps {
		trie.Insert(key, val)
	}

	prefix := "fo"
	actual := trie.Prefix(prefix)
	expect := []string{"foo", "foobar", "foball", "foojar", "fomang"}
	sort.Strings(expect)
	assert.Truef(t, reflect.DeepEqual(expect, actual), "Expect:%v, actual:%v", expect, actual)

	prefix = "foo"
	actual = trie.Prefix(prefix)
	expect = []string{"foo", "foobar", "foojar"}
	sort.Strings(expect)
	fmt.Println()
	assert.Truef(t, reflect.DeepEqual(expect, actual), "Expect:%v, actual:%v", expect, actual)

	prefix = "l"
	actual = trie.Prefix(prefix)
	expect = []string{"liusi"}
	assert.Truef(t, reflect.DeepEqual(expect, actual), "Expect:%v, actual:%v", expect, actual)

	prefix = "sdf"
	actual = trie.Prefix(prefix)
	assert.Nilf(t, actual, "Expect:nil, actual:%v", actual)

	prefix = ""
	actual = trie.Prefix(prefix)
	expect = []string{"foo", "foobar", "foball", "foojar", "fomang", "liusi"}
	sort.Strings(expect)
	assert.Truef(t, reflect.DeepEqual(expect, actual), "Expect:%v, actual:%v", expect, actual)
}
