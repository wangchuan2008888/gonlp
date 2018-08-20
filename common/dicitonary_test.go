package common

import (
	"testing"
	"fmt"
)

func TestLoadMap(t *testing.T) {
	conf := GetConf()
	fmt.Println(conf)
	dict_path := conf.SegmentDictionaryPath
	fmt.Println(dict_path)
	dictionary := LoadHashDictionary(dict_path)
	fmt.Println(dictionary.GetWord("常驻"))
}
