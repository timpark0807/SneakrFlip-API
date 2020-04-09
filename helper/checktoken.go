package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/timpark0807/restapi/model"
)

// CheckToken Comment
func CheckToken(token string) (model.Property, error) {

	var bearerToken model.BearerToken
	var property model.Property

	if len(token) == 0 {
		return property, errors.New("Invalid")
	}

	authToken := strings.Split(token, " ")[1]
	googleURL := "https://www.googleapis.com/oauth2/v1/tokeninfo?access_token="
	resp, err := http.Get(googleURL + authToken)
	fmt.Println(resp)
	if err != nil {
		return property, errors.New("Bad Request")
	}

	_ = json.NewDecoder(resp.Body).Decode(&bearerToken)
	if len(bearerToken.Error) != 0 {
		fmt.Println(bearerToken.Error)
		return property, errors.New(bearerToken.ErrorDescription)
	}
	property.CreatedBy = bearerToken.Email

	return property, nil
}
