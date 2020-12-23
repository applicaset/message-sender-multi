package messagesendermulti

import (
	"context"
	"github.com/applicaset/sms-svc"
)

type messageSender struct {
	drivers []smssvc.MessageSender
}

func (ms *messageSender) Send(ctx context.Context, phoneNumber, message string) error {
	errs := make([]error, 0)

	for i := range ms.drivers {
		err := ms.drivers[i].Send(ctx, phoneNumber, message)
		if err != nil {
			errs = append(errs, err)

			continue
		}

		return nil
	}

	return ErrSendMessageFailed{Errors: errs}
}

func New(driver1 smssvc.MessageSender, drivers ...smssvc.MessageSender) smssvc.MessageSender {
	ms := messageSender{
		drivers: append([]smssvc.MessageSender{driver1}, drivers...),
	}

	return &ms
}
