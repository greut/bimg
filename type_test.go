package bimg

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestDeterminateImageType(t *testing.T) {
	files := []struct {
		name     string
		expected ImageType
	}{
		{"test.jpg", JPEG},
		{"test.png", PNG},
		{"test.webp", WEBP},
		{"test.gif", GIF},
		{"test.pdf", PDF},
		{"test.svg", SVG},
		{"test.jp2", MAGICK},
	}

	for _, file := range files {
		if !VipsIsTypeSupported(file.expected) {
			continue
		}
		img, _ := os.Open(path.Join("fixtures", file.name))
		buf, _ := ioutil.ReadAll(img)
		defer img.Close()

		if DetermineImageType(buf) != file.expected {
			t.Fatalf("Image type %#v is not valid", ImageTypes[file.expected])
		}

		if DetermineImageTypeName(buf) != ImageTypes[file.expected] {
			t.Fatal("Image type %#v is not valid", ImageTypes[file.expected])
		}
	}
}

func TestIsTypeSupported(t *testing.T) {
	types := []struct {
		name ImageType
	}{
		{JPEG}, {PNG}, {WEBP}, {GIF}, {PDF},
	}

	for _, n := range types {
		if IsTypeSupported(n.name) == false {
			t.Fatalf("Image type %#v is not valid", n.name)
		}
	}
}

func TestIsTypeNameSupported(t *testing.T) {
	types := []struct {
		name     string
		expected bool
	}{
		{"jpeg", true},
		{"png", true},
		{"webp", true},
		{"gif", true},
		{"pdf", true},
	}

	for _, n := range types {
		if IsTypeNameSupported(n.name) != n.expected {
			t.Fatalf("Image type %#v is not valid", n.name)
		}
	}
}
