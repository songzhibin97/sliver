package request

import "github.com/SliverHorn/sliver/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}