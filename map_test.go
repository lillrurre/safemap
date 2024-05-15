package safemap

import (
	"fmt"
	"slices"
	"testing"
)

func TestNewMap(t *testing.T) {
	type testType struct{}
	safeMap := NewMap[int, []*testType]()

	if safeMap == nil {
		t.Errorf("NewMap() returned nil")
	}

	if safeMap.m == nil {
		t.Errorf("Expected NewMap() to not be nil")
	}

	if len(safeMap.m) != 0 {
		t.Errorf("Expected empty map from NewMap()")
	}
}

func TestSafeMap_Store(t *testing.T) {
	safeMap := NewMap[string, string]()

	if s, ok := safeMap.m["foo"]; ok || s != "" {
		t.Errorf("NewMap() returned non-empty Map, got %v", safeMap)
	}

	safeMap.m["foo"] = "bar"
	if s, ok := safeMap.m["foo"]; !ok || s != "bar" {
		t.Errorf("Expected 'foo' to return 'bar', got %v", s)
	}

}

func TestSafeMap_Load(t *testing.T) {
	safeMap := NewMap[string, string]()
	val := safeMap.Load("foo")
	if val != "" {
		t.Errorf("NewMap() returned value, got %s", val)
	}
	safeMap.Store("foo", "bar")
	val = safeMap.Load("foo")
	if val != "bar" {
		t.Errorf("Expected 'foo' to return 'bar', got %s", val)
	}
}

func TestSafeMap_Delete(t *testing.T) {
	safeMap := NewMap[string, string]()
	safeMap.m["foo"] = "bar"
	safeMap.m["rab"] = "baz"
	if val, ok := safeMap.m["foo"]; !ok || val != "bar" {
		t.Errorf("Expected 'foo' to return 'bar', got %s", val)
	}
	safeMap.Delete("foo")
	if val, ok := safeMap.m["foo"]; ok || val != "" {
		t.Errorf("Expected key 'foo' to be deleted, got map %v", safeMap)
	}
}

func TestSafeMap_LoadBool(t *testing.T) {
	safeMap := NewMap[string, string]()
	val, ok := safeMap.LoadBool("foo")
	if ok || val != "" {
		t.Errorf("Expected empty map, got %v", val)
	}
	safeMap.m["foo"] = "bar"
	val, ok = safeMap.LoadBool("foo")
	if !ok || val != "bar" {
		t.Errorf("Expected 'foo' to return 'true' and 'bar', got %s and map %v", val, safeMap)
	}
}

func TestSafeMap_Swap(t *testing.T) {
	safeMap := NewMap[string, string]()
	safeMap.m["foo"] = "bar"
	if val, ok := safeMap.m["foo"]; !ok || val != "bar" {
		t.Errorf("Expected 'foo' to return 'bar', got %s", val)
	}
	safeMap.Swap("foo", "lol")
	if val, ok := safeMap.m["foo"]; !ok || val != "lol" {
		t.Errorf("Expected 'foo' to return 'lol', got %s", val)
	}
}

func TestSafeMap_Range(t *testing.T) {
	expected := []string{"foo=bar", "rab=baz", "lol=lol"}

	// Test case where func returns true
	{
		safeMap := NewMap[string, string]()
		safeMap.m["foo"] = "bar"
		safeMap.m["rab"] = "baz"
		safeMap.m["lol"] = "lol"

		safeMap.Range(func(k string, v string) bool {
			if !slices.Contains(expected, fmt.Sprintf("%s=%s", k, v)) {
				t.Errorf("Expected slice %v to contain %q with value %q", expected, k, v)
			}
			return true
		})
	}

	// Test case where func returns false
	{
		safeMap := NewMap[string, string]()
		safeMap.m["uuga"] = "buuga"

		safeMap.Range(func(k string, v string) bool {
			if slices.Contains(expected, fmt.Sprintf("%s=%s", k, v)) {
				t.Errorf("Expected slice %v to contain %q with value %q", expected, k, v)
			}
			return false
		})
	}

}

func TestSafeMap_RangeValue(t *testing.T) {
	expected := []string{"bar", "baz", "lol"}

	// Test case where func returns true
	{
		safeMap := NewMap[string, string]()
		safeMap.m["foo"] = "bar"
		safeMap.m["rab"] = "baz"
		safeMap.m["lol"] = "lol"

		safeMap.RangeValue(func(v string) bool {
			if !slices.Contains(expected, v) {
				t.Errorf("Expected slice %v to contain value %q", expected, v)
			}
			return true
		})
	}

	// Test case where func returns false
	{
		safeMap := NewMap[string, string]()
		safeMap.m["uuga"] = "buuga"

		safeMap.RangeValue(func(v string) bool {
			if slices.Contains(expected, v) {
				t.Errorf("Expected slice %v to contain value %q", expected, v)
			}
			return false
		})
	}
}

func TestSafeMap_RangeKey(t *testing.T) {
	expected := []string{"foo", "rab", "lol"}

	// Test case where func returns true
	{
		safeMap := NewMap[string, string]()
		safeMap.m["foo"] = "bar"
		safeMap.m["rab"] = "baz"
		safeMap.m["lol"] = "lol"

		safeMap.RangeKey(func(k string) bool {
			if !slices.Contains(expected, k) {
				t.Errorf("Expected slice %v to contain key %q", expected, k)
			}
			return true
		})
	}

	// Test case where func returns false
	{
		safeMap := NewMap[string, string]()
		safeMap.m["uuga"] = "buuga"

		safeMap.RangeKey(func(k string) bool {
			if slices.Contains(expected, k) {
				t.Errorf("Expected slice %v to contain key %q", expected, k)
			}
			return false
		})
	}
}

func TestSafeMap_Len(t *testing.T) {
	safeMap := NewMap[string, string]()
	if safeMap.Len() != 0 {
		t.Errorf("Expected 0, got %d", safeMap.Len())
	}

	safeMap.m["foo"] = "bar"
	if safeMap.Len() != 1 {
		t.Errorf("Expected 1, got %d", safeMap.Len())
	}

	delete(safeMap.m, "foo")

	if safeMap.Len() != 0 {
		t.Errorf("Expected 0, got %d", safeMap.Len())
	}
}
