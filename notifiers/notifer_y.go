package notifiers

import "log"

type NotifierY struct{}

func (n NotifierY) SendMessage(mobileNumber string, message string) error {
	log.Printf("Sending message '%s' to (%s) through Notifier Y", message, mobileNumber)
	return nil
}
