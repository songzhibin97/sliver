package result

import "github.com/SliverHorn/sliver/global"

type Interface struct {
   Result interface{}
   Error error
}

type empty struct {}

func NewInterface(result interface{}, error error) *Interface {
   return &Interface{Result: result, Error: error}
}

func (i *Interface) UnWarp() interface{} {
   if i.Error != nil {
      global.Zap.Error(global.I18n.TranslateFormat(`{#InterfaceUnWarpFail} %v`, i.Error))
      return empty{}
   }
   return i.Result
}

func (i *Interface) UnWarpDefault(d interface{}) interface{} {
   if i.Error != nil {
      return d
   }
   return i.Result
}