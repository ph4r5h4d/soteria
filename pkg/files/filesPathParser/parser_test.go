package filesPathParser

import (
	"os"
	"reflect"
	"testing"
)

func TestParseFilesPath(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		i   []string
		exp []string
	}{
		{
			[]string{"/home/file.json", "/home/.farshad"},
			[]string{"/home/file.json", "/home/.farshad"},
		},
		{
			[]string{"~/file.json", "/home/.farshad"},
			[]string{home + "/file.json", "/home/.farshad"},
		},
	}

	for _, e := range cases {
		res, err := ParseFilesPath(e.i)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(e.exp, res) {
			t.Errorf("Expected: \n %v \n Got: %v", e.exp, res)
		}
	}
}
