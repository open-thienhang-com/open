package base

import (
	"errors"
	"reflect"

	"api_thienhang_com/pkg/gocqrs/base/aggregate"
	"api_thienhang_com/pkg/gocqrs/base/command"
	"api_thienhang_com/pkg/gocqrs/base/repository"
)

// ErrInvalidID missing initial event
var ErrInvalidID = errors.New("Invalid ID, initial event missign")

// Handler contains the info to manage commands
type Handler struct {
	repository     *repository.Repository
	aggregate      reflect.Type
	bucket, subset string
}

// NewCommandHandler return a handler
func NewCommandHandler(repository *repository.Repository, aggregate aggregate.AggregateHandler) command.CommandHandle {
	return &Handler{
		repository: repository,
		aggregate:  reflect.TypeOf(aggregate).Elem(),
	}
}

// Handle a command
func (h *Handler) Handle(cmd command.Command) error {
	var err error

	version := cmd.GetVersion()
	aggregate := reflect.New(h.aggregate).Interface().(aggregate.AggregateHandler)

	if version != 0 {
		if err = h.repository.Load(aggregate, cmd.GetAggregateID()); err != nil {
			return err
		}
	}

	if err = aggregate.HandleCommand(cmd); err != nil {
		return err
	}

	// if not contain a valid ID,  the initial event (some like createAggreagate event) is missing
	if aggregate.GetID() == "" {
		return ErrInvalidID
	}

	if err = h.repository.Save(aggregate, version); err != nil {
		return err
	}

	if err = h.repository.PublishEvents(aggregate, h.bucket, h.subset); err != nil {
		return err
	}

	return nil
}
