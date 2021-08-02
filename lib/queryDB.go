package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func QueryDB() DatabaseQueryResponse {

	url := os.Getenv("url_root") + os.Getenv("database_id") + "/query"
	method := "POST"

	payload := strings.NewReader(`{
    "filter": {
        "and": [
            {
                "property": "Date",
                "date": {
                    "equals": "2021-08-02"
                }
            }
        ]
    }
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("notion_api"))
	req.Header.Add("Notion-Version", "2021-05-13")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	// data := &notion.DatabaseQueryResponse{body}
	// result := json.Marshal(body)
	var ans DatabaseQueryResponse
	json.Unmarshal([]byte(body), &ans)
	// fmt.Println(ans.Results)
	return ans
}

type DatabaseQueryResponse struct {
	Object  string `json:"object"`
	Results []struct {
		Object         string    `json:"object"`
		ID             string    `json:"id"`
		CreatedTime    time.Time `json:"created_time"`
		LastEditedTime time.Time `json:"last_edited_time"`
		Parent         struct {
			Type       string `json:"type"`
			DatabaseID string `json:"database_id"`
		} `json:"parent"`
		Archived   bool `json:"archived"`
		Properties struct {
			Property struct {
				ID       string        `json:"id"`
				Type     string        `json:"type"`
				RichText []interface{} `json:"rich_text"`
			} `json:"Property"`
			GCalID struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				RichText []struct {
					Type string `json:"type"`
					Text struct {
						Content string      `json:"content"`
						Link    interface{} `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string      `json:"plain_text"`
					Href      interface{} `json:"href"`
				} `json:"rich_text"`
			} `json:"GCal_ID"`
			Date struct {
				ID          string    `json:"id"`
				Type        string    `json:"type"`
				CreatedTime time.Time `json:"created_time"`
			} `json:"Date"`
			Name struct {
				ID    string `json:"id"`
				Type  string `json:"type"`
				Title []struct {
					Type string `json:"type"`
					Text struct {
						Content string      `json:"content"`
						Link    interface{} `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string      `json:"plain_text"`
					Href      interface{} `json:"href"`
				} `json:"title"`
			} `json:"Name"`
		} `json:"properties"`
		URL string `json:"url"`
	} `json:"results"`
	NextCursor interface{} `json:"next_cursor"`
	HasMore    bool        `json:"has_more"`
}
