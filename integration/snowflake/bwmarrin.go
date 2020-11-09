package snowflake

import (
	"fmt"
	"github.com/SliverHorn/sliver/global"
	"github.com/bwmarrin/snowflake"
	"time"
)
var node *snowflake.Node

type Bwmarrin struct {}

func (b *Bwmarrin) Init() error {
	t, err := time.Parse("2006-01-02", global.Config.Snowflake.StartTime)
	if err != nil {
		return fmt.Errorf("func time.Parse() failed!, err:%v", err.Error())
	}
	snowflake.Epoch = t.UnixNano() / 1000000
	if node, err = snowflake.NewNode(int64(global.Config.Snowflake.MachineId)); err != nil {
		return fmt.Errorf("func snowflake.NewNode() failed!, err:%v", err.Error())
	}
	return nil
}

func (b *Bwmarrin) GenId() (uint64, error) {
	return uint64(node.Generate().Int64()), nil
}
