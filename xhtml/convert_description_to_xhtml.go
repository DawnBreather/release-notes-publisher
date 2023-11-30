package xhtml

import (
	"bufio"
	"bytes"
	"strings"
)

func ConvertDescriptionToXHTML(input string) string {
	var buffer bytes.Buffer
	scanner := bufio.NewScanner(strings.NewReader(input))
	var inList bool

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Check for bullet point
		if strings.HasPrefix(line, "*") {
			line = strings.TrimPrefix(line, "*")
			line = strings.TrimSpace(line)

			if !inList {
				buffer.WriteString("<ul>\n")
				inList = true
			}

			buffer.WriteString("<li>")
			buffer.WriteString(line)
			buffer.WriteString("</li>\n")
		} else {
			if inList {
				buffer.WriteString("</ul>\n")
				inList = false
			}

			if line != "" {
				buffer.WriteString("<p>")
				buffer.WriteString(line)
				buffer.WriteString("</p>\n")
			}
		}
	}

	// Close the list tag if it's still open
	if inList {
		buffer.WriteString("</ul>\n")
	}

	return buffer.String()
}
