package url

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var parseTests = []struct { // this struct is the format for the table test inputs and required outputs
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

var stringTests = []struct {
	name string
	url  *URL
	want string
}{
	{
		name: "complete_url",
		url:  &URL{Scheme: "https", Host: "go.dev", Path: "play"},
		want: "https://go.dev/play",
	},
	{
		name: "without_path",
		url:  &URL{Scheme: "https", Host: "go.dev"},
		want: "https://go.dev/",
	},
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

func TestParseTable(t *testing.T) {
	for _, tt := range parseTests {
		t.Logf("run %s", tt.name)
		got, err := Parse(tt.uri)
		if err != nil {
			t.Fatalf("Parse(%q) err = %q, want <nil>", tt.uri, err)
		}

		if diff := cmp.Diff(got, tt.want); diff != "" {
			t.Errorf("Parse(%q) mismatch (-want +got):\n%s", tt.uri, diff)
		}
	}
}

func TestURLStringTable(t *testing.T) { // here we have used subtests to run multiple tests
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("run %s", tt.name)
			got := tt.url.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("String() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
