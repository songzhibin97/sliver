package cache

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/SliverHorn/sliver/global"
	"github.com/SliverHorn/sliver/integration/redis/operation"
	"github.com/SliverHorn/sliver/integration/redis/policy"
	"time"
)

const (
	SerializerGob  = "gob"
	SerializerJson = "json"
)

type DbGetter func() interface{}

type Simple struct {
	// 操作类String
	String *operation.String
	// 过期时间
	Expire time.Duration
	// redis没有数据就去gorm去
	DbGetter DbGetter
	// 序列化方式
	Serializer string

	// 策略
	Policy policy.CachePolicy
}

func NewSimple(string *operation.String, expire time.Duration, serializer string, policy policy.CachePolicy) *Simple {
	return &Simple{String: string, Expire: expire, Serializer: serializer, Policy: policy}
}


func (s *Simple) Set(key string, value interface{}) {
	s.String.Set(key, value).UnWarp()
}

func (s *Simple) Get(key string) (result interface{}) {
	if s.Policy != nil { // 检查策略
		s.Policy.Before(key)
	}
	if s.Serializer == SerializerJson { // json 序列化
		f := func() string {
			getter := s.DbGetter()
			b, err := json.Marshal(getter)
			if err != nil {
				global.Zap.Error(global.I18n.TranslateFormat(`{#JsonMarshalFail} %v`, err))
				return ""
			}
			return string(b)
		}
		result = s.String.Get(key).UnWarpElse(f)
	} else if s.Serializer == SerializerGob { // Gob 序列化
		f := func() string {
			getter := s.DbGetter()
			buffer := new(bytes.Buffer)       // 创建一个buffer区
			if err := gob.NewEncoder(buffer).Encode(getter); err != nil { // 创建新的需要转化二进制区域对象
				global.Zap.Error(global.I18n.TranslateFormat(`{#GobNewEncoderFail} %v`, err))
				return ""
			}
			return buffer.String()
		}
		result = s.String.Get(key).UnWarpElse(f)
	}
	if result.(string) == "" && s.Policy == nil {
		s.Policy.IsNil(key, "")
		return result
	}
	s.Set(key, result)
	return result
}

func (s *Simple) GetModel(key string, model interface{}) {
	result := s.Get(key)
	if result == nil {
		return
	}
	if s.Serializer == SerializerJson {
		if err := json.Unmarshal([]byte(result.(string)), model); err != nil {
			global.Zap.Error(global.I18n.TranslateFormat(`{#JsonUnmarshalFail} %v`, err))
			return
		}
		return
	} else if s.Serializer == SerializerGob {
		buffer := new(bytes.Buffer)
		buffer.WriteString(result.(string))
		if err := gob.NewDecoder(buffer).Decode(model); err != nil {
			global.Zap.Error(global.I18n.TranslateFormat(`{#GobNewDecoderFail} %v`, err))
			return
		}
	}

}
