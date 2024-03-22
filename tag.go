package fred

// Tag is a single instance of a FRED tag.
type Tag struct {
	Name        string `json:"name" xml:"name,attr"`
	GroupID     string `json:"group_id" xml:"group_id,attr"`
	Notes       string `json:"notes" xml:"notes,attr"`
	Created     string `json:"created" xml:"created,attr"`
	Popularity  int    `json:"popularity" xml:"popularity,attr"`
	SeriesCount int    `json:"series_count" xml:"series_count,attr"`
}

func (f *FredClient) GetTags(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetTags").Msg("getting tags")
	return f.Request("fred/tags", params)
}

func (f *FredClient) GetRelatedTags(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetRelatedTags").Msg("getting related tags")
	return f.Request("fred/related_tags", params)
}

func (f *FredClient) GetTagSeries(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetTagSeries").Msg("getting tag series")
	return f.Request("fred/tag/series", params)
}
