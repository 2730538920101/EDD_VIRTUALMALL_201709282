//En este paquete definimir la estructura Tienda que tendra los datos de las tiendas
package Tiendas

//Definir la structura Tienda con sus atributos y señalando que pueden venir en formato json
type Tienda struct{
	//Definir un atributo ID para obtener el codigo ascii del nombre y poder hacer el ordenamiento
	Id int 
	Nombre string `json:"Nombre"`
	Descripcion string `json:"Descripcion"`
	Contacto string `json:"Contacto"`
	Calificacion int `json:"Calificacion"`
	Logo string `json: "Logo"`
}
//Definir una estructura para las busquedas por tienda especifica
type Busc_Esp struct{
	Departamento string `json: "Departamento"`
	Nombre string `json: "Nombre"`
	Calificacion int `json: "Calificacion"`
}
//Definir una estructura para las eliminaciones por tienda especifica
type Eliminar_Esp struct{
	Nombre string `json: "Nombre"`
	Categoria string `json: "Categoria"`
	Calificacion int `json: "Calificacion"`
}
//Definir una funcion que genere el Id de las tiendas basados en el codigo ascii para odenarlos al ser insertados en la lista
func (t *Tienda) GenerarId(Nombre string) int{
	str := Nombre
	runes := []rune(str)
	var arr []int
	for i :=0; i < len(runes); i++{
		arr = append(arr,int(runes[i]))
	}
	var resultado = 0
	for _, valor := range arr{
		
		resultado = resultado + valor
	}
	return resultado
}
//Definir la estructura de los departamentos que contiene el nombre del departamento y un arrego de tiendas
type Departamento struct{
	Nombre string `json:"Nombre"`
	Tiendas [] Tienda `json:"Tiendas"`
}
//Definir la estructura Datos que contienen el nombre del indice y un arreglo de departamentos
type Datos struct{
	Indice string `json:"Indice"`
	Departamentos [] Departamento `json:"Departamentos"`
}

type Inicio struct{
	Data []Datos `json:"Datos"`
}

