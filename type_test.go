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
		{"test.tiff", TIFF},
	}

	for _, file := range files {
		img, _ := os.Open(path.Join("fixtures", file.name))
		buf, _ := ioutil.ReadAll(img)
		defer img.Close()

		actual := DetermineImageType(buf)
		if IsTypeSupported(file.expected) && file.expected != actual {
			t.Fatalf("Image type %#v != %#v", ImageTypes[file.expected], ImageTypes[actual])
		}
	}
}

func TestDeterminateImageTypeName(t *testing.T) {
	files := []struct {
		name     string
		expected string
		typ      ImageType
	}{
		{"test.jpg", "jpeg", JPEG},
		{"test.png", "png", PNG},
		{"test.webp", "webp", WEBP},
		{"test.gif", "gif", GIF},
		{"test.pdf", "pdf", PDF},
		{"test.svg", "svg", SVG},
		{"test.jp2", "magick", MAGICK},
		{"test.tiff", "tiff", TIFF},
	}

	for _, file := range files {
		img, _ := os.Open(path.Join("fixtures", file.name))
		buf, _ := ioutil.ReadAll(img)
		defer img.Close()

		actual := DetermineImageTypeName(buf)
		if IsTypeSupported(file.typ) && file.expected != actual {
			t.Fatalf("Image type %#v != %#v", file.expected, actual)
		}
	}
}

func TestIsTypeSupported(t *testing.T) {
	types := []struct {
		name     ImageType
		expected bool
	}{
		{JPEG, true},
		{PNG, true},
		{WEBP, true},
		{GIF, VipsVersion >= "8.3.0"},
		{PDF, true},
		{SVG, true},
		{TIFF, VipsVersion >= "8.0.0"},
	}

	for _, n := range types {
		actual := IsTypeSupported(n.name)
		if n.expected != actual {
			t.Fatalf("Image type %#v support: %#v != %#v", ImageTypes[n.name], n.expected, actual)
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
		{"gif", VipsVersion >= "8.3.0"},
		{"pdf", true},
		{"svg", true},
		{"tiff", VipsVersion >= "8.0.0"},
	}

	for _, n := range types {
		actual := IsTypeNameSupported(n.name)
		if n.expected != actual {
			t.Fatalf("Image type %#v support: %#v != %#v", n.name, n.expected, actual)
		}
	}
}

func TestIsTypeSupportedSave(t *testing.T) {
	types := []struct {
		name ImageType
	}{
		{JPEG}, {PNG}, {WEBP},
	}
	if VipsVersion >= "8.5.0" {
		types = append(types, struct{ name ImageType }{TIFF})
	}

	for _, n := range types {
		if IsTypeSupportedSave(n.name) == false {
			t.Fatalf("Image type %#v is not valid", ImageTypes[n.name])
		}
	}
}

func TestIsTypeNameSupportedSave(t *testing.T) {
	types := []struct {
		name     string
		expected bool
	}{
		{"jpeg", true},
		{"png", true},
		{"webp", true},
		{"gif", false},
		{"pdf", false},
		{"tiff", VipsVersion >= "8.5.0"},
	}

	for _, n := range types {
		if IsTypeNameSupportedSave(n.name) != n.expected {
			t.Fatalf("Image type %#v is not valid", n.name)
		}
	}
}
