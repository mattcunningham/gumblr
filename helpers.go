package tumblr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// This method GET requests a URL and unmarshals it based on a specified blank struct
// url - The GET URL
// responseObject - A pointer to the blank struct type
func (api Tumblr) info(url string, responseObject interface{}) {
	response := api.get(url)
	if response.Meta.Status != 200 {
		log.Fatalln(fmt.Sprintf("http get error: response status %d with %s",
			response.Meta.Status, response.Meta.Msg))
	}

	err := json.Unmarshal(response.Response, &responseObject)
	if err != nil {
		log.Fatalln(err)
	}
}

// This method GET requests only returning the []byte found
// url - The GET URL
func (api Tumblr) rawGet(url string) []byte {
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	api.oauthService.Sign(request, &api.config)
	client := new(http.Client)
	clientResponse, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}
	defer clientResponse.Body.Close()

	body, err := ioutil.ReadAll(clientResponse.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

// This method GET requests a URL
// url - The GET URL
func (api Tumblr) get(url string) Response {
	body := api.rawGet(url)

	var response Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}
	return response
}

// This method POSTs to a URL
// url - The URL to post to
// params - A string of the encoded parameters
func (api Tumblr) post(url string, params string) Response {
	request, err := http.NewRequest("POST", url, strings.NewReader(params))

	if err != nil {
		log.Fatalln(err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	api.oauthService.Sign(request, &api.config)
	client := new(http.Client)
	clientResponse, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}
	defer clientResponse.Body.Close()

	body, err := ioutil.ReadAll(clientResponse.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}
	return response
}
