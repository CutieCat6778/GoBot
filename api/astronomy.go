package api

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"encoding/json"
	"image"
	"io"
	"net/http"
)

type Astronomy struct {
	HttpClient *http.Client
}

var (
	APODURL string = "https://api.nasa.gov/planetary/apod?api_key=" + class.NasaKey
)

func NewAstronomy() Astronomy {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return Astronomy{
		HttpClient: &client,
	}
}

func (handler Astronomy) APOD() APOD {
	resp, err := handler.HttpClient.Get(APODURL)
	if err != nil {
		utils.HandleServerError(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.HandleServerError(err)
	}

	defer resp.Body.Close()

	res := APOD{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		utils.HandleServerError(err)
	}

	return res
}

func (handler Astronomy) GetImageSize(url string) (int, int) {
	resp, err := handler.HttpClient.Get(url)
	if err != nil {
		utils.HandleServerError(err)
	}
	defer resp.Body.Close()

	m, _, err := image.Decode(resp.Body)
	if err != nil {
		utils.HandleServerError(err)
	}
	g := m.Bounds()

	// Get height and width
	height := g.Dy()
	width := g.Dx()

	return height, width
}
