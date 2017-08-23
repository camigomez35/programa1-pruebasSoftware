package testNilton


func validateCode() string  {
	return "dsfa"
}

func readFile() bool  {
	file:= abrirArchivo("archivo.txt")
	if file != "" {
		return true
	}else {
		return false
	}
}