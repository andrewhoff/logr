package logr

import (
	"github.com/andrewhoff/logr/config"
	"github.com/andrewhoff/logr/reader"
	"github.com/andrewhoff/logr/store"
	"github.com/andrewhoff/logr/writers"
)

const (
	// LowPriority ...
	LowPriority = config.LowPriority

	// MedPriority ...
	MedPriority = config.MedPriority

	// HighPriority ...
	HighPriority = config.HighPriority
)

// InitWithOpts ...
func InitWithOpts(opts *config.Opts) {
	store.Init()
	config.SetOpts(opts)
}

// NewLogReader ...
func NewLogReader() (*reader.LogReader, error) {
	return reader.NewLogReader(), nil
}

// NewGenericLogWriter ...
func NewGenericLogWriter() (*writers.GenericLogWriter, error) {
	return writers.NewGenericLogWriter(), nil
}

// NewSevereLogWriter ...
func NewSevereLogWriter() (*writers.SevereLogWriter, error) {
	return writers.NewSevereLogWriter(), nil
}
