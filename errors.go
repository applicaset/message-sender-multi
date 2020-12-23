package messagesendermulti

type ErrSendMessageFailed struct {
	Errors []error
}

func (err ErrSendMessageFailed) Error() string {
	return "sending message with all registered drivers failed"
}
