package reader

import (
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
		color.Yellow("[WARNING] logr hasn't finished initializing")
		return ""
	}

	if store.InternalDataStore.Empty() {
		color.Yellow("[WARNING] log is empty")
		return ""
	}

	dequeued := store.InternalDataStore.Dequeue()
	if dequeued == nil {
		return ""
	}

	switch dequeued.Priority {
	case config.HighPriority:
		return color.RedString("[HIGH PRIORITY] - %s\n", dequeued.Value)
	case config.MedPriority:
		return color.YellowString("[MED PRIORITY] - %s\n", dequeued.Value)
	case config.LowPriority:
		return color.GreenString("[LOW PRIORITY] - %s\n", dequeued.Value)
	}

	return dequeued.Value
}
