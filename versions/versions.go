package versions

import (
	"encoding/json"
	"fmt"
	. "html_to_xhtml_converter/config"
	. "html_to_xhtml_converter/jira"
	"os"
)

type Component struct {
	ChartVer   string            `json:"chartVer"`
	Comment    string            `json:"comment"`
	JiraVer    string            `json:"jiraVer"`
	JiraFixVer string            `json:"jiraFixVer"`
	Namespaces map[string]string `json:"namespaces"`
}

// ProjectVersions directly maps component names to their data.
type ProjectVersions map[string]Component

// ProjectVersionsMocks has separate maps for connectors and mocks.
type ProjectVersionsMocks struct {
	Connectors map[string]Component `json:"connectors"`
	Mocks      map[string]Component `json:"mocks"`
}

func Parse(config Config) string {
	// Load the JSON files into structs
	versions := loadVersions(config.VersionsFilePath)
	mocks := loadMocks(config.MocksVersionsFilePath)

	// Process the data and generate XHTML
	//xhtml := generateXHTML(versions, mocks)

	return generateXHTML(versions, mocks, config)

	//// Output XHTML to a file
	//err := os.WriteFile("output.xhtml", []byte(xhtml), 0644)
	//if err != nil {
	//	panic(err)
	//}
}

func loadVersions(filename string) ProjectVersions {
	var versions ProjectVersions
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &versions)
	if err != nil {
		panic(err)
	}
	return versions
}

func loadMocks(filename string) ProjectVersionsMocks {
	var mocks ProjectVersionsMocks
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &mocks)
	if err != nil {
		panic(err)
	}
	return mocks
}

func generateXHTML(versions ProjectVersions, mocks ProjectVersionsMocks, config Config) string {
	// Start XHTML document with table
	xhtml := "<table>\n"
	xhtml += "<tr><th>Version</th><th>Release Notes Description</th></tr>\n"

	// Process and add components, connectors, and mocks
	xhtml += processItems(versions, config)
	xhtml += processItems(mocks.Connectors, config)
	xhtml += processItems(mocks.Mocks, config)

	// Close table
	xhtml += "</table>\n"

	return xhtml
}

func processItems(items map[string]Component, config Config) string {
	var result string
	processedVersions := make(map[string]bool)

	for _, item := range items {
		if !processedVersions[item.JiraVer] && item.JiraVer != "" {
			releaseNotesTickets := GetReleaseNotesTicketsByVersion(config, item.JiraVer)
			description := "DESCRIPTION"
			jiraVer := item.JiraVer
			if len(releaseNotesTickets) > 0 {
				description = releaseNotesTickets[item.JiraVer].Description
				jiraVer = fmt.Sprintf("<a href=\"%s/browse/%s\">%s</a>", config.JiraBaseUrl, releaseNotesTickets[item.JiraVer].Key, jiraVer)
			}
			result += fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>\n", jiraVer, description)
			processedVersions[item.JiraVer] = true
		}
	}

	return result
}
