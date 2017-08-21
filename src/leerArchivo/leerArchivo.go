package main

import (
    "fmt"
    "io/ioutil"
)

func main(){
  abrirArchivo("./src/leerArchivo/archivo.txt")
}

func abrirArchivo(s_nombreArchivo string) string{
  dat, err := ioutil.ReadFile(s_nombreArchivo)
  if err != nil {
      return ""
  }
  fmt.Print(string(dat))
  return string(dat)
}
