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

// func TestURLString(t *testing.T) {
// 	u := &URL{
// 		Scheme: "https",
// 		Host:   "go.dev",
// 		Path:   "play",
// 	}
// 	got := u.String()
// 	want := "https://go.dev/play"

// 	if got != want {
// 		t.Errorf("String() = %q, want %q", got, want)
// 	}
// }

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

func TestParseError(t *testing.T) {
	tests := []struct{ name, uri string }{
		{name: "without_scheme", uri: "github.dev/play"},
		{name: "empty_scheme", uri: "://github.dev"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("run %s", tt.name)
			_, err := Parse(tt.uri)
			if err == nil {
				t.Errorf("Parse(%q) err = <nil>, want an error", tt.uri)
			}
		})
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

func TestURLString(t *testing.T) {
	tests := []struct {
		name string
		uri  *URL
		want string
	}{
		{name: "nil", uri: nil, want: ""},
		{name: "empty", uri: &URL{}, want: ""},
		/* we'll add more test cases soon */
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.uri.String()
			if got != tt.want {
				t.Errorf("\ngot %q\nwant %q\nfor %#v", got, tt.want, tt.uri)
			}
		})
	}
}
