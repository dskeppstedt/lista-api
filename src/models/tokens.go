package models

type UserTokens struct {
	Auth    string `json:"auth"`
	Refresh string `json:"refresh"`
}

func NewUserTokens(auth, refresh string) *UserTokens {
	ut := &UserTokens{auth, refresh}
	return ut
}
