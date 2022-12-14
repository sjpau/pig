package grep

import (
	"testing"
)

func TestIsFileFormat(t *testing.T) {
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
			if result := IsFileFormat(j.in); result != j.expected {
				t.Errorf("Expected %t, received %t", j.expected, result)
			}
		})
	}
	unitsWithExt := []struct {
		desc     string
		in       string
		expected bool
	}{
		{"Match with argument", "/tmp/private/file.aseprite", true},
	}
	for _, j := range unitsWithExt {
		t.Run(j.desc, func(t *testing.T) {
			if result := IsFileFormat(j.in, ".aseprite"); result != j.expected {
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

func TestDomainNameFromURL(t *testing.T) {
	units := []struct {
		desc     string
		in       string
		expected string
	}{
		{"https regular", "https://mobyus.xyz", "mobyus.xyz"},
		{"https + trailing slash", "https://mobyus.xyz/", "mobyus.xyz"},
		{"http regular", "http://mobyus.xyz", "mobyus.xyz"},
		{"http + trailing slash", "http://mobyus.xyz/", "mobyus.xyz"},
	}
	for _, j := range units {
		t.Run(j.desc, func(t *testing.T) {
			if result := DomainNameFromURL(j.in); result != j.expected {
				t.Errorf("Expected %s, received %s", j.expected, result)
			}
		})
	}
}
