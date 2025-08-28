package common

import (
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJsonToAny(t *testing.T) {
	raw := map[string]any{
		"name": "Dev Isho",
		"age":  25,
	}

	var u User
	err := JsonToAny(raw, &u)
	if err != nil {
		t.Fatalf("JsonToAny failed: %v", err)
	}

	if u.Name != "Dev Isho" || u.Age != 25 {
		t.Errorf("unexpected result: got %+v", u)
	}
}

func TestAnyToJson(t *testing.T) {
	u := User{Name: "Dev Isho", Age: 25}

	jsonStr, err := AnyToJson(u)
	if err != nil {
		t.Fatalf("AnyToJson failed: %v", err)
	}

	expected := `{"name":"Dev Isho","age":25}`
	if jsonStr != expected {
		t.Errorf("expected %s, got %s", expected, jsonStr)
	}
}
