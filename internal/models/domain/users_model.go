package domain

type User struct {
	ID             int
	UserName       string
	Email          string
	EmailNew       string
	Password       string
	PasswordNew    string
	PasswordHash   string
	SecretWord     string
	SecretWordNew  string
	SecretWordHash string
	CreatedAt      string
}
