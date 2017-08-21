package main

import (
  "testing")

func TestAbrirArchivo(t_test *testing.T){
  s_texto := abrirArchivo()

  if (s_texto == "") {
		t_test.Fatalf("No fue posible abrir el archivo")
	}
}
