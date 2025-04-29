package authentication

import "github.com/aws/aws-sdk-go/aws/session"

type AWSAuthenticator struct {
	Session *session.Session
}

func (a *AWSAuthenticator) Init() error {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}

	a.Session = sess

	return nil
}
