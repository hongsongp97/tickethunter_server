package helpers

import (
	"log"

	fcm "github.com/appleboy/go-fcm"
)

type FcmHelper struct {
}

func (fcmHelper *FcmHelper) Init() {
	log.SetFlags(log.Lshortfile)
}

// CreateMessage func creates the message to be sent.
func (fcmHelper *FcmHelper) CreateMessage() *(fcm.Message) {
	msg := &fcm.Message{
		To: "sample_device_token",
		Data: map[string]interface{}{
			"foo": "bar",
		},
	}

	return msg
}

// Create a FCM client to send the message.
func (fcmHelper *FcmHelper) CreateClient() (*(fcm.Client), error) {
	client, err := fcm.NewClient("sample_api_key")
	if err != nil {
		log.Fatalln(err)
	}

	return client, err
}

// Send the message and receive the response without retries.
func (fcmHelper *FcmHelper) SendMessage(client *(fcm.Client), msg *(fcm.Message)) (*(fcm.Response), error) {
	response, err := client.Send(msg)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%#v\n", response)

	return response, err
}
