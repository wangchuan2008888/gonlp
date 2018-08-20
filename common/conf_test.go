package common

import (
	"testing"
	"fmt"
)

func TestGetConf(t *testing.T) {
	conf := GetConf()
	fmt.Println(conf.SegmentDictionaryPath)
}
