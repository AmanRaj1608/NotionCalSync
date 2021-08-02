package main

import (
	"NotionCalSync/lib"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func getEnvVariable() {
	err := godotenv.Load("./config.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	getEnvVariable()

	// Instance getCalendar
	svc := lib.GetCalendar()

	// Query the full list of calendars
	c_list, err := svc.CalendarList.List().Fields("items/id").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve list of calendars: %v", err)
	}

	// For each calendar in calendars save summary and id
	calendars := []calendarType{}
	for _, v := range c_list.Items {
		// fmt.Printf("GCal_ID - %v \nSummary - %v\n", v.Id, v.Summary)
		calendars = append(calendars, calendarType{v.Id, v.Summary, []eventItem{}})
	}

	// Get eventItem for each calendar in calendars
	for ind, calendar := range calendars {
		// fmt.Printf("Fetching eventItem for GCal_ID - %v\n", calendar.c_id)
		// filter only upcoming events for today
		t1 := time.Now().Format(time.RFC3339)
		t2 := time.Now().AddDate(0, 0, 5).Format(time.RFC3339)
		events, _ := svc.Events.List(calendar.c_id).ShowDeleted(false).
			SingleEvents(true).TimeMin(t1).TimeMax(t2).MaxResults(15).OrderBy("startTime").Do()

		// fmt.Println("Upcoming events...")
		if len(events.Items) != 0 {
			for _, item := range events.Items {
				date := item.Start.DateTime
				if date == "" {
					date = item.Start.Date
				}
				// fmt.Println(item.Id, item.Summary, date)
				calendars[ind].events = append(calendars[ind].events, eventItem{item.Id, item.Summary, date})
			}
		}
	}

	// Init Notion client
	// lib.GetNotion(notion_api)
	// Query entries from notion db
	notionToday := lib.QueryDB()
	cal_today := []eventItem{}
	// filter items not in notionDb using maps
	mp := make(map[string]bool)
	for _, u := range notionToday.Results {
		mp[u.ID] = true
	}
	for _, gCal := range calendars {
		for _, eventList := range gCal.events {
			fmt.Println(eventList.id, mp[eventList.id])
			if !mp[eventList.id] {
				cal_today = append(cal_today, eventList)
			}
		}
	}
	fmt.Println(cal_today)
	// lib.GetPage("fd120265ab174d8aa19d0ae4d0db44ca")

	// Create page for every new events
	for _, cal := range cal_today {
		// Create page
		lib.CreatePage(cal.id, cal.summary, cal.date)
	}
}

type calendarType struct {
	c_id    string
	summary string
	events  []eventItem
}

type eventItem struct {
	id      string
	summary string
	date    string
}
