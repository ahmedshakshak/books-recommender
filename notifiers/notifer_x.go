package notifiers

import "log"

type NotifierX struct{}

func (n NotifierX) SendMessage(mobileNumber string, message string) error {
	log.Printf("Sending message '%s' to (%s) through Notifier X", message, mobileNumber)
	return nil
}
