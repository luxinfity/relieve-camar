package helper

import "github.com/pamungkaski/camar/datamodel"

func ResponsifyEvent(event datamodel.Event) (datamodel.ResponseEvent, error) {
	return datamodel.ResponseEvent{
		ID: event.ID,
	}
}
