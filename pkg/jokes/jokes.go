package jokes

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
)

const jokesApi = "https://fatherhood.gov/jsonapi/node/dad_jokes"

type Joke struct {
	Opener    string
	Punchline string
}

type jokesResponse struct {
	Data []jokeData
}

type jokeData struct {
	Attributes jokeAttributes
}

type jokeAttributes struct {
	Opener   string `json:"field_joke_opener"`
	Response string `json:"field_joke_response"`
}

func GetRandomJoke() Joke {
	jds := getJokes()
	jd := jds[rand.Intn(len(jds))]

	return Joke{
		Opener:    jd.Attributes.Opener,
		Punchline: jd.Attributes.Response,
	}
}

func getJokes() []jokeData {
	resp, err := http.Get(jokesApi)
	if err != nil {
		panic(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		panic(readErr)
	}

	jr := jokesResponse{}
	jsonErr := json.Unmarshal(body, &jr)
	if jsonErr != nil {
		panic(jsonErr)
	}

	return jr.Data
}
