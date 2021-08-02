package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func CreatePage(summary string, title string, date string) {
	url := "https://api.notion.com/v1/pages"
	method := "POST"

	afterOneDay := time.Now().AddDate(0, 0, 4).Format(time.RFC3339)

	payload := strings.NewReader(`{
    "parent": {
        "database_id": ` + `"` + os.Getenv("database_id") + `"` + `
    },
    "properties": {
        "GCal_ID": {
            "type": "rich_text",
            "rich_text": [
                {
                    "id": "_}uo",
                    "type": "text",
                    "text": {
                        "content": ` + `"` + summary + `"` + `
                    }
                }
            ]
        },
        "Event Start": {
            "id": "u|mS",
            "type": "date",
            "date": {
                "id": "u|mS",
                "start": ` + `"` + date + `"` + `,
                "end": ` + `"` + afterOneDay + `"` + `
            }
        },
        "Name": {
            "type": "title",
            "title": [
                {
                    "type": "text",
                    "text": {
                        "content": ` + `"` + title + `"` + `
                    }
                }
            ]
        }
    }
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("notion_api"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", "2021-05-13")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

type createPageType struct {
	Parent struct {
		DatabaseID string `json:"database_id"`
	} `json:"parent"`
	Properties struct {
		GCalID struct {
			Type     string `json:"type"`
			RichText []struct {
				ID   string `json:"id"`
				Type string `json:"type"`
				Text struct {
					Content string `json:"content"`
				} `json:"text"`
			} `json:"rich_text"`
		} `json:"GCal_ID"`
		Date struct {
			ID   string `json:"id"`
			Type string `json:"type"`
			Date struct {
				ID    string `json:"id"`
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"date"`
		} `json:"Date"`
		Name struct {
			Type  string `json:"type"`
			Title []struct {
				Type string `json:"type"`
				Text struct {
					Content string `json:"content"`
				} `json:"text"`
			} `json:"title"`
		} `json:"Name"`
	} `json:"properties"`
}
