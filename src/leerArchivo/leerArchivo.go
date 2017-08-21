package main

import (
    "fmt"
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main(){
}

func abrirArchivo() string{
  dat, err := ioutil.ReadFile("archivo.txt")
  check(err)
  fmt.Print(string(dat))
  return string(dat)
}
