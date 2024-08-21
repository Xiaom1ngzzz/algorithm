package trie

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	word := "abc"

	for i := 0; i < len(word); i++ {
		path := byte(word[i]) - 'a'
		xType := reflect.TypeOf(path)
		fmt.Println(path, xType)
	}
}
