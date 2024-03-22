package fred

type Results struct {
	Start        string        `json:"realtime_start" xml:"realtime_start"`
	End          string        `json:"realtime_end" xml:"realtime_end"`
	ObsStart     string        `json:"observation_start" xml:"observation_start"`
	ObsEnd       string        `json:"observation_end" xml:"observation_end"`
	Units        string        `json:"units" xml:"units"`
	OutputType   int           `json:"output_type" xml:"output_type"`
	FileType     string        `json:"file_type" xml:"file_type"`
	OrderBy      string        `json:"order_by" xml:"order_by"`
	SortOrder    string        `json:"sort_order" xml:"sort_order"`
	Count        int           `json:"count" xml:"count"`
	Offset       int           `json:"offset" xml:"offset"`
	Limit        int           `json:"limit" xml:"limit"`
	Categories   []Category    `json:"categories" xml:"category"`
	Releases     []Release     `json:"releases" xml:"release"`
	Seriess      []Series      `json:"seriess" xml:"series"`
	Observations []Observation `json:"observations" xml:"observation"`
	VintageDates []string      `json:"vintage_dates" xml:"vintage_date"`
	Tags         []Tag         `json:"tags" xml:"tag"`
	Sources      []Source      `json:"sources" xml:"source"`
	ReleaseDates []ReleaseDate `json:"release_dates" xml:"release_date"`
}
