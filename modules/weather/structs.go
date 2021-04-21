package weather

type WttrInHour struct {
	Dewpointc        string `json:"DewPointC"`
	Dewpointf        string `json:"DewPointF"`
	Feelslikec       string `json:"FeelsLikeC"`
	Feelslikef       string `json:"FeelsLikeF"`
	Heatindexc       string `json:"HeatIndexC"`
	Heatindexf       string `json:"HeatIndexF"`
	Windchillc       string `json:"WindChillC"`
	Windchillf       string `json:"WindChillF"`
	Windgustkmph     string `json:"WindGustKmph"`
	Windgustmiles    string `json:"WindGustMiles"`
	Chanceoffog      string `json:"chanceoffog"`
	Chanceoffrost    string `json:"chanceoffrost"`
	Chanceofhightemp string `json:"chanceofhightemp"`
	Chanceofovercast string `json:"chanceofovercast"`
	Chanceofrain     string `json:"chanceofrain"`
	Chanceofremdry   string `json:"chanceofremdry"`
	Chanceofsnow     string `json:"chanceofsnow"`
	Chanceofsunshine string `json:"chanceofsunshine"`
	Chanceofthunder  string `json:"chanceofthunder"`
	Chanceofwindy    string `json:"chanceofwindy"`
	Cloudcover       string `json:"cloudcover"`
	Humidity         string `json:"humidity"`
	Precipinches     string `json:"precipInches"`
	Precipmm         string `json:"precipMM"`
	Pressure         string `json:"pressure"`
	Pressureinches   string `json:"pressureInches"`
	Tempc            string `json:"tempC"`
	Tempf            string `json:"tempF"`
	Time             string `json:"time"`
	Uvindex          string `json:"uvIndex"`
	Visibility       string `json:"visibility"`
	Visibilitymiles  string `json:"visibilityMiles"`
	Weathercode      string `json:"weatherCode"`
	Weatherdesc      []struct {
		Value string `json:"value"`
	} `json:"weatherDesc"`
	Weathericonurl []struct {
		Value string `json:"value"`
	} `json:"weatherIconUrl"`
	Winddir16Point string `json:"winddir16Point"`
	Winddirdegree  string `json:"winddirDegree"`
	Windspeedkmph  string `json:"windspeedKmph"`
	Windspeedmiles string `json:"windspeedMiles"`
}

type WttrIn struct {
	CurrentCondition []struct {
		Feelslikec       string `json:"FeelsLikeC"`
		Feelslikef       string `json:"FeelsLikeF"`
		Cloudcover       string `json:"cloudcover"`
		Humidity         string `json:"humidity"`
		Localobsdatetime string `json:"localObsDateTime"`
		ObservationTime  string `json:"observation_time"`
		Precipinches     string `json:"precipInches"`
		Precipmm         string `json:"precipMM"`
		Pressure         string `json:"pressure"`
		Pressureinches   string `json:"pressureInches"`
		TempC            string `json:"temp_C"`
		TempF            string `json:"temp_F"`
		Uvindex          string `json:"uvIndex"`
		Visibility       string `json:"visibility"`
		Visibilitymiles  string `json:"visibilityMiles"`
		Weathercode      string `json:"weatherCode"`
		Weatherdesc      []struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
		Weathericonurl []struct {
			Value string `json:"value"`
		} `json:"weatherIconUrl"`
		Winddir16Point string `json:"winddir16Point"`
		Winddirdegree  string `json:"winddirDegree"`
		Windspeedkmph  string `json:"windspeedKmph"`
		Windspeedmiles string `json:"windspeedMiles"`
	} `json:"current_condition"`
	NearestArea []struct {
		Areaname []struct {
			Value string `json:"value"`
		} `json:"areaName"`
		Country []struct {
			Value string `json:"value"`
		} `json:"country"`
		Latitude   string `json:"latitude"`
		Longitude  string `json:"longitude"`
		Population string `json:"population"`
		Region     []struct {
			Value string `json:"value"`
		} `json:"region"`
		Weatherurl []struct {
			Value string `json:"value"`
		} `json:"weatherUrl"`
	} `json:"nearest_area"`
	Request []struct {
		Query string `json:"query"`
		Type  string `json:"type"`
	} `json:"request"`
	Weather []struct {
		Astronomy []struct {
			MoonIllumination string `json:"moon_illumination"`
			MoonPhase        string `json:"moon_phase"`
			Moonrise         string `json:"moonrise"`
			Moonset          string `json:"moonset"`
			Sunrise          string `json:"sunrise"`
			Sunset           string `json:"sunset"`
		} `json:"astronomy"`
		Avgtempc    string        `json:"avgtempC"`
		Avgtempf    string        `json:"avgtempF"`
		Date        string        `json:"date"`
		Hourly      []*WttrInHour `json:"hourly"`
		Maxtempc    string        `json:"maxtempC"`
		Maxtempf    string        `json:"maxtempF"`
		Mintempc    string        `json:"mintempC"`
		Mintempf    string        `json:"mintempF"`
		Sunhour     string        `json:"sunHour"`
		TotalsnowCm string        `json:"totalSnow_cm"`
		Uvindex     string        `json:"uvIndex"`
	} `json:"weather"`
}
