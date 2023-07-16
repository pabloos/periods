package periods

import (
	"sync"
	"time"
)

type Report struct {
	sync.Mutex

	hints []time.Time
}

func New() *Report {
	return &Report{}
}

func (report *Report) Hit() {
	defer report.Unlock()
	report.Lock()

	report.hints = append(report.hints, time.Now())
}

func (report *Report) Since(duration time.Duration) int {
	defer report.Unlock()
	report.Lock()

	var hints int

	for _, hint := range report.hints {
		if time.Since(hint) < duration {
			hints++
		}
	}

	return hints
}

func (report *Report) Length() int {
	defer report.Unlock()
	report.Lock()

	return len(report.hints)
}
