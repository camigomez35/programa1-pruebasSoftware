package leerArchivo

import (
	"testing"
	"fmt"
)

func TestEmpty(t *testing.T ){
	expected := "dsfa"
	actual := validateCode()
	fmt.Println(actual)
	fmt.Println(expected)
	/*if actual != expected {
		//t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}*/
}