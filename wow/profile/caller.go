package profile

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *CharacterClient) Caller(endpoint string) ([]byte, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", c.BaseUrl+endpoint+"?locale="+c.Locale, nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Battlenet-Namespace": []string{c.Namespace},
		"Authorization":       []string{"Bearer " + c.Token.Get()},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(string(body))
	}
	return body, nil
}
