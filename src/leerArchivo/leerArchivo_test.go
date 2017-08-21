package main

import (
  "testing")

func TestAbrirArchivo(t_test *testing.T){
  s_texto := abrirArchivo()
  s_textoEsperado := "Texto de prueba"
  if (s_texto != s_textoEsperado) {
		t_test.Fatalf("Esperado %s pero se obtuvo %s", s_textoEsperado , s_texto)
	}
}
