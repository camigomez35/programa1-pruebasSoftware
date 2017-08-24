package programaDos

import (
	"strings"
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
	array:=abrirArchivo2("archivo.txt")
	for i:=0;i<len(array) ;i++  {
		if strings.Index(array[i],"func")!=-1{
			lineas=append(lineas, i)
		}
	}
	return lineas
}

func conteoDeLineas() int {
	numeroLineas:=0
	return numeroLineas

}