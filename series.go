package fred

// Series is a single instance of a FRED series.
type Series struct {
	ID                      string `json:"id" xml:"id,attr"`
	Start                   string `json:"realtime_start" xml:"realtime_start,attr"`
	End                     string `json:"realtime_end" xml:"realtime_end,attr"`
	Title                   string `json:"title" xml:"title,attr"`
	ObsStart                string `json:"observation_start" xml:"observation_start,attr"`
	ObsEnd                  string `json:"observation_end" xml:"observation_end,attr"`
	Frequency               string `json:"annual" xml:"annual,attr"`
	FrequencyShort          string `json:"frequency_short" xml:"frequency_short,attr"`
	Units                   string `json:"units" xml:"units,attr"`
	UnitsShort              string `json:"units_short" xml:"units_short,attr"`
	SeasonalAdjustment      string `json:"seasonal_adjustment" xml:"seasonal_adjustment,attr"`
	SeasonalAdjustmentShort string `json:"seasonal_adjustment_short" xml:"seasonal_adjustment_short,attr"`
	LastUpdated             string `json:"last_updated" xml:"last_updated,attr"`
	Popularity              int    `json:"popularity" xml:"popularity,attr"`
	GroupPopularity         int    `json:"group_popularity" xml:"group_popularity,attr"`
	Notes                   string `json:"notes" xml:"notes,attr"`
}

// Observation is a single instance of a FRED observation.
type Observation struct {
	Start string `json:"realtime_start" xml:"realtime_start,attr"`
	End   string `json:"realtime_end" xml:"realtime_end,attr"`
	Date  string `json:"date" xml:"date,attr"`
	Value string `json:"value" xml:"value,attr"`
}

func (f *FredClient) GetSeries(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetSeries").Msg("getting series")
	return f.Request("fred/series", params)
}

func (f *FredClient) GetSeriesCategories(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetSeriesCategories").Msg("getting series categories")
	return f.Request("fred/series/categories", params)
}

func (f *FredClient) GetSeriesObservations(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetSeriesObservations").Msg("getting series observations")
	return f.Request("fred/series/observations", params)
}

func (f *FredClient) GetSeriesRelease(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetSeriesRelease").Msg("getting series release")
	return f.Request("fred/series/release", params)
}

func (f *FredClient) GetSeriesSearch(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetSeriesSearch").Msg("getting series search")
	return f.Request("fred/series/search", params)
}

func (f *FredClient) GetSeriesSearchTags(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetSeriesSearchTags").Msg("getting series search tags")
	return f.Request("fred/series/search/tags", params)
}

func (f *FredClient) GetSeriesSearchRelatedTags(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetSeriesSearchRelatedTags").Msg("getting series search related tags")
	return f.Request("fred/series/search/related_tags", params)
}

func (f *FredClient) GetSeriesTags(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetSeriesTags").Msg("getting series tags")
	return f.Request("fred/series/tags", params)
}

func (f *FredClient) GetSeriesUpdates(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetSeriesUpdates").Msg("getting series updates")
	return f.Request("fred/series/updates", params)
}

func (f *FredClient) GetSeriesVintageDates(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetSeriesVintageDates").Msg("getting series vintage dates")
	return f.Request("fred/series/vintagedates", params)
}
