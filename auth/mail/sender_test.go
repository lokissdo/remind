package mail

import (
	"github.com/elliotnguyen/auth/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSenderEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailConfig.EmailSenderName, config.EmailConfig.EmailSenderAddress, config.EmailConfig.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from <a href="https://github.com/elliotnguyen">Elliot Nguyen</a></p>
	`
	to := []string{"nguyentienthanhlqd@gmail.com"}
	attachFile := []string{"../../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFile)
	require.NoError(t, err)
}
