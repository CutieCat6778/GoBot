package api

import (
	"bytes"
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Map struct {
	HttpClient *http.Client
}

var (
	MapURL     string = "https://maps.googleapis.com/maps/api/staticmap?center=%v,%v8&zoom=%v&size=600x600&maptype=%v&key=" + class.GGAPIKey
	AddressURL string = "http://dev.virtualearth.net/REST/v1/Locations/%q?maxResults=1&key=" + class.BINGKey
)

func NewMap() Map {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return Map{
		HttpClient: &client,
	}
}

func (handler Map) Get(lat float64, long float64, zoom int64, maptype string) *bytes.Reader {
	url := fmt.Sprintf(MapURL, lat, long, zoom, maptype)
	resp, err := handler.HttpClient.Get(url)
	if err != nil {
		utils.HandleServerError(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.HandleServerError(err)
	}

	defer resp.Body.Close()

	return bytes.NewReader(body)
}

func (handler Map) GetAddress(query string) BingRes {
	url := fmt.Sprintf(AddressURL, query)
	resp, err := handler.HttpClient.Get(url)
	if err != nil {
		utils.HandleServerError(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.HandleServerError(err)
	}

	defer resp.Body.Close()

	res := BingRes{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		utils.HandleServerError(err)
	}

	return res
}

func (handler Map) GetMapImage(query string, zoom int64, maptype string) *bytes.Reader {

	var point = []float64{0.0, 0.0}

	res := handler.GetAddress(query)
	if len(res.ResourceSets) > 0 {
		if len(res.ResourceSets[0].Resources) > 0 {
			point = res.ResourceSets[0].Resources[0].Point.Coordinates
		}
	}

	if point[0] == 0.0 || point[0] == 0.0 {
		return nil
	}
	return handler.Get(point[0], point[1], zoom, maptype)
}
