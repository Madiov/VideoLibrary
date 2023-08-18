package contracts

import (
	"testproject/internal/services/searchService/dto"
)

type ISearchService interface {
	Query(searchTerm string) (dto.SearchResults, bool)
	AddData(term string, docID int)
}
