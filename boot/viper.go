package boot

import (
	"flag"
	"fmt"
	"github.com/SliverHorn/sliver/global"
	"github.com/SliverHorn/sliver/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func InitViper() {
	var p string
	flag.StringVar(&p, "c", "", global.I18n.TranslateFormatLang("{#ChooseConfigFile}", "zh-CN"))
	flag.Parse()
	if p == "" { // 优先级: 命令行 > 环境变量 > 默认值
		if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
			p = utils.ConfigFile
			fmt.Println(global.I18n.TranslateFormatLang("{#DefaultConfigFile} %s", "zh-CN", p))
		} else {
			p = utils.ConfigEnv
			fmt.Println(global.I18n.TranslateFormatLang("{#ConfigEnv} %v", "zh-CN", p))
		}
	} else {
		fmt.Println(global.I18n.TranslateFormatLang("{#CommandConfigFile} %v", "zh-CN", p))
	}
	v := viper.New()
	v.SetConfigFile(p)
	if err := v.ReadInConfig(); err != nil {
		panic(global.I18n.TranslateFormatLang(`{#ReadConfigFileFatal} %v`, "zh-CN", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(global.I18n.TranslateFormat(`{#ConfigChanged} %v`, e.Name))
		if err := v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}
	global.I18n.SetLanguage(global.Config.System.Language)
	global.Viper = v
}
