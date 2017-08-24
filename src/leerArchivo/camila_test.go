package programaDos

import (
  "testing"
  "strings")

func TestVerificarPackage(t_test *testing.T){
  s_texto := abrirArchivo("codigo.txt")
  i_existe := strings.Index(s_texto, "package")
  if (i_existe == -1) {
		t_test.Fatalf("No se encontró el package")
	}
}
