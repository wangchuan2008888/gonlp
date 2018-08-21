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
	fmt.Println(len(dictionary.word_map))
}

func TestSimpleCut(t *testing.T) {
	conf := GetConf()
	fmt.Println(conf)
	dict_path := conf.SegmentDictionaryPath
	fmt.Println(dict_path)
	dictionary := LoadHashDictionary(dict_path)

	str := "由于需要指定配置文件，故需要在命令中运."
	str_arr := []rune(str)
	rslt := []string{}
	_len := len(str_arr)
	i := 0
	j := 1
	for ; i < _len-1 && j < _len; {
		current_word := string(str_arr[i:j])
		test_word := string(str_arr[i : j+1])
		has := dictionary.GetWord(test_word)
		if has == nil {
			rslt = append(rslt, current_word)
			i = j
			j = i + 1
		} else {
			j++
		}
	}
	rslt = append(rslt, string(str_arr[i:j]))
	fmt.Println(rslt)

}
