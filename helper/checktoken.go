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
func CheckToken(token string) (model.BearerToken, error) {

	var bearerToken model.BearerToken

	if len(token) == 0 {
		return bearerToken, errors.New("Invalid")
	}

	authToken := strings.Split(token, " ")[1]
	googleURL := "https://www.googleapis.com/oauth2/v1/tokeninfo?access_token="
	resp, err := http.Get(googleURL + authToken)

	if err != nil {
		return bearerToken, errors.New("Bad Request")
	}

	_ = json.NewDecoder(resp.Body).Decode(&bearerToken)

	if len(bearerToken.Error) != 0 {
		fmt.Println(bearerToken.Error)
		return bearerToken, errors.New(bearerToken.ErrorDescription)
	}

	return bearerToken, nil
}
