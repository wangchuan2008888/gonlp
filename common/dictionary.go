package common

import (
	"os"
	"bufio"
	"strconv"
	"regexp"
)

type Dictionary interface {
	SaveDictionary(string) (*error)
	AddWord(string) (*error)
	GetWord(string) (*WordTag)
}

type WordTag struct {
	Pos            []string
	Count          []int
	TotalFrequency int
}

func NewWordTag(n int) *WordTag {
	return &WordTag{make([]string, 0, n), make([]int, 0, n), 0}
}

type HashMapDictionary struct {
	word_map map[string]*WordTag
}

func newHashMapDictionary() *HashMapDictionary {
	m := make(map[string]*WordTag)
	return &HashMapDictionary{m}
}
func loadLine(line []string) (rslt *WordTag) {
	_len := len(line)
	rslt = NewWordTag(_len / 2)
	i := 0
	for {
		rslt.Pos = append(rslt.Pos, line[i])
		value, _ := strconv.Atoi(line[i+1])
		rslt.TotalFrequency += value
		rslt.Count = append(rslt.Count, value)
		i += 2
		if i > _len-1 {
			break
		}
	}

	return
}
func LoadHashDictionary(dictionary_path string) *HashMapDictionary {
	dictionary_file, err := os.Open(dictionary_path)
	if err != nil {
		panic(err)
	}
	defer dictionary_file.Close()
	hash_dict := newHashMapDictionary()
	dictionary_scanner := bufio.NewScanner(dictionary_file)
	counter := 0
	splitter := regexp.MustCompile("\\s+")

	for dictionary_scanner.Scan() {
		line := dictionary_scanner.Text()
		eles := splitter.Split(line, -1)

		word := eles[0]
		tag := loadLine(eles[1:])
		hash_dict.word_map[word] = tag
		counter += 1
	}
	println("loaded %d words", counter)
	return hash_dict
}
func (hash_dict *HashMapDictionary) SaveDictionary(s string) (*error) {
	return nil
}
func (hash_dict *HashMapDictionary) AddWord(s string) (*error) {
	return nil
}
func (hash_dict *HashMapDictionary) GetWord(s string) *WordTag {
	_, ok := hash_dict.word_map[s]
	if ok {
		return hash_dict.word_map[s]
	} else {
		return nil
	}
}
