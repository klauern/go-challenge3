package mosaic

import (
	"encoding/json"
	"image"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const endpoint string = "https://api.flickr.com/services/rest/?"

type Photo struct {
	Id         string `json:"id"`
	Owner      string `json:"owner"`
	Secret     string `json:"secret"`
	Server     string `json:"server"`
	Farm       int    `json:"farm"`
	Title      string `json:"title"`
	IsPublic   int    `json:"ispublic"`
	IsFriendly int    `json:"isfriendly"`
	IsFamily   int    `json:"isfamily"`
}

type Photos struct {
	Page      int     `json:"page"`
	Pages     int     `json:"pages"`
	PerPage   int     `json:"perpage"`
	Total     string  `json:"total"`
	PhotoList []Photo `json:"photo"`
	Stat      string  `json:"stat"`
}

type FPhotos struct {
	Photos Photos `json:"photos"`
}

func (p *Photos) Unmarshal(body io.Reader) {
	content, err := ioutil.ReadAll(body)

	if err != nil {
		panic(err.Error())
	}

	var fphotos FPhotos
	json.Unmarshal(content, &fphotos)

	*p = fphotos.Photos
	//out, err := json.Marshal(p)
	//if err != nil {
	//panic(err.Error())
	//}
	//fmt.Println(string(out))
}

func GetInterestingness() *Photos {
	v := url.Values{}
	v.Set("method", "flickr.interestingness.getList")
	v.Set("api_key", os.Getenv("GOCHALLENGEkey"))
	v.Set("format", "json")
	v.Set("nojsoncallback", "1")

	resp, err := http.Get(endpoint + v.Encode())

	if err != nil {
		panic(err.Error())
	}

	photos := new(Photos)
	photos.Unmarshal(resp.Body)

	return photos
}

func GetPhoto(photo Photo) (img *image.Image, err error) {

	return img, err
}
