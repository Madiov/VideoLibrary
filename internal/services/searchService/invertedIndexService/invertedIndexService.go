package invertedIndexService

import (
	"strings"
	"testproject/internal/database"
	"testproject/internal/services/searchService/dto"
)

type InvertedIndexService struct {
	data map[string]dto.SearchResults
	uow  *database.UnitOfWork
}

func NewInvertedIndex(uow *database.UnitOfWork, capacity int) *InvertedIndexService {
	return &InvertedIndexService{
		uow:  uow,
		data: make(map[string]dto.SearchResults, capacity),
	}
}

func (inv *InvertedIndexService) Query(searchTerm string) (dto.SearchResults, bool) {
	res, ok := inv.data[searchTerm]
	return res, ok
}

func (inv *InvertedIndexService) AddData(text string, docId int) {
	if text == "" {
		return
	}
	words := strings.Split(text, ` `)
	for _, w := range words {
		w = strings.ToLower(w)
		res, ok := inv.data[w]
		if !ok {
			inv.data[w] = dto.SearchResults{
				{DocID: docId, TermFrequency: 1},
			}
		} else {
			index := res.Find(docId)
			if index == -1 {
				inv.data[w] = append(inv.data[w], dto.SearchResult{DocID: docId, TermFrequency: 1})
			} else {
				res[index].TermFrequency++
			}

		}
	}
	return
}
