package versions

import (
	"encoding/json"
	"fmt"
	"github.com/thoas/go-funk"
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
	xhtml += "<tr><<th>Version</th><th>Release Notes Description</th></tr>\n"

	// Add components to XHTML
	processedComponentsVersions := []string{}
	for _, comp := range versions {
		if !funk.Contains(processedComponentsVersions, comp.JiraVer) && comp.JiraVer != "" {
			releaseNotesTickets := GetReleaseNotesTicketsByVersion(config, comp.JiraVer)
			description := "DESCRIPTION"
			if len(releaseNotesTickets) > 0 {
				description = releaseNotesTickets[comp.JiraVer].Description
			}
			xhtml += fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>\n", comp.JiraVer, description)
			processedComponentsVersions = append(processedComponentsVersions, comp.JiraVer)
		}
	}

	// Add connectors to XHTML
	processedConnectorsVersions := []string{}
	for _, conn := range mocks.Connectors {
		if !funk.Contains(processedConnectorsVersions, conn.JiraVer) && conn.JiraVer != "" {
			releaseNotesTickets := GetReleaseNotesTicketsByVersion(config, conn.JiraVer)
			description := "DESCRIPTION"
			if len(releaseNotesTickets) > 0 {
				description = releaseNotesTickets[conn.JiraVer].Description
			}
			xhtml += fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>\n", conn.JiraVer, description)
			processedConnectorsVersions = append(processedConnectorsVersions, conn.JiraVer)
		}
	}

	// Add mocks to XHTML
	processedMocksVersions := []string{}
	for _, mock := range mocks.Mocks {
		if !funk.Contains(processedMocksVersions, mock.JiraVer) && mock.JiraVer != "" {
			releaseNotesTickets := GetReleaseNotesTicketsByVersion(config, mock.JiraVer)
			description := "DESCRIPTION"
			if len(releaseNotesTickets) > 0 {
				description = releaseNotesTickets[mock.JiraVer].Description
			}
			xhtml += fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>\n", mock.JiraVer, description)
			processedMocksVersions = append(processedMocksVersions, mock.JiraVer)
		}
	}

	// Close table
	xhtml += "</table>\n"

	return xhtml
}
