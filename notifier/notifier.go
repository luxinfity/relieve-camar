package notifier

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type FCM struct {
	app *firebase.App
}

// Alerting is the business logic contract for alerting service.
// the main idea is to send alert to all device.
type Notifier interface {
	// SendAlert is a function to send Disastrous Event alert to specific Device using the alerting service.
	SendAlert(disaster interface{}, token string, errc chan error)
}

func NewAlerter() *FCM {
	opt := option.WithCredentialsFile("FCM.json")
	fb, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return &FCM{
		app: fb,
	}
}

func (f *FCM) SendAlert(disaster interface{}, token string, errc chan error) {
	alert := f.createAlertMessage(disaster, token)

	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := f.app.Messaging(ctx)

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, &alert)
	if err != nil {
		errc <- err
		return
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

	errc <- nil
}

func (f *FCM) createAlertMessage(disaster interface{}, token string) messaging.Message {
	message := messaging.Message{
		Data: map[string]string{
			"data": fmt.Sprintf("%b", disaster),
		},
		Notification: &messaging.Notification{
			Title: "Warning",
			Body:  "Disaster just happen in your area",
		},
		Token: token,
	}
	return message
}
