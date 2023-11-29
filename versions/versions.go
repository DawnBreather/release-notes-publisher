package versions

import (
	"encoding/json"
	"fmt"
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

func Parse(versionsPath, mocksPath string) string {
	// Load the JSON files into structs
	versions := loadVersions(versionsPath)
	mocks := loadMocks(mocksPath)

	// Process the data and generate XHTML
	//xhtml := generateXHTML(versions, mocks)

	return generateXHTML(versions, mocks)

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

func generateXHTML(versions ProjectVersions, mocks ProjectVersionsMocks) string {
	// Start XHTML document with table
	xhtml := "<table>\n"
	xhtml += "<tr><th>Component ID</th><th>Version</th><th>Release Notes Description</th></tr>\n"

	// Add components to XHTML
	for id, comp := range versions {
		xhtml += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td></tr>\n", id, comp.JiraVer, "DESCRIPTION")
	}

	// Add connectors to XHTML
	for id, conn := range mocks.Connectors {
		xhtml += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td></tr>\n", id, conn.JiraVer, "DESCRIPTION")
	}

	// Add mocks to XHTML
	for id, mock := range mocks.Mocks {
		xhtml += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td></tr>\n", id, mock.JiraVer, "DESCRIPTION")
	}

	// Close table
	xhtml += "</table>\n"

	return xhtml
}
