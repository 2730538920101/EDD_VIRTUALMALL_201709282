//En este paquete definimir la estructura Tienda que tendra los datos de las tiendas
package Tiendas
//Importar los paquetes necesarios
import "fmt"
//Definir la structura Tienda con sus atributos y señalando que pueden venir en formato json
type Tienda struct{
	//Definir un atributo ID para obtener el codigo ascii del nombre y poder hacer el ordenamiento
	Id int 
	Nombre string `json:"Nombre, omitempty"`
	Descripcion string `json:"Descripcion, omitempty"`
	Contacto string `json:"Contacto, omitempty"`
	Calificacion int `json:"Calificacion, omitempty"`
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
	Nombre string `json:"Nombre, omitempty"`
	Tiendas [] Tienda `json:"Tiendas, omitempty"`
}
//Definir la estructura Datos que contienen el nombre del indice y un arreglo de departamentos
type Datos struct{
	Indice string `json:"Indice, omitempty"`
	Departamentos [] Departamento `json:"Departamentos, omitempty"`
}

type Inicio struct{
	Data []Datos `json:"Datos, omitempty"`
}

//Definir una funcion para crear nuevas tiendas
func Nueva_Tienda(Id int, Nombre string, Descripcion string, Contacto string, Calificacion int) *Tienda{
	fmt.Println("Se ha creado una nueva tienda exitosamene")
	return &Tienda{Id, Nombre, Descripcion, Contacto, Calificacion}
}