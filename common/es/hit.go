package es

type BaseModel = any

type SearchResult[T BaseModel] struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Source T `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (res SearchResult[T]) GetSourceList() []*T {
	var list []*T
	for i, _ := range res.Hits.Hits {
		list = append(list, &res.Hits.Hits[i].Source)
	}

	return list
}

type DocResult[T BaseModel] struct {
	Source T `json:"_source"`
}
