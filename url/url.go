package url

import (
	"errors"
	"fmt"
	"strings"
)

type URL struct {
	Scheme string
	Host   string
	Path   string
}

func Parse(rawURL string) (*URL, error) {
	scheme, rest, ok := strings.Cut(rawURL, "://")
	if !ok {
		return nil, errors.New("missing scheme")
	}

	host, path, _ := strings.Cut(rest, "/")

	return &URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}, nil
}

func (u *URL) String() string {
	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
}
