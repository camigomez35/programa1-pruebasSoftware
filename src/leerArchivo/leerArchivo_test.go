package main

import (
  "testing")

func TestAbrirArchivo(t_test *testing.T){
  s_texto := abrirArchivo("archivo.txt")

  if (s_texto == "") {
		t_test.Fatalf("No fue posible abrir el archivo")
	}
}

func TestArchivoNoEncontrado(t_test *testing.T){
  s_texto := abrirArchivo("archivoNoEncontrado.txt")
  if (s_texto != ""){
      t_test.Fatalf("El archivo no estiste")
  }
}
