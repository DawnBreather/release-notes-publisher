package xhtml

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func ConvertHTMLToXHTML(htmlContent string) []byte {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer
	renderNode(&buffer, doc)

	return buffer.Bytes()
}

// renderNode renders a single html.Node as XHTML.
func renderNode(b *bytes.Buffer, n *html.Node) {
	// Render the node itself
	if n.Type == html.ElementNode {
		// Self-closing for XHTML
		b.WriteString("<" + n.Data)
		for _, a := range n.Attr {
			b.WriteString(fmt.Sprintf(` %s="%s"`, a.Key, html.EscapeString(a.Val)))
		}
		if isSelfClosingTag(n.Data) {
			b.WriteString(" /")
		}
		b.WriteString(">")
	} else if n.Type == html.TextNode {
		b.WriteString(html.EscapeString(n.Data))
	}

	// Render child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		renderNode(b, c)
	}

	// Close the tag for non-self-closing elements
	if n.Type == html.ElementNode && !isSelfClosingTag(n.Data) {
		b.WriteString("</" + n.Data + ">")
	}
}

// isSelfClosingTag checks if a tag is self-closing in XHTML.
func isSelfClosingTag(tagName string) bool {
	switch tagName {
	case "area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param", "source", "track", "wbr":
		return true
	default:
		return false
	}
}
