package lib

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dstotijn/go-notion"
)

var client *notion.Client

func GetNotion() *notion.Client {
	notion_api := os.Getenv("notion_api")
	fmt.Println(notion_api)
	client := notion.NewClient(notion_api)
	return client
}

func GetPage(pageId string) {
	// view page
	fmt.Println("page print")
	page, err := client.FindPageByID(context.Background(), pageId)
	if err != nil {
		log.Fatalf("Error while getting page %v", err)
	}
	fmt.Println("print..")
	fmt.Println(page)
}

func queryDB(database_id string) (result notion.DatabaseQueryResponse, err error) {
	t := time.Now()
	fmt.Println("here")
	res, err := client.QueryDatabase(context.Background(), database_id, &notion.DatabaseQuery{
		Filter: &notion.DatabaseQueryFilter{
			Property: "Date",
			Date: &notion.DateDatabaseQueryFilter{
				Equals: &t,
			},
		},
	})
	fmt.Println("after")
	if err != nil {
		noData := notion.DatabaseQueryResponse{
			Results:    []notion.Page{},
			HasMore:    false,
			NextCursor: new(string),
		}
		fmt.Println(fmt.Printf("%v", err))
		return noData, err
	}
	fmt.Println(fmt.Printf("%v", res))
	return res, nil
}

func Createpage_() {
	fmt.Printf("createPage\n")

	parentID := os.Getenv("database_id")

	ctx := context.Background()
	titleText := &notion.Text{
		Content: "this is a title of the created page",
	}
	title := []notion.RichText{
		{
			Type: notion.RichTextTypeText,
			Text: titleText,
		},
	}
	params := notion.CreatePageParams{
		ParentID:   parentID,
		ParentType: notion.ParentTypePage,
		Title:      title,
	}
	rsp, err := client.CreatePage(ctx, params)
	if err != nil {
		log.Fatalf("CreatePage() failed with '%s'\n", err)
	}
	fmt.Println("Created a page!\n")
	fmt.Println(rsp)
}

// func ShowPageInfo(page *notion.Page) {
// 	fmt.Printf("ShowPageInfo:\n")
// 	fmt.Printf("  ID: '%s'\n", page.ID)
// 	fmt.Printf("  CreatedTime: '%s'\n", page.CreatedTime)
// 	fmt.Printf("  LastEditedTime: '%s'\n", page.LastEditedTime)
// 	if page.Parent.PageID != nil {
// 		fmt.Printf("  Parent: page with ID '%s'\n", *page.Parent.PageID)
// 	} else if page.Parent.DatabaseID != nil {
// 		fmt.Printf("  Parent: database with ID '%s'\n", *page.Parent.DatabaseID)
// 	} else {
// 		fmt.Printf("both page.Parent.PageID or page.Parent.DatabaseID are nil")
// 	}
// 	fmt.Printf("  Archived: %v\n", page.Archived)
// 	switch prop := page.Properties.(type) {
// 	case notion.PageProperties:
// 		fmt.Printf("  page properties:\n")
// 		ShowRichText(2, "Title", prop.Title.Title)
// 	case notion.DatabasePageProperties:
// 		fmt.Printf("  database properties (NYI):\n")
// 	}
// }
