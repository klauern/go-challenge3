package mosaic

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const endpoint string = "https://api.flickr.com/services/rest/?"

type Photo struct {
	Id         int
	Owner      string
	Secret     string
	Server     int
	Farm       int
	Title      string
	IsPublic   int
	IsFriendly int
	IsFamily   int
}

type Photos struct {
	Page      int
	Pages     int
	PerPage   int
	Total     int
	PhotoList []Photo `json:photo`
	Stat      string
}

func (p *Photos) Unmarshal(body io.ReadCloser) {
	content, err := ioutil.ReadAll(body)

	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(content, &p)
}

func GetInterestingness() Photos {
	v := url.Values{}
	v.Set("method", "flickr.interestingness.getList")
	v.Set("api_key", os.Getenv("GOCHALLENGEkey"))
	v.Set("format", "json")
	v.Set("nojsoncallback", "1")

	resp, err := http.Get(endpoint + v.Encode())

	if err != nil {
		panic(err.Error())
	}

	var photos Photos
	photos.Unmarshal(resp.Body)

	return photos

}
