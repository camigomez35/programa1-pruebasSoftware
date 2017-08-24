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

func conteoDeLineas() (int,[]int) {
	numeroLineas:=0
	minusculas:="abcdefghijklmnñopqrstuvwxyz"
	var errores []int
	logico, array:=readFile()
	arrayLineas:=encontrarMetodos()
	if logico {
		for i:=0;i<len(arrayLineas) ;i++  {
			indice:=strings.Index(array[arrayLineas[i]],"func")
			substring := array[arrayLineas[i]][strings.IndexAny(array[arrayLineas[i]], " ")+1:len(array[arrayLineas[i]])]
			primeraLetra := substring[0]
			if indice!=0 || strings.IndexAny(minusculas,string(primeraLetra))==-1 {
				errores=append(errores,arrayLineas[i])
			}else {
				numeroLineas++
				//substring[0].IsLower()
				//numeroLineas++

				//numeroLineas:=numeroLineas+contar()
			}

		}
	}
	return numeroLineas, errores
}

