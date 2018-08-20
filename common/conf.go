package common

import "gopkg.in/ini.v1"

type Conf struct {
	SegmentDictionaryPath string `ini:"segment-dictionary-path"`
}

func GetConf() *Conf {
	var config Conf
	var conf_file = "../configure.conf"
	conf, err := ini.Load(conf_file)
	if err != nil {
		panic(err)
	}
	conf.BlockMode = false
	conf.MapTo(&config)
	return &config
}
