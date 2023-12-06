package jira

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getReleaseNotesTicketsByVersion(jiraBaseUrl, fixVersion, authBearerToken string) []byte {

	url := fmt.Sprintf("%s/rest/api/2/search", jiraBaseUrl)
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{"jql": "issuetype in (\"Release Note\") AND fixVersion in (%s) AND status not in (REJECTED)", "maxResults": 1000}`, fixVersion))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authBearerToken))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return body
}
