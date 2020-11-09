package result

import "github.com/SliverHorn/sliver/global"

type String struct {
   Result string
   Error error
}

func NewString(result string, error error) *String {
   return &String{Result: result, Error: error}
}

func (s *String) UnWarp() string {
   if s.Error != nil {
      global.Zap.Error(global.I18n.TranslateFormat(`{#StringUnWarpFail} %v`, s.Error))
      return ""
   }
   return s.Result
}

func (s *String) UnWarpElse(f func() string) string {
   if s.Error != nil {
      return f()
   }
   return s.Result
}

func (s *String) UnWarpDefault(d string) string {
   if s.Error != nil {
      return d
   }
   return s.Result
}