package alerter

import (
	"context"
	"fmt"
	"sync"

	"firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
	"log"
)

type FCM struct {
	app *firebase.App
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

func (f *FCM) SendAlert(alert messaging.Message, errc chan []error, wg *sync.WaitGroup) {
	var errs []error
	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := f.app.Messaging(ctx)

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, &alert)
	if err != nil {
		errs = append(errs, err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

	errc <- errs

	wg.Done()
}
