package stats

import (
	"reflect"
	"testing"
)

func TestMakeTimer(t *testing.T) {
	tests := []struct {
		key  string
		name string
		tags []Tag
	}{
		{
			key:  "?",
			name: "",
			tags: nil,
		},
		{
			key:  "M?",
			name: "M",
			tags: nil,
		},
		{
			key:  "M?A=1",
			name: "M",
			tags: []Tag{{"A", "1"}},
		},
		{
			key:  "M?A=1&B=2",
			name: "M",
			tags: []Tag{{"B", "2"}, {"A", "1"}},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			if timer := makeTimer(nil, test.name, test.tags); !reflect.DeepEqual(timer, Timer{
				eng:  nil,
				key:  test.key,
				name: test.name,
				tags: test.tags,
			}) {
				t.Errorf("makeTimer(nil, %#v, %#v) => %#v", test.name, test.tags, timer)
			}
		})
	}
}