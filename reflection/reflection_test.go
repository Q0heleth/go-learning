package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestReflection(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		Expectedcalls []string
	}{
		{
			"Struct with one string field",
			struct{ Name string }{"chris"},
			[]string{"chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"Arrays",
			[2]Profile {
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}
	for _, test := range cases {
		var got []string
		walk(test.Input, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, test.Expectedcalls) {
			t.Errorf("got %v, want %v", got, test.Expectedcalls)
		}
	}
}
func walk(x interface{}, fn func(s string)) {
	v := getValueof(x)
	var getField func(int) reflect.Value
	fieldNumbers := 0
	switch v.Kind() {
	case reflect.String:
		fn(v.String())
	case reflect.Struct:
		fieldNumbers = v.NumField()
		getField =v.Field
	case reflect.Slice,reflect.Array:
		fieldNumbers=v.Len()
		getField=v.Index
	case reflect.Map:
		for _,key:=range v.MapKeys(){
			walk(v.MapIndex(key).Interface(),fn)
		}
	}
	for i:=0;i<fieldNumbers;i++{
		walk(getField(i).Interface(),fn)
	}
}
func getValueof(x interface{}) reflect.Value {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}
