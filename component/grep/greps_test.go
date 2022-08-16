package grep

import (
	"reflect"
	"testing"
)

func TestIsKnownFile(t *testing.T) {
	units := []struct {
		desc     string
		in       string
		expected bool
	}{
		{"Match path", "/tmp/private/file.pdf", true},
		{"Match default path", "/file.json", true},
		{"Irregular path", "file.part", true},
		{"No ext", "/tmp/private/file", false},
		{"No ext default path", "/file", false},
		{"No ext irregular path", "file", false},
	}
	for _, j := range units {
		t.Run(j.desc, func(t *testing.T) {
			if result := IsKnownFile(j.in); result != j.expected {
				t.Errorf("Expected %t, received %t", j.expected, result)
			}
		})
	}
}

func TestFileNameFromPath(t *testing.T) {
	units := []struct {
		desc     string
		in       string
		expected string
	}{
		{"", "/asd/das/filename.xxxx", "filename.xxxx"},
		{"", "/asd/das/filename", "filename"},
		{"", "/asd/das/filename/", ""},
	}
	for _, j := range units {
		t.Run(j.desc, func(t *testing.T) {
			if result := FileNameFromPath(j.in); result != j.expected {
				t.Errorf("Expected %s, received %s", j.expected, result)
			}
		})
	}
}

func TestInternalDocumentPaths(t *testing.T) {
	t.Run("With specified extension.", func(t *testing.T) {
		if result := InternalDocumentPaths("<tag><a href=\"/asd/das/file.odt></a></tag>\"", ".odt", ".pdf"); !reflect.DeepEqual(result, []string{"/asd/das/file.odt"}) {
			t.Errorf("Failed to grep specified extension. %s", result)
		}
	})
}
