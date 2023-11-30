package jira

import (
	. "html_to_xhtml_converter/config"
)

func GetReleaseNotesTicketsByVersion(config Config, fixVersion string) map[string]Ticket {
	return parseTickets(getReleaseNotesTicketsByVersion(config.JiraBaseUrl, fixVersion, config.JiraAuthPersonalToken))
}
