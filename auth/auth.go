package auth

import (
	"encoding/json"
	"net/http"
)

type DcUser struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func FetchUserData(r *http.Request) (DcUser, error) {
	getUrl := "http://idp/getUser"
	println(getUrl)
	request, _ := http.NewRequest("GET", getUrl, nil)
	authCookie, err := r.Cookie("access_token")
	if err == nil {
		request.AddCookie(authCookie)
	}

	client := &http.Client{}
	res, _ := client.Do(request)
	defer res.Body.Close()
	var answer DcUser
	err = json.NewDecoder(res.Body).Decode(&answer)
	if err != nil {
		return DcUser{}, err
	}
	return answer, nil
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
