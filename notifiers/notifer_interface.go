package notifiers

type INotifier interface {
	SendMessage(string, string) error
}
