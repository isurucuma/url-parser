package url

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var parseTests = []struct {
	name string
	uri  string
	want *URL
}{
	{
		name: "full",
		uri:  "https://go.dev/play",
		want: &URL{Scheme: "https", Host: "go.dev", Path: "play"},
	},
	{
		name: "without_path",
		uri:  "http://github.com",
		want: &URL{Scheme: "http", Host: "github.com", Path: ""},
	},
	/* many more test cases can be easily added */
}

func TestParse(t *testing.T) {
	const url = "https://go.dev/play"

	got, err := Parse(url)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want <nil>", url, err)
	}
	want := &URL{
		Scheme: "https",
		Host:   "go.dev",
		Path:   "play",
	}
	if *got != *want {
		t.Errorf("Parse(%q)\ngot %#v\nwant %#v", url, got, want)
	}
}

func TestURLString(t *testing.T) {
	u := &URL{
		Scheme: "https",
		Host:   "go.dev",
		Path:   "play",
	}
	got := u.String()
	want := "https://go.dev/play"

	if got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}

func TestParseWithoutPath(t *testing.T) {
	const url = "https://github.dev"

	got, err := Parse(url)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want <nil>", url, err)
	}

	want := &URL{
		Scheme: "https",
		Host:   "github.dev",
		Path:   "",
	}

	// here we have used the google/go-cmp/cmp package to compare the two structs
	// this is a better way to compare two structs and it will give you a better error message
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Parse(%q) mismatch (-want +got):\n%s", url, diff)
	}
}

func TestParseWithoutScheme(t *testing.T) {
	const url = "github.dev/play"

	_, err := Parse(url)
	if err == nil {
		t.Fatalf("Parse(%q) err = <nil>, want error", url)
	}
}
