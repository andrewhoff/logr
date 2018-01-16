package writers

import (
	"time"

	"github.com/andrewhoff/logr/store"
	"github.com/fatih/color"
)

// LazyLogWriter ...
type LazyLogWriter struct {
	aggregateMsg string
	priority     int
}

// NewLazyLogWriter ...
func NewLazyLogWriter(priority int) *LazyLogWriter {
	return &LazyLogWriter{
		priority: priority,
	}
}

// Log - this implementation logs string message along with a specified priority
func (lw *LazyLogWriter) Log(val string) {
	lw.aggregateMsg += val
}

// Flush - Actually send the log to the system
func (lw *LazyLogWriter) Flush() error {
	_, err := store.InternalDataStore.Enqueue(&store.Item{
		Priority: lw.priority,
		Value:    lw.aggregateMsg,
		DateTime: time.Now(),
	})

	if err == nil {
		lw.reset()
		color.Green("[LAZY LOG WRITER] - Write Succeeded!")
	}

	return err
}

func (lw *LazyLogWriter) reset() {
	lw.aggregateMsg = ""
}
