package url_test

import (
	"fmt"
	"log"
	"url-parser/url"
)

func ExampleParse() {
	u, err := url.Parse("https://go.dev/play")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)

	// Output:
	// https://go.dev/play
}
