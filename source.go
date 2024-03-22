package fred

// Source is a single instance of a FRED Source.
type Source struct {
	ID    int    `json:"id" xml:"id,attr"`
	Start string `json:"realtime_start" xml:"realtime_start,attr"`
	End   string `json:"realtime_end" xml:"realtime_end,attr"`
	Name  string `json:"name" xml:"name,attr"`
	Link  string `json:"link" xml:"link,attr"`
}

func (f *FredClient) GetSources(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetSources").Msg("getting sources")
	return f.Request("fred/sources", params)
}

func (f *FredClient) GetSource(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetSource").Msg("getting source")
	return f.Request("fred/source", params)
}

func (f *FredClient) GetSourceReleases(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetSourceReleases").Msg("getting source releases")
	return f.Request("fred/source/releases", params)
}
