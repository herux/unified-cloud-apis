package authentication

type Authenticator interface {
	Init() error
}
