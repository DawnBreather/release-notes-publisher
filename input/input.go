package input

import (
	"bufio"
	. "html_to_xhtml_converter/config"
	"html_to_xhtml_converter/versions"
	"os"
)

func ReadInputContent(config Config) string {
	if config.InputFilePath != "" {
		content, err := os.ReadFile(config.InputFilePath)
		if err != nil {
			panic(err)
		}
		return versions.Parse(config) + "<br></br><br></br>" + string(content)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var htmlContent string
	for scanner.Scan() {
		htmlContent += scanner.Text() + "\n"
	}
	return htmlContent
}
