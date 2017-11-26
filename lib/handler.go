package random

import (
	"fmt"
	"net/http"
)

func NewRootHandler(s *LinkStore) *RootHandler {
	return &RootHandler{store: s}
}

type RootHandler struct {
	store *LinkStore
}

func (r *RootHandler) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {
	colorLinks := r.store.All()
	links := renderLinks(colorLinks)

	index := renderIndex(links)

	rw.Write(index)
}

func renderLinks(cs []ColorLink) string {
	var links string
	for _, c := range cs {
		links += fmt.Sprintf(linkFormat, c.Link, c.Color)
	}
	return links
}

func renderIndex(links string) []byte {
	index := fmt.Sprintf(indexFormat, links)
	return []byte(index)
}

const (
	linkFormat  = `<a href="%s"><div class="swatch" style="background-color:%s;"></div></a>`
	indexFormat = `<html>
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>random</title>
	<style>
		.swatch {
			display: inline-block;
			height: 1em;
			width: 1em;
		}
		.grid .swatch {
			margin-right: 0.25em;
			margin-bottom: 0.25em;
		}
	</style>
</head>
<body>
	<div class="grid">
		%s
	</div>
</body>
</html>`
)
