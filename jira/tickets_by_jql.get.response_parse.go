package jira

import (
	"encoding/json"
	"fmt"
	"html_to_xhtml_converter/xhtml"
)

// Define the JSON structure in Go structs
type Response struct {
	Issues []Issue `json:"issues"`
}

type Issue struct {
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

type Fields struct {
	Description string    `json:"description"`
	FixVersions []Version `json:"fixVersions"`
	Status      Status    `json:"status"`
	Summary     string    `json:"summary"`
}

type Status struct {
	Name string `json:"name"`
}

//type Issuetype struct {
//	Description string `json:"description"`
//}

type Version struct {
	Name string `json:"name"`
}

// The desired map structure
type Ticket struct {
	Key         string
	Description string
	Summary     string
	Status      string
}

func parseTickets(jsonResponse []byte) map[string]Ticket {
	// Unmarshal JSON to the struct
	var resp Response
	err := json.Unmarshal([]byte(jsonResponse), &resp)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}

	// Extract data and populate the map
	ticketMap := make(map[string]Ticket)
	for _, jIssue := range resp.Issues {
		if len(jIssue.Fields.FixVersions) > 0 {
			ticketMap[jIssue.Fields.FixVersions[0].Name] = Ticket{
				Key:         jIssue.Key,
				Description: xhtml.ConvertDescriptionToXHTML(jIssue.Fields.Description),
				Summary:     jIssue.Fields.Summary,
				Status:      jIssue.Fields.Status.Name,
			}
		}
	}

	// Print the map or use it as needed
	return ticketMap
}
