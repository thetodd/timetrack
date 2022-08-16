package domain

import "time"

type TimeEntries struct {
	TimeEntries []TimeEntry
}

type TimeEntry struct {
	Start, End  time.Time
	ProjectId   int
	Description string
}
