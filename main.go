package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/tdewolff/minify/v2"
	minifyhtml "github.com/tdewolff/minify/v2/html"
	"golang.org/x/net/html"
	"html_to_xhtml_converter/versions"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Command line flags
	inputFilePath := flag.String("i", "", "Input HTML file path (optional, reads from stdin if not provided)")
	outputFilePath := flag.String("o", "", "Output XHTML file path (optional, writes to stdout if not provided)")
	shouldMinify := flag.Bool("minify", false, "Minify the XHTML output")
	escapeForJSON := flag.Bool("escape-for-json", false, "Escape XHTML for embedding into JSON")
	versionsFilePath := flag.String("versions-filepath", "project_versions.json", "Path of versions JSON file")
	mocksVersionsFilePath := flag.String("mocks-versions-filepath", "project_versions_mocks.json", "Path of mocks versions JSON file")
	confluencePageTitle := flag.String("confluence-page-title", "", "Title of the Confluence page")
	confluenceSpaceCode := flag.String("confluence-space-code", "", "Space code of the Confluence space")
	confluenceAncestorPageId := flag.Int("confluence-ancestor-page-id", 0, "ID of the ancestor Confluence page")
	confluenceAuthPersonalToken := flag.String("confluence-auth-personal-token", "", "Personal access token for Confluence API")
	flag.Parse()

	// Read HTML input
	var htmlContent string
	if *inputFilePath != "" {
		// Read from file
		content, err := os.ReadFile(*inputFilePath)
		if err != nil {
			panic(err)
		}
		htmlContent = versions.Parse(*versionsFilePath, *mocksVersionsFilePath) + "<br></br><br></br>" + string(content)
	} else {
		// Read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			htmlContent += scanner.Text() + "\n"
		}
	}

	var finalHTML string
	if *shouldMinify {
		// Initialize minifier and minify HTML
		m := minify.New()
		m.AddFunc("text/html", minifyhtml.Minify)
		minifiedHTML, err := m.String("text/html", htmlContent)
		if err != nil {
			log.Fatalf("Failed to minify: %v", err)
		}
		finalHTML = minifiedHTML
	} else {
		finalHTML = htmlContent
	}

	// Parse the HTML
	doc, err := html.Parse(strings.NewReader(finalHTML))
	if err != nil {
		panic(err)
	}

	// Convert to XHTML
	var b bytes.Buffer
	renderNode(&b, doc)

	// Store XHTML output
	output := b.Bytes()

	// Escape for JSON if flag is set
	if *escapeForJSON {
		escapedOutput := strings.ReplaceAll(string(output), "\"", "\\\"")
		output = []byte(escapedOutput)
	}

	// Output the XHTML
	if *outputFilePath != "" {
		if strings.HasPrefix(*outputFilePath, "file://") {
			// File output
			filePath := strings.TrimPrefix(*outputFilePath, "file://")
			err := os.WriteFile(filePath, output, 0644)
			if err != nil {
				panic(err)
			}
		} else if strings.HasPrefix(*outputFilePath, "confluence://") {

			if *confluencePageTitle == "" || *confluenceSpaceCode == "" || *confluenceAuthPersonalToken == "" {
				log.Fatal("Confluence details are required: page title, space code, and auth token")
			}
			// Confluence output
			confluenceURL := strings.TrimPrefix(*outputFilePath, "confluence://")
			SendToConfluence(*confluencePageTitle, *confluenceSpaceCode, string(output), *confluenceAuthPersonalToken, *confluenceAncestorPageId, confluenceURL)
		} else {
			fmt.Println("Invalid output destination")
		}
	} else {
		// Write to stdout
		fmt.Print(string(output))
	}
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

func SendToConfluence(pageTitle, spaceCode, xhtmlContent, authBearerToken string, ancestorId int, confluenceURL string) {
	url := confluenceURL + "/rest/api/content"
	method := "POST"

	jsonPayload := fmt.Sprintf(`{
		"type": "page",
		"title": "%s",
		"ancestors": [{"id":%d}],
		"space": {
			"key": "%s"
		},
		"body": {
			"storage": {
				"value": "%s",
				"representation": "storage"
			}
		}
	}`, pageTitle, ancestorId, spaceCode, xhtmlContent)

	payload := strings.NewReader(jsonPayload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatalf("An error occurred: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authBearerToken))

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("An error occurred: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("An error occurred: %v", err)
	}
	fmt.Println(string(body))
}
