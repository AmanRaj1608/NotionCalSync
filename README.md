<p align="center"><img src="" align="center" width="350"></p>
<h2 align="center">Notion Calendar Sync</h2>
<p align="center"><b>Notion <> Google Calendar two-way sync</b></p>

A two way sync between Google calendar and Notion Pages, db. Fetch all tasks from all calendars for next 24hrs (or anything you set) and update them into Notion todo page.

### To run

```bash
git clone https://github.com/AmanRaj1608/NotionCalSync.git
go get
go run main.go
```

### Usage Guide

- Update the `config/credentials.json` with [Create Credentials](https://developers.google.com/workspace/guides/create-credentials).
- Update `config.env` with [Notion Secrets](https://www.notion.so/my-integrations).
- Now, when you run the program for first time, you will be redirected to get the tokens.
- Once you enter the token it will be saved in `config/token.json` for reuse.
- Now it will fetch all the events from the Calendar and update the Notion Page.
