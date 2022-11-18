package groupings

import "time"

type Grouping struct {
}

type PomodoroGrouping struct {
	// in minutes
	mainLength      time.Duration
	breakLength     time.Duration
	longBreakLength time.Duration
	numberIntervals int16
}
