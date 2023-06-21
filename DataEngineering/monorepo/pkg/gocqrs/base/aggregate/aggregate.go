// The aggregate is a logical boundary for things that can change in a business transaction of a given context.
// In the Eventhus context, it simplifies the process the commands and produce events.
package aggregate

import (
	"time"

	"api_thienhang_com/pkg/gocqrs/base/command"
	"api_thienhang_com/pkg/gocqrs/base/event"
)

// BaseAggregate contains the basic info
// that all aggregates should have
type BaseAggregate struct {
	ID        string
	Type      string
	Version   int
	Changes   []event.Event
	CreatedAt time.Time
	CreatedBy string
	CreatedFB string
}

// AggregateHandler defines the methods to process commands
type AggregateHandler interface {
	// LoadsFromHistory(events []Event)
	ApplyChange(event event.Event)
	ApplyChangeHelper(aggregate AggregateHandler, event event.Event, commit bool)
	HandleCommand(command.Command) error
	Uncommited() []event.Event
	ClearUncommited()
	IncrementVersion()
	GetID() string
}

// Uncommited return the events to be saved
func (b *BaseAggregate) Uncommited() []event.Event {
	return b.Changes
}

// ClearUncommited the events
func (b *BaseAggregate) ClearUncommited() {
	b.Changes = []event.Event{}
}

// IncrementVersion ads 1 to the current version
func (b *BaseAggregate) IncrementVersion() {
	b.Version++
}

// ApplyChangeHelper increments the version of an aggregate and apply the change itself
func (b *BaseAggregate) ApplyChangeHelper(aggregate AggregateHandler, evt event.Event, commit bool) {
	// increments the version in event and aggregate
	b.IncrementVersion()

	// apply the event itself
	aggregate.ApplyChange(evt)
	if commit {
		evt.Version = b.Version
		_, evt.Type = event.GetTypeName(evt.Data)
		b.Changes = append(b.Changes, evt)
	}
}

// GetID of the current aggregate
func (b *BaseAggregate) GetID() string {
	return b.ID
}

func (b *BaseAggregate) GetCreatedBy() string {
	return b.CreatedBy
}

func (b *BaseAggregate) GetCreatedFB() string {
	return b.CreatedFB
}
