package programaDos

import (
)

func validateCode() string  {
	return "dsfa"
}

func readFile() (bool,[]string){
	array:=abrirArchivo2("archivo.txt")
	if(len(array)>0){
		return true,array
	}else {
		return false,array
	}
}

func encontrarMetodos() []int {
	var lineas []int
	return lineas
}