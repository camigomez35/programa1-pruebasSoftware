package testNilton

import (
    "fmt"
    "io/ioutil"
)

func abrirArchivo(s_nombreArchivo string) string{
  dat, err := ioutil.ReadFile(s_nombreArchivo)
  if err != nil {
      return ""
  }
  fmt.Print(string(dat))
  return string(dat)
}
