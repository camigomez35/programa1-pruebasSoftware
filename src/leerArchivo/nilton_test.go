package programaDos

import (
	"testing"
)

func TestEmpty(t *testing.T ){
	expected := "dsfa"
	actual := validateCode()
	if actual!=expected {
		t.Errorf("valor no esperado")
	}
}

func TestRead(t *testing.T){
	valor, array:=readFile()
	if valor!=true && len(array)<=0 {
		t.Errorf("El archivo no fue correctamente leido o esta vacio")
	}
}