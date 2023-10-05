package eventstore

import (
	"context"
	"github.com/mikaelim-id/go-apt/pkg/esfazz"
)

// EventStore is interface for event store
type EventStore interface {
	Save(ctx context.Context, agg esfazz.Aggregate, events ...*esfazz.EventPayload) ([]*esfazz.Event, error)
	FindNotApplied(ctx context.Context, agg esfazz.Aggregate) ([]*esfazz.Event, error)
}
