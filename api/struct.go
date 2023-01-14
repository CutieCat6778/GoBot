package api

type BingRes struct {
	AuthenticationResultCode string `json:"authenticationResultCode"`
	BrandLogoURI             string `json:"brandLogoUri"`
	Copyright                string `json:"copyright"`
	ResourceSets             []struct {
		EstimatedTotal int `json:"estimatedTotal"`
		Resources      []struct {
			Type  string    `json:"__type"`
			Bbox  []float64 `json:"bbox"`
			Name  string    `json:"name"`
			Point struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"point"`
			Address struct {
				AddressLine      string `json:"addressLine"`
				AdminDistrict    string `json:"adminDistrict"`
				AdminDistrict2   string `json:"adminDistrict2"`
				CountryRegion    string `json:"countryRegion"`
				FormattedAddress string `json:"formattedAddress"`
				Locality         string `json:"locality"`
				PostalCode       string `json:"postalCode"`
			} `json:"address"`
			Confidence    string `json:"confidence"`
			EntityType    string `json:"entityType"`
			GeocodePoints []struct {
				Type              string    `json:"type"`
				Coordinates       []float64 `json:"coordinates"`
				CalculationMethod string    `json:"calculationMethod"`
				UsageTypes        []string  `json:"usageTypes"`
			} `json:"geocodePoints"`
			MatchCodes []string `json:"matchCodes"`
		} `json:"resources"`
	} `json:"resourceSets"`
	StatusCode        int    `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	TraceID           string `json:"traceId"`
}

type CurrentWeatherStruct struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Rain struct {
		OneH float64 `json:"1h"`
	} `json:"rain"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type RoadRiskStruct []struct {
	Dt      int       `json:"dt"`
	Coord   []float64 `json:"coord"`
	Weather struct {
		Temp                   float64 `json:"temp"`
		WindSpeed              float64 `json:"wind_speed"`
		WindDeg                int     `json:"wind_deg"`
		PrecipitationIntensity float64 `json:"precipitation_intensity"`
		DewPoint               float64 `json:"dew_point"`
	} `json:"weather"`
	Road struct {
		State int     `json:"state"`
		Temp  float64 `json:"temp"`
	} `json:"road"`
	Alerts []struct {
		SenderName string `json:"sender_name"`
		Event      string `json:"event"`
		EventLevel int    `json:"event_level"`
	} `json:"alerts"`
}

type ForecastStruct struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			SeaLevel  int     `json:"sea_level"`
			GrndLevel int     `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
			Gust  float64 `json:"gust"`
		} `json:"wind"`
		Visibility int     `json:"visibility"`
		Pop        float64 `json:"pop"`
		Rain       struct {
			ThreeH float64 `json:"3h"`
		} `json:"rain"`
		Sys struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
	} `json:"list"`
	City struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country    string `json:"country"`
		Population int    `json:"population"`
		Timezone   int    `json:"timezone"`
		Sunrise    int    `json:"sunrise"`
		Sunset     int    `json:"sunset"`
	} `json:"city"`
}

type APOD struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
}

type EarthStruct []struct {
	Identifier          string `json:"identifier"`
	Caption             string `json:"caption"`
	Image               string `json:"image"`
	Version             string `json:"version"`
	CentroidCoordinates struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"centroid_coordinates"`
	DscovrJ2000Position struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"dscovr_j2000_position"`
	LunarJ2000Position struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"lunar_j2000_position"`
	SunJ2000Position struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"sun_j2000_position"`
	AttitudeQuaternions struct {
		Q0 float64 `json:"q0"`
		Q1 float64 `json:"q1"`
		Q2 float64 `json:"q2"`
		Q3 float64 `json:"q3"`
	} `json:"attitude_quaternions"`
	Date   string `json:"date"`
	Coords struct {
		CentroidCoordinates struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"centroid_coordinates"`
		DscovrJ2000Position struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"dscovr_j2000_position"`
		LunarJ2000Position struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"lunar_j2000_position"`
		SunJ2000Position struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"sun_j2000_position"`
		AttitudeQuaternions struct {
			Q0 float64 `json:"q0"`
			Q1 float64 `json:"q1"`
			Q2 float64 `json:"q2"`
			Q3 float64 `json:"q3"`
		} `json:"attitude_quaternions"`
	} `json:"coords"`
}
