package authentication

type PasswordResetEmailer interface {
	Email(email string, token string, valid string) error
}
