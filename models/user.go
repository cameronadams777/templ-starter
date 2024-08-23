package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type User struct {
	UUIDBaseModel
	FirstName           string       `json:"first_name"`
	LastName            string       `json:"last_name"`
	Email               string       `json:"email"`
	Password            string       `json:"password"`
	ConfirmedAt         *time.Time    `json:"confirmed_at"`
	ConfirmationToken   string       `json:"confirmation_token"`
	ResetPasswordCode   string       `json:"reset_password_code"`
	ResetPasswordExpiry sql.NullTime `json:"reset_password_expiry"`
}

func (u User) MarshalJSON() ([]byte, error) {
	type user User // prevent recursion
	x := user(u)
	x.Password = ""
	return json.Marshal(x)
}

func (u User) IsConfirmed() bool {
	return u.ConfirmedAt != nil
}
