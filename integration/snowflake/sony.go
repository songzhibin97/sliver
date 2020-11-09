package snowflake

import (
	"fmt"
	"github.com/SliverHorn/sliver/global"
	"time"

	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

type Sony struct{}

func (s *Sony) Init() error {
	t, err := time.Parse("2006-01-02", global.Config.Snowflake.StartTime)
	if err != nil {
		return fmt.Errorf("func time.Parse() failed!, err:%v", err.Error())
	}
	settings := sonyflake.Settings{StartTime: t, MachineID: func() (uint16, error) { return global.Config.Snowflake.MachineId, nil }}
	sf = sonyflake.NewSonyflake(settings)
	return err
}

func (s *Sony) GenId() (uint64, error) {
	if sf == nil {
		return 0, fmt.Errorf("var sonyflake.Sonyflake is Null, Not inited!, sf:%v", sf)
	}
	return sf.NextID()
}
