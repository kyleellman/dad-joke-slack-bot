package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const postMessageApi = "https://slack.com/api/chat.postMessage"

type SimpleSlackClient struct {
	token      string
	channelId  string
	httpClient http.Client
}

type SlackMessageResponse struct {
	Ok    bool
	Error string
	Ts    string
}

func NewSimpleClient(token string, channelId string) SimpleSlackClient {
	return SimpleSlackClient{
		token:      token,
		channelId:  channelId,
		httpClient: http.Client{},
	}
}

func (client SimpleSlackClient) SendMessage(text string, threadTs *string) SlackMessageResponse {
	payload := map[string]string{
		"channel": client.channelId,
		"text":    text,
	}

	if threadTs != nil {
		payload["thread_ts"] = *threadTs
	}

	json_data, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", postMessageApi, bytes.NewBuffer(json_data))
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.token))

	resp, err := client.httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode > 299 {
		panic(resp.Status)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		panic(readErr)
	}

	smr := SlackMessageResponse{}
	jsonErr := json.Unmarshal(body, &smr)
	if jsonErr != nil {
		panic(jsonErr)
	}

	if !smr.Ok {
		panic(smr.Error)
	}

	return smr
}
