package confluence

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func CreatePage(pageTitle, spaceCode, xhtmlContent, authBearerToken string, ancestorId int, confluenceURL string) {
	url := confluenceURL + "/rest/api/content"
	method := "POST"

	jsonPayload := fmt.Sprintf(`{
		"type": "page",
		"title": "%s",
		"ancestors": [{"id":%d}],
		"space": {
			"key": "%s"
		},
		"body": {
			"storage": {
				"value": "%s",
				"representation": "storage"
			}
		}
	}`, pageTitle, ancestorId, spaceCode, xhtmlContent)

	payload := strings.NewReader(jsonPayload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatalf("An error occurred: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authBearerToken))

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("An error occurred: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("An error occurred: %v", err)
	}
	fmt.Println(string(body))
}
