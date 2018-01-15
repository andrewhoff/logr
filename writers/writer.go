package writers

import (
	"time"

	"github.com/andrewhoff/logr/config"
	"github.com/andrewhoff/logr/store"
)

// GenericLogWriter ...
type GenericLogWriter struct {
}

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

	return err
}

// SevereLogWriter - this type of LogWriter will always try to log messages with the highest priority, automatically and specified priority needed
type SevereLogWriter struct {
}

// NewSevereLogWriter ...
func NewSevereLogWriter() *SevereLogWriter {
	return &SevereLogWriter{}
}

// Log - this implementation will always try to log messages with the highest priority, automatically and specified priority needed
func (lw *SevereLogWriter) Log(val string) error {
	_, err := store.InternalDataStore.Enqueue(&store.Item{
		Priority: config.HighPriority,
		Value:    val,
		DateTime: time.Now(),
	})

	return err
}
