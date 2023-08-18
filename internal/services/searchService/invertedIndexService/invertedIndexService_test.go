package invertedIndexService

//
//import (
//	"reflect"
//	"testing"
//	"testproject/internal/services/searchService/dto"
//)
//
//type addDataInput struct {
//	text  string
//	docID int
//}
//type QueryDataInput struct {
//	text        []string
//	searchTerms []string
//}
//type QueryDataOutput struct {
//	result dto.SearchResults
//	found  bool
//}
//type fakeDb struct{}
//
//func newFakeDb() *fakeDb {
//	return &fakeDb{}
//}
//func (f *fakeDb) AddData(text string, docId int) {
//	return
//}
//
//func TestInvertedIndex_AddData(t *testing.T) {
//	tests := map[string]struct {
//		input    []addDataInput
//		expected map[string]dto.SearchResults
//	}{
//		`empty`: {
//			input: []addDataInput{
//				{text: "", docID: 1},
//			},
//			expected: map[string]dto.SearchResults{},
//		},
//		`simpleText`: {
//			input: []addDataInput{
//				{text: "alex Bob draws birds bob", docID: 1},
//				{text: "bob goes", docID: 2},
//			},
//			expected: map[string]dto.SearchResults{
//				"alex": dto.SearchResults{
//					{DocID: 1, TermFrequency: 1},
//				},
//				"bob": dto.SearchResults{
//					{DocID: 1, TermFrequency: 2},
//					{DocID: 2, TermFrequency: 1},
//				},
//				"draws": dto.SearchResults{
//					{DocID: 1, TermFrequency: 1},
//				},
//				"birds": dto.SearchResults{
//					{DocID: 1, TermFrequency: 1},
//				},
//				"goes": dto.SearchResults{
//					{DocID: 2, TermFrequency: 1},
//				},
//			},
//		},
//	}
//	for name, tt := range tests {
//		t.Run(name, func(t *testing.T) {
//			invertedIndex := NewInvertedIndex(newFakeDb(), 100)
//			for _, v := range tt.input {
//				invertedIndex.AddData(v.text, v.docID)
//			}
//			if !reflect.DeepEqual(invertedIndex.data, tt.expected) {
//				t.Errorf("expected %s ,\n               got %s", tt.expected, invertedIndex.data)
//			}
//		})
//	}
//}
//
//func TestInvertedIndex_Query(t *testing.T) {
//	tests := map[string]struct {
//		input    QueryDataInput
//		expected []QueryDataOutput
//	}{
//		`test1`: {
//			input: QueryDataInput{
//				text: []string{
//					"hasan went To Bob with Alex",
//					"Bird Goes To Ammar BadFaz With aLex the Alex",
//				},
//				searchTerms: []string{
//					"alex", "hasan", "david",
//				},
//			},
//			expected: []QueryDataOutput{
//				QueryDataOutput{
//					result: dto.SearchResults{
//						dto.SearchResult{
//							DocID: 0, TermFrequency: 1,
//						},
//						dto.SearchResult{
//							DocID: 1, TermFrequency: 2,
//						},
//					},
//					found: true,
//				},
//				QueryDataOutput{
//					result: dto.SearchResults{
//						dto.SearchResult{
//							DocID: 0, TermFrequency: 1,
//						},
//					},
//					found: true,
//				},
//				QueryDataOutput{
//					result: nil,
//					found:  false,
//				},
//			},
//		},
//	}
//	for name, tt := range tests {
//		t.Run(name, func(t *testing.T) {
//			invertedIndex := NewInvertedIndex(newFakeDb(), 100)
//			for i, v := range tt.input.text {
//				invertedIndex.AddData(v, i)
//			}
//			for i, v := range tt.input.searchTerms {
//				res, ok := invertedIndex.Query(v)
//				got := QueryDataOutput{
//					found:  ok,
//					result: res,
//				}
//				if !reflect.DeepEqual(got, tt.expected[i]) {
//					t.Errorf("expected %v ,\n               got %v", tt.expected[i], got)
//				}
//			}
//		})
//	}
//}
//
//func TestNewInvertedIndex(t *testing.T) {
//
//}
