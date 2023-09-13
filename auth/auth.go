package auth

import (
	"encoding/json"
	"errors"
	"github.com/logotipiwe/dc_go_utils/src/config"
	"net/http"
)

type DcUser struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func FetchUserData(r *http.Request) (DcUser, error) {
	idpHost := config.GetConfig("IDP_HOST")
	idpSubpath := config.GetConfig("IDP_SUBPATH")
	getUrl := idpHost + idpSubpath + "/getUser"
	println(getUrl)
	request, _ := http.NewRequest("GET", getUrl, nil)
	authCookie, err := r.Cookie("access_token")
	if err == nil {
		request.AddCookie(authCookie)
	}

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return DcUser{}, err
	}
	defer res.Body.Close()
	var answer DcUser
	err = json.NewDecoder(res.Body).Decode(&answer)
	if err != nil {
		return DcUser{}, err
	}
	return answer, nil
}

func AuthAsMachine(r *http.Request) error {
	mToken := config.GetConfig("M_TOKEN")
	providedToken := r.URL.Query().Get("mToken")
	if mToken != providedToken {
		return errors.New("not a machine")
	}
	return nil
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
