// models/user.go
package models

type User struct {

	ID int `json:"id"`
	FIRST_NAME string `json:"first_name"`
    LAST_NAME string `json:"last_name"`
	EMAIL string `json:"email"`
	COUNTRY_CODE string `json:"country_code"`
    MOBILE  string `json:"mobile"`
}
