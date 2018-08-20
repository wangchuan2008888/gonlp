package common

import (
	"testing"
	"fmt"
)

func TestLoadMap(t *testing.T){
	dictionary := LoadHashDictionary("/home/andrew/go/src/learnnlp/data/dictionary/CoreNatureDictionary.txt")
	fmt.Println(dictionary.GetWord("北京"))
}
