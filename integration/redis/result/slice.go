package result

import (
	"github.com/SliverHorn/sliver/global"
	"github.com/SliverHorn/sliver/integration/redis/iterator"
)

type Slice struct {
   Results []interface{}
   Error error
}

func NewSlice(results []interface{}, error error) *Slice {
	return &Slice{Results: results, Error: error}
}

func (s *Slice) Iter(results []interface{}) *iterator.Iterator {
	return iterator.NewIterator(results)
}

func (s *Slice) UnWarp() []interface{} {
	if s.Error != nil {
		global.Zap.Error(global.I18n.TranslateFormat(`{#SliceUnWarpFail} %v`, s.Error))
		return []interface{}{}
	}
	return s.Results
}

func (s *Slice) UnWarpDefault(results []interface{}) []interface{} {
	if s.Error != nil {
		return results
	}
	return s.Results
}
