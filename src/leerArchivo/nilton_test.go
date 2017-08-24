package programaDos

import (
	"testing"
	"fmt"
)

func TestVacio(t *testing.T ){
	expected := "dsfa"
	actual := validateCode()
	if actual!=expected {
		t.Errorf("valor no esperado")
	}
}

func TestLeer(t *testing.T){
	valor, array:=readFile()
	if valor!=true && len(array)<=0 {
		t.Errorf("El archivo no fue correctamente leido o esta vacio")
	}
}

func TestEncontrarMetodos(t *testing.T){
	lineas:=encontrarMetodos()
	if len(lineas)<=0 {
		t.Errorf("En el archivo no hay metodos")
	}
}