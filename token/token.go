package token

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	ClientID     string
	ClientSecret string
	Expiration   time.Time
	Region       string
}

func (t *Token) Get() string {
	if t.Expiration.Unix() < time.Now().Unix() {
		t.Renew()
	}
	return t.AccessToken
}

func (t *Token) Renew() error {

	// Turning our webhook payload into json data
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	urlString := fmt.Sprintf("https://%s.battle.net/oauth/token", t.Region)

	client := http.Client{}
	req, err := http.NewRequest("POST", urlString, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(url.QueryEscape(t.ClientID), url.QueryEscape(t.ClientSecret))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf(string(respBody))
	}

	if err = json.Unmarshal(respBody, t); err != nil {
		return err
	}

	t.Expiration = time.Now().Add(time.Second * 86399)
	return nil

}
