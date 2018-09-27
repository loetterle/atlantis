package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"
	events "github.com/runatlantis/atlantis/server/events"
)

func AnyEventsCommandName() events.CommandName {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(events.CommandName))(nil)).Elem()))
	var nullValue events.CommandName
	return nullValue
}

func EqEventsCommandName(value events.CommandName) events.CommandName {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue events.CommandName
	return nullValue
}