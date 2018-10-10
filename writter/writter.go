package writter

import (
	"firebase.google.com/go/messaging"
	"github.com/pamungkaski/camar/datamodel"
)

type Writer struct{}

func (a *Writer) CreateAlertMessage(disaster datamodel.GeoJSON, alerts []string) (messaging.Message, error) {
	// Cut AlertBody unnecessary part
	alertBody := disaster.Properties.Title

	message := messaging.Message{
		Data: map[string]string{
			"URL": disaster.URL,
		},
		Notification: &messaging.Notification{
			Title: "Warning",
			Body:  alertBody,
		},
	}
	return message, nil
}
