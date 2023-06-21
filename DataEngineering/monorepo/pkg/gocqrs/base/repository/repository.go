package repository

import (
	"api_thienhang_com/pkg/gocqrs/base/aggregate"
	"api_thienhang_com/pkg/gocqrs/base/event"
)

// Repository is responsible to generate an Aggregate
// save events and publish it
type Repository struct {
	eventStore event.EventStore
	eventBus   event.EventBus
}

// NewRepository creates a repository wieh a eventstore and eventbus access
func NewRepository(store event.EventStore, bus event.EventBus) *Repository {
	return &Repository{
		store,
		bus,
	}
}

// Load restore the last state of an aggregate
func (r *Repository) Load(agg aggregate.AggregateHandler, ID string) error {
	events, err := r.eventStore.Load(ID)

	if err != nil {
		return err
	}

	for _, event := range events {
		agg.ApplyChangeHelper(agg, event, false)
	}
	return nil
}

// Save the events and publish it to eventbus
func (r *Repository) Save(aggregate aggregate.AggregateHandler, version int) error {
	return r.eventStore.Save(aggregate.Uncommited(), version)
}

// PublishEvents to an eventBus
func (r *Repository) PublishEvents(aggregate aggregate.AggregateHandler, bucket, subset string) error {
	var err error

	for _, event := range aggregate.Uncommited() {
		if err = r.eventBus.Publish(event, bucket, subset); err != nil {
			return err
		}
	}

	return nil
}

// SafeSave the events without check the version
func (r *Repository) SafeSave(aggregate aggregate.AggregateHandler, version int) error {
	return r.eventStore.SafeSave(aggregate.Uncommited(), version)
}
