package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetAccessToken - returns an access token from Spotify
func GetAccessToken() {
	form := url.Values{}
	form.Add("client_id", ClientCredentials.ClientID)
	form.Add("client_secret", ClientCredentials.ClientSecret)
	form.Add("grant_type", "authorization_code")
	form.Add("code", ClientCredentials.Code)
	form.Add("redirect_uri", ClientCredentials.RedirectURI)
	response, err := http.PostForm("https://accounts.spotify.com/api/token", form)
	if err != nil {
		panic("--- Could not authenticate ---")
	}
	defer response.Body.Close()
	buffer, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(buffer, &authCredentials)
}

func refreshAccessToken() {
	form := url.Values{}
	form.Add("client_id", ClientCredentials.ClientID)
	form.Add("client_secret", ClientCredentials.ClientSecret)
	form.Add("grant_type", "refresh_token")
	form.Add("refresh_token", authCredentials.RefreshToken)
	response, err := http.PostForm("https://accounts.spotify.com/api/token", form)
	if err != nil {
		panic("--- Could not authenticate ---")
	}
	defer response.Body.Close()
	buffer, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(buffer, &authCredentials)
}
