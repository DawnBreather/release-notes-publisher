package config

import "flag"

var ConfigStorage Config

type Config struct {
	InputFilePath               string
	OutputFilePath              string
	ShouldMinify                bool
	EscapeForJSON               bool
	VersionsFilePath            string
	MocksVersionsFilePath       string
	ConfluencePageTitle         string
	ConfluenceSpaceCode         string
	ConfluenceAncestorPageId    int
	ConfluenceAuthPersonalToken string
	JiraBaseUrl                 string
	JiraAuthPersonalToken       string
}

func ParseFlags() Config {
	inputFilePath := flag.String("i", "", "Input HTML file path")
	outputFilePath := flag.String("o", "", "Output XHTML file path")
	shouldMinify := flag.Bool("minify", false, "Minify the XHTML output")
	escapeForJSON := flag.Bool("escape-for-json", false, "Escape XHTML for embedding into JSON")
	versionsFilePath := flag.String("versions-filepath", "project_versions.json", "Path of versions JSON file")
	mocksVersionsFilePath := flag.String("mocks-versions-filepath", "project_versions_mocks.json", "Path of mocks versions JSON file")
	confluencePageTitle := flag.String("confluence-page-title", "", "Title of the Confluence page")
	confluenceSpaceCode := flag.String("confluence-space-code", "", "Space code of the Confluence space")
	confluenceAncestorPageId := flag.Int("confluence-ancestor-page-id", 0, "ID of the ancestor Confluence page")
	confluenceAuthPersonalToken := flag.String("confluence-auth-personal-token", "", "Personal access token for Confluence API")
	jiraBaseUrl := flag.String("jira-base-url", "", "JIRA base url")
	jiraAuthPersonalToken := flag.String("jira-auth-personal-token", "", "Personal access token for JIRA API")

	flag.Parse()

	return Config{
		*inputFilePath, *outputFilePath, *shouldMinify, *escapeForJSON,
		*versionsFilePath, *mocksVersionsFilePath, *confluencePageTitle,
		*confluenceSpaceCode, *confluenceAncestorPageId, *confluenceAuthPersonalToken,
		*jiraBaseUrl, *jiraAuthPersonalToken,
	}
}
