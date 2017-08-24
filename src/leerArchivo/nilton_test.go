package programaDos

import (
	"testing"
	"strconv"
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

func TestValidarFuncPrimeraPalabra(t *testing.T)  {
	numeroLineas, errores :=conteoDeLineas()
	if numeroLineas==0 {
		t.Errorf("Algúna funcion no empieza por la pabra reservada func")
	}else if(len(errores)>0){
		for i:=0;i<len(errores);i++  {
			t.Errorf("la función de la línea "+strconv.Itoa(errores[i])+" no esta acorde al estandar")
		}
	}
}

func TestValidarMinusculaNombreMetodo(t *testing.T){
	numeroLineas, errores :=conteoDeLineas()
	if numeroLineas==0 {
		t.Errorf("Algúna funcion no empieza por la pabra reservada func")
	}else if(len(errores)>0){
		for i:=0;i<len(errores);i++  {
			t.Errorf("la función de la línea "+strconv.Itoa(errores[i])+" no esta acorde al estandar")
		}
	}
}