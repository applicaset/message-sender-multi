package messagesendermulti_test

import (
	"context"
	"github.com/applicaset/message-sender-multi"
	"github.com/applicaset/message-sender-output"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMessageSender_Send(t *testing.T) {
	ms1 := messagesenderoutput.New()
	mms := messagesendermulti.New(ms1)

	ctx := context.Background()
	phoneNumber := "+1234567890"
	msg := "Dummy Text"

	err := mms.Send(ctx, phoneNumber, msg)
	require.NoError(t, err)
}
