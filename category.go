package fred

// Category is a single instance of a FRED category.
type Category struct {
	ID       int    `json:"id" xml:"id,attr"`
	Name     string `json:"name" xml:"name,attr"`
	ParentID int    `json:"parent_id" xml:"parent_id,attr"`
}

func (f *FredClient) GetCategory(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetCategory").Msg("getting category")
	return f.Request("fred/category", params)
}

func (f *FredClient) GetCategoryChildren(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetCategoryChildren").Msg("getting category children")
	return f.Request("fred/category/children", params)
}

func (f *FredClient) GetRelatedCategory(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetRelatedCategory").Msg("getting related category")
	return f.Request("fred/category/related", params)
}

func (f *FredClient) GetCategorySeries(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetCategorySeries").Msg("getting category series")
	return f.Request("fred/category/series", params)
}

func (f *FredClient) GetCategoryTags(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetCategoryTags").Msg("getting category tags")
	return f.Request("fred/category/tags", params)
}

func (f *FredClient) GetCategoryRelatedTags(params map[string]interface{}) (results *Results, err error) {
	f.Logger.Debug().Str("func", "GetCategoryRelatedTags").Msg("getting category related tags")
	return f.Request("fred/category/related_tags", params)
}
