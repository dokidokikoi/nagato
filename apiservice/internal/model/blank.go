package model

type Blank struct {
	ID      uint     `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	Matters []Matter `json:"matters"`
}
