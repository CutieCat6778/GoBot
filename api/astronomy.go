package api

import (
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"encoding/json"
	"fmt"
	"image"
	"io"
	"net/http"
	"strings"
)

type Astronomy struct {
	HttpClient *http.Client
}

var (
	APODURL   string = "https://api.nasa.gov/planetary/apod?api_key=" + class.NasaKey
	EARTHURL  string = "https://api.nasa.gov/EPIC/api/natural/date/%v?api_key=" + class.NasaKey
	EARTHURL2 string = "https://api.nasa.gov/EPIC/api/natural?api_key=" + class.NasaKey
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

func (handler Astronomy) Earth2() EarthStruct {
	resp, err := handler.HttpClient.Get(EARTHURL2)
	if err != nil {
		utils.HandleServerError(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.HandleServerError(err)
	}

	defer resp.Body.Close()

	res := EarthStruct{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		utils.HandleServerError(err)
	}

	return res
}

func (handler Astronomy) Earth(date string) EarthStruct {
	resp, err := handler.HttpClient.Get(fmt.Sprintf(EARTHURL, date))
	if err != nil {
		utils.HandleServerError(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.HandleServerError(err)
	}

	defer resp.Body.Close()

	res := EarthStruct{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		utils.HandleServerError(err)
	}

	return res
}

func (handler Astronomy) EarthImage(date string, id string) string {
	//2023-01-13 00:22:24
	format1 := strings.Split(date, " ")
	format2 := strings.Split(format1[0], "-")
	dates := strings.Join(format2, "/")

	//https://api.nasa.gov/EPIC/archive/natural/2019/05/30/png/epic_1b_20190530011359.png?api_key=DEMO_KEY
	return fmt.Sprintf("https://api.nasa.gov/EPIC/archive/natural/%v/png/%v.png?api_key=%v", dates, id, class.NasaKey)
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
