//En este paquete definimir la estructura Tienda que tendra los datos de las tiendas
package Tiendas
//Importar los paquetes necesarios
import "fmt"
//Definir la structura Tienda con sus atributos
type Tienda struct{
	Id int
	Nombre string
	Descripcion string
	Contacto string
	Calificacion int
}

//Definir una funcion para crear nuevas tiendas
func Nueva_Tienda(Id int, Nombre string, Descripcion string, Contacto string, Calificacion int) *Tienda{
	fmt.Println("Se ha creado una nueva tienda exitosamene")
	return &Tienda{Id, Nombre, Descripcion, Contacto, Calificacion}
}