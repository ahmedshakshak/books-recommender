package notifiers

import "os"

func NewNotifier() INotifier {
	isNotifierX := os.Getenv("NOTIFIER") == "X"
	if isNotifierX {
		return &NotifierX{}
	} else {
		return &NotifierY{}
	}
}
