package output

import (
	"fmt"
	. "html_to_xhtml_converter/config"
	"html_to_xhtml_converter/confluence"
	"log"
	"os"
	"strings"
)

func WriteOutput(output []byte, config Config) {
	if config.EscapeForJSON {
		output = []byte(strings.ReplaceAll(string(output), "\"", "\\\""))
	}

	if config.OutputFilePath != "" {
		handleFileOutput(output, config)
	} else {
		fmt.Print(string(output))
	}
}

func handleFileOutput(output []byte, config Config) {
	if strings.HasPrefix(config.OutputFilePath, "file://") {
		writeToFile(output, strings.TrimPrefix(config.OutputFilePath, "file://"))
	} else if strings.HasPrefix(config.OutputFilePath, "confluence://") {
		writeToConfluence(output, config)
	} else {
		fmt.Println("Invalid output destination")
	}
}

func writeToFile(output []byte, filePath string) {
	err := os.WriteFile(filePath, output, 0644)
	if err != nil {
		panic(err)
	}
}

func writeToConfluence(output []byte, config Config) {
	if config.ConfluencePageTitle == "" || config.ConfluenceSpaceCode == "" || config.ConfluenceAuthPersonalToken == "" {
		log.Fatal("Confluence details are required: page title, space code, and auth token")
	}
	confluenceURL := strings.TrimPrefix(config.OutputFilePath, "confluence://")
	confluence.CreatePage(config.ConfluencePageTitle, config.ConfluenceSpaceCode, string(output), config.ConfluenceAuthPersonalToken, config.ConfluenceAncestorPageId, confluenceURL)
}
