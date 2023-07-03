package main

import (
	"encoding/json"
	"errors"
	env "github.com/logotipiwe/dc_go_env_lib"
	"net/http"
	"net/url"
)

type User struct {
	Sub        string `json:"sub"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	Locale     string `json:"locale"`
}

func GetAccessTokenFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return "", err
	}
	var accessToken string
	if cookie != nil {
		accessToken = cookie.Value
	} else {
		accessToken = ""
	}
	return accessToken, nil
}

func GetUserData(r *http.Request) (User, error) {
	accessToken, err := GetAccessTokenFromCookie(r)
	if err != nil {
		return User{}, err
	}
	return GetUserDataFromToken(accessToken)
}

func GetUserDataFromToken(accessToken string) (User, error) {
	bearer := "Bearer " + accessToken
	getUrl := "https://www.googleapis.com/oauth2/v3/userinfo"
	request, _ := http.NewRequest("GET", getUrl, nil)
	request.Header.Add("Authorization", bearer)

	client := &http.Client{}
	res, _ := client.Do(request)
	defer res.Body.Close()
	var answer User
	err := json.NewDecoder(res.Body).Decode(&answer)
	if err != nil {
		return User{}, err
	}
	if answer.Sub != "" {
		return answer, nil
	} else {
		return answer, errors.New("WTF HUH")
	}
}

func GetLoginUrl() string {
	loginUrl, _ := url.Parse(env.GetCurrUrl() + "/oauth2/auth")
	q := loginUrl.Query()
	q.Set("redirect", env.GetCurrUrl()+env.GetSubpath())
	loginUrl.RawQuery = q.Encode()

	return loginUrl.String()
}
