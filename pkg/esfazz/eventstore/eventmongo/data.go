package eventmongo

import "github.com/mikaelim-id/go-apt/pkg/esfazz"

type eventLog struct {
	Type      string
	Aggregate esfazz.BaseAggregate
	Data      map[string]interface{}
}
