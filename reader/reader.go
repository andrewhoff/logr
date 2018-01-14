package reader

import (
	"log"

	"github.com/andrewhoff/logr/config"
	"github.com/andrewhoff/logr/store"
	"github.com/fatih/color"
)

// LogReader ...
type LogReader struct {
}

// NewLogReader ...
func NewLogReader() *LogReader {
	return &LogReader{}
}

// Get ...
func (reader *LogReader) Get() string {
	if store.InternalDataStore == nil {
		log.Println("[WARNING] logr hasn't finished initializing")
		return ""
	}

	if store.InternalDataStore.Empty() {
		log.Println("[WARNING] log is empty")
		return ""
	}

	dequeued := store.InternalDataStore.Dequeue()
	if dequeued == nil {
		return ""
	}

	switch dequeued.Priority {
	case config.HighPriority:
		return color.GreenString("[HIGH PRIORITY] - %s\n", dequeued.Value)
	case config.MedPriority:
		return color.YellowString("[MED PRIORITY] - %s\n", dequeued.Value)
	case config.LowPriority:
		return color.RedString("[LOW PRIORITY] - %s\n", dequeued.Value)
	}

	return dequeued.Value
}
