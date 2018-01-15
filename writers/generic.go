package writers

import (
	"time"

	"github.com/andrewhoff/logr/store"
	"github.com/fatih/color"
)

// GenericLogWriter ...
type GenericLogWriter struct{}

// NewGenericLogWriter ...
func NewGenericLogWriter() *GenericLogWriter {
	return &GenericLogWriter{}
}

// Log - this implementation logs string message along with a specified priority
func (lw *GenericLogWriter) Log(priority int, val string) error {
	_, err := store.InternalDataStore.Enqueue(&store.Item{
		Priority: priority,
		Value:    val,
		DateTime: time.Now(),
	})

	if err == nil {
		color.Green("[GENERIC LOG WRITER] - Write Succeeded!")
	}

	return err
}
