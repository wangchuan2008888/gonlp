package common

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
	"regexp"
)

type Dictionary interface {
	SaveDictionary(string) (*error)
	AddWord(string) (*error)
	GetWord(string) (*WordTag)
}

type WordTag struct {
	Pos   string
	Count int
}

type HashMapDictionary struct {
	word_map map[string]*WordTag
}

func newHashMapDictionary() *HashMapDictionary {
	m := make(map[string]*WordTag)
	return &HashMapDictionary{m}
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
		if len(eles) < 3 {
			fmt.Printf("line broken %s\n", line)
			fmt.Println(eles)
			continue
		}
		word := eles[0]
		pos := eles[1]
		count, erro := strconv.Atoi(eles[2])
		if erro != nil {
			panic(erro)
		}
		hash_dict.word_map[word] = &WordTag{pos, count}
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
