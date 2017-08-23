package testNilton

import (
	"testing"
	"fmt"
)

func TestEmpty(t *testing.T ){
	expected := "dsfa"
	actual := validateCode()
	fmt.Println(actual)
	fmt.Println(expected)
}

func TestRead(t *testing.T){
	res:=readFile()
	if res!=true {
		t.Errorf("Test failed, expected: true, got: false")
	}
}