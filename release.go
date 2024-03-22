package fred

// Release is a single instance of a release of FRED economic data.
type Release struct {
	ID           int    `json:"id" xml:"id,attr"`
	Start        string `json:"realtime_start" xml:"realtime_start,attr"`
	End          string `json:"realtime_end" xml:"realtime_end,attr"`
	Name         string `json:"name" xml:"name,attr"`
	PressRelease bool   `json:"press_release" xml:"press_release,attr"`
	Link         string `json:"link" xml:"link,attr"`
	Notes        string `json:"notes" xml:"notes,attr"`
}

// ReleaseDate is a single instance of a release date of FRED economic data.
type ReleaseDate struct {
	ID   int    `json:"release_id" xml:"release_id,attr"`
	Name string `json:"release_name" xml:"release_name,attr"`
	Date string `json:"date" xml:",chardata"`
}

func (f *FredClient) GetReleases(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetReleases").Msg("getting releases")
	return f.Request("fred/releases", params)
}

func (f *FredClient) GetReleasesDates(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetReleasesDates").Msg("getting releases dates")
	return f.Request("fred/releases/dates", params)
}

func (f *FredClient) GetRelease(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetRelease").Msg("getting release")
	return f.Request("fred/release", params)
}

func (f *FredClient) GetReleaseDates(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetReleaseDates").Msg("getting release dates")
	return f.Request("fred/release/dates", params)
}

func (f *FredClient) GetReleaseSeries(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetReleaseSeries").Msg("getting release series")
	return f.Request("fred/release/series", params)
}

func (f *FredClient) GetReleaseSources(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetReleaseSources").Msg("getting release sources")
	return f.Request("fred/release/sources", params)
}

func (f *FredClient) GetReleaseTags(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetReleaseTags").Msg("getting release tags")
	return f.Request("fred/release/tags", params)
}

func (f *FredClient) GetReleaseRelatedTags(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetReleaseRelatedTags").Msg("getting release related tags")
	return f.Request("fred/release/related_tags", params)
}
