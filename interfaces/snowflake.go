package interfaces

import (
    "github.com/SliverHorn/sliver/global"
	"github.com/SliverHorn/sliver/integration/snowflake"
)

type Snowflake interface {
    // 初始化
    Init() error

    // ⽣生成id
    GenId() (uint64, error)
}

func NewSnowflake() Snowflake {
    switch global.Config.Snowflake.Package {
    case "Sony":
        return &snowflake.Sony{}
    case "Bwmarrin":
        return &snowflake.Bwmarrin{}
    default:
        return &snowflake.Sony{}
    }
}