package store

import (
	"testing"
	"time"

	"github.com/andrewhoff/logr/config"
)

func TestInit(t *testing.T) {
	Init()

	if InternalDataStore.ds == nil {
		t.Fail()
	}

	if InternalDataStore.mutex == nil {
		t.Fail()
	}
}

func TestEnqueue(t *testing.T) {
	Init()
	config.SetOpts(&config.Opts{})

	item := &Item{
		Priority: config.LowPriority,
		Value:    "hi",
		DateTime: time.Now(),
	}

	enqueued, err := InternalDataStore.Enqueue(item)
	if err != nil {
		t.Error(err)
	}

	if enqueued != item {
		t.Fail()
	}

	if InternalDataStore.Len() != 1 {
		t.Fail()
	}

	if InternalDataStore.LenWithPriority(config.LowPriority) != 1 {
		t.Fail()
	}

	if InternalDataStore.Empty() {
		t.Fail()
	}

}

func TestEnqueueCapOne(t *testing.T) {
	Init()
	config.SetOpts(&config.Opts{Capacity: 1})

	item := &Item{
		Priority: config.LowPriority,
		Value:    "hi",
		DateTime: time.Now(),
	}

	enqueued, err := InternalDataStore.Enqueue(item)
	if err != nil {
		t.Error(err)
	}

	if enqueued != item {
		t.Fail()
	}

	if InternalDataStore.Len() != 1 {
		t.Fail()
	}

	if InternalDataStore.LenWithPriority(config.LowPriority) != 1 {
		t.Fail()
	}

	if InternalDataStore.Empty() {
		t.Fail()
	}

}

func TestEnqueueCapZero(t *testing.T) {
	Init()
	config.SetOpts(&config.Opts{Capacity: 0})

	item := &Item{
		Priority: config.LowPriority,
		Value:    "hi",
		DateTime: time.Now(),
	}

	enqueued, err := InternalDataStore.Enqueue(item)
	if err != nil {
		t.Error(err)
	}

	if enqueued != item {
		t.Fail()
	}

	if InternalDataStore.Len() != 1 {
		t.Fail()
	}

	if InternalDataStore.LenWithPriority(config.LowPriority) != 1 {
		t.Fail()
	}

	if InternalDataStore.Empty() {
		t.Fail()
	}
}

func TestDequeue(t *testing.T) {
	Init()
	config.SetOpts(&config.Opts{})

	item := &Item{
		Priority: config.LowPriority,
		Value:    "hi",
		DateTime: time.Now(),
	}

	_, err := InternalDataStore.Enqueue(item)
	if err != nil {
		t.Error(err)
	}

	dequeued := InternalDataStore.Dequeue()
	if dequeued != item {
		t.Fail()
	}

	if InternalDataStore.Len() != 0 {
		t.Fail()
	}

	if InternalDataStore.LenWithPriority(config.LowPriority) != 0 {
		t.Fail()
	}

	if !InternalDataStore.Empty() {
		t.Fail()
	}
}

func TestDequeueOldestLowestIsLowest(t *testing.T) {
	Init()
	config.SetOpts(&config.Opts{})

	item1 := &Item{
		Priority: config.LowPriority,
		Value:    "hi",
		DateTime: time.Now(),
	}

	item2 := &Item{
		Priority: config.MedPriority,
		Value:    "hi",
		DateTime: time.Now(),
	}

	_, err := InternalDataStore.Enqueue(item1)
	if err != nil {
		t.Error(err)
	}

	_, err = InternalDataStore.Enqueue(item2)
	if err != nil {
		t.Error(err)
	}

	dequeued := InternalDataStore.DequeueOldestLowest()
	if dequeued != item1 {
		t.Fail()
	}

	if InternalDataStore.Len() != 1 {
		t.Fail()
	}

	if InternalDataStore.LenWithPriority(config.MedPriority) == 0 {
		t.Fail()
	}

	if InternalDataStore.Empty() {
		t.Fail()
	}
}

func TestDequeueOldestLowestIsOldest(t *testing.T) {
	Init()
	config.SetOpts(&config.Opts{})

	item1 := &Item{
		Priority: config.LowPriority,
		Value:    "I'm the oldest log message",
		DateTime: time.Now(),
	}

	item2 := &Item{
		Priority: config.LowPriority,
		Value:    "hi",
		DateTime: time.Now(),
	}

	_, err := InternalDataStore.Enqueue(item1)
	if err != nil {
		t.Error(err)
	}

	_, err = InternalDataStore.Enqueue(item2)
	if err != nil {
		t.Error(err)
	}

	dequeued := InternalDataStore.DequeueOldestLowest()
	if dequeued != item1 {
		t.Fail()
	}

	if InternalDataStore.Len() != 1 {
		t.Fail()
	}

	if InternalDataStore.LenWithPriority(config.LowPriority) != 1 {
		t.Fail()
	}

	if InternalDataStore.Empty() {
		t.Fail()
	}
}

func TestEnqueueFullWithOverwrites(t *testing.T) {
	Init()
	config.SetOpts(&config.Opts{
		Capacity:  1,
		Overwrite: true,
	})

	item1 := &Item{
		Priority: config.LowPriority,
		Value:    "I'm the oldest log message",
		DateTime: time.Now(),
	}

	item2 := &Item{
		Priority: config.LowPriority,
		Value:    "hi",
		DateTime: time.Now(),
	}

	_, err := InternalDataStore.Enqueue(item1)
	if err != nil {
		t.Error(err)
	}

	_, err = InternalDataStore.Enqueue(item2)
	if err != nil {
		t.Error(err)
	}

	if InternalDataStore.Len() > 1 {
		t.Fail()
	}

	dequeued := InternalDataStore.Dequeue()
	if dequeued != item2 {
		t.Fail()
	}

	if !InternalDataStore.Empty() {
		t.Fail()
	}
}

func TestEnqueueFullWithoutOverwrites(t *testing.T) {
	Init()
	config.SetOpts(&config.Opts{
		Capacity:  1,
		Overwrite: false,
	})

	item1 := &Item{
		Priority: config.LowPriority,
		Value:    "I'm the oldest log message",
		DateTime: time.Now(),
	}

	item2 := &Item{
		Priority: config.LowPriority,
		Value:    "hi",
		DateTime: time.Now(),
	}

	_, err := InternalDataStore.Enqueue(item1)
	if err != nil {
		t.Error(err)
	}

	_, err = InternalDataStore.Enqueue(item2)
	if err == nil {
		//it's supposed to return an error here
		t.Fail()
	}

	if InternalDataStore.Len() > 1 {
		t.Fail()
	}

	dequeued := InternalDataStore.Dequeue()
	if dequeued != item1 {
		t.Fail()
	}

	if !InternalDataStore.Empty() {
		t.Fail()
	}
}

func TestDequeueEmpty(t *testing.T) {
	Init()
	config.SetOpts(&config.Opts{})

	item1 := &Item{
		Priority: config.LowPriority,
		Value:    "I'm the oldest log message",
		DateTime: time.Now(),
	}

	_, err := InternalDataStore.Enqueue(item1)
	if err != nil {
		t.Error(err)
	}

	if InternalDataStore.Len() > 1 {
		t.Fail()
	}

	dequeued := InternalDataStore.Dequeue()
	if dequeued != item1 {
		t.Fail()
	}

	if !InternalDataStore.Empty() {
		t.Fail()
	}

	dequeuedAgain := InternalDataStore.Dequeue()
	if dequeuedAgain != nil {
		t.Fail()
	}
}

func TestLenWithPriority(t *testing.T) {
	Init()
	config.SetOpts(&config.Opts{})

	item1 := &Item{
		Priority: config.LowPriority,
		Value:    "I'm the oldest log message",
		DateTime: time.Now(),
	}
	_, err := InternalDataStore.Enqueue(item1)
	if err != nil {
		t.Error(err)
	}

	item2 := &Item{
		Priority: config.LowPriority,
		Value:    "I'm the second log message",
		DateTime: time.Now(),
	}
	_, err = InternalDataStore.Enqueue(item2)
	if err != nil {
		t.Error(err)
	}

	item3 := &Item{
		Priority: config.MedPriority,
		Value:    "Med Pri",
		DateTime: time.Now(),
	}
	_, err = InternalDataStore.Enqueue(item3)
	if err != nil {
		t.Error(err)
	}

	item4 := &Item{
		Priority: config.HighPriority,
		Value:    "High Pri",
		DateTime: time.Now(),
	}
	_, err = InternalDataStore.Enqueue(item4)
	if err != nil {
		t.Error(err)
	}

	if InternalDataStore.LenWithPriority(config.LowPriority) != 2 {
		t.Fail()
	}

	if InternalDataStore.LenWithPriority(config.MedPriority) != 1 {
		t.Fail()
	}

	if InternalDataStore.LenWithPriority(config.HighPriority) != 1 {
		t.Fail()
	}
}
