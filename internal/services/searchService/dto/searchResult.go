package dto

import "fmt"

type SearchResult struct {
	DocID         int `json:"doc_id"`
	TermFrequency int `json:"term_frequency"`
}

func (s SearchResult) String() string {
	return fmt.Sprintf("{doc ID : %d , term frequency : %d }", s.DocID, s.TermFrequency)
}

type SearchResults []SearchResult

func (s *SearchResults) Find(docId int) int {
	for i, d := range *s {
		if d.DocID == docId {
			return i
		}
	}
	return -1
}
