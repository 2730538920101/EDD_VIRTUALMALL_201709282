//En este paquete definir una Lista Doblemente enlazada para su implementacion
package Estructura
//Importar los paquetes necesarios
import (
	"fmt"
	"../Tiendas"
)

//Definir la structura Nodo que contendra como atributo una estructura de tipo Tienda
type Nodo struct{
	//declarar los punteros de tipo nodo, uno para el anterior y uno para el siguiente
	anterior *Nodo
	siguiente *Nodo
	//declarar la variable tipo Tienda que contendra la structura Tienda que se creara
	NodoTienda *Tiendas.Tienda
}

//Definir la estructura Lista que contendra las estructuras de tipo Nodo
type Lista struct{
	//declarar los punteros de tipo nodo al inicio y al final de la lista
	inicio *Nodo
	ultimo *Nodo
	//declarar una variable para el tamaño de la lista
	tam int
}

//Definir una funcion para crear una nueva lista
func Nueva_Lista() *Lista{

	return &Lista{nil,nil,0}
}

//Definir una funcion para insertar un nuevo nodo
func (l *Lista)Insertar(t *Tiendas.Tienda){
	//Declarar un nuevo nodo
	nuevo := &Nodo{nil,nil,t}
	//Si el dato a ingresar es el primero de la lista
	if l.inicio == nil{
		//El inicio y el final apuntan a null
		l.inicio = nuevo
		l.ultimo = nuevo
		//Si no es el primer elemento de la lista
	}else{
		//el nodo que se crea de ultimo apuntara al siguiente y almacena al nuevo nodo
		l.ultimo.siguiente = nuevo
		//el nuevo nodo apunta el valor en la ultimo posicion de la lista
		nuevo.anterior = l.ultimo
		//El ultimo nodo almacena al nuevo
		l.ultimo = nuevo
	}
	l.tam ++

}
//Definir una funcion para imprimir la lista
func (l *Lista)Imprimir(){
	//Declarar nodo auxiliar para recorrer la lista
	aux := l.inicio
	//Iniciar un ciclo for que funciones mientras el nodo auxiliar sea diferente de null
	for aux !=nil{
		fmt.Println("DATO: ", aux.NodoTienda)
		aux = aux.siguiente
	}
	fmt.Println()
	fmt.Println("EL TAMAÑO DE LA LISTA ES: ", l.tam)
}

//Definir una funcion para buscar un nodo dentro de la lista
func (l *Lista)Buscar(Identificador int) *Nodo{
	//definir un nodo auxiliar para recorrer la lista
	aux := l.inicio
	for aux != nil{
		//Verificar si el nodo en su propiedad de Id es igual al Id ingresado
		if aux.NodoTienda.Id == Identificador{
			//Si es igual devolver el nodo encontrado
			fmt.Println("Se encontro el nodo")
			return aux
		}
		//Si no lo encuentra en la lista, pasar al siguiente nodo
		aux = aux.siguiente
	}
	//Si al terminar de leer la lista no lo encontro, se envia un msj y retorna el nodo aux apuntando a null
	fmt.Println("No se encontro el nodo")
	return aux
}

//Definir un metodo para eliminar de la lista por Id del nodo
func (l *Lista) Eliminar(Identificador int){
	//Declarar un nodo auxiliar que busque el nodo y si existe para poder hacer las operaciones de comparacion
	aux := l.Buscar(Identificador)
	//Si el nodo encontrado esta al inicio de la lista
	if l.inicio == aux{
		l.inicio = aux.siguiente
		aux.siguiente.anterior = nil 
		aux.siguiente = nil 
		fmt.Println("Se ha eliminado un elemento de la lista")
	//Si el nodo encontrado se encuentra al final de la lista
	}else if l.ultimo == aux {
		l.ultimo = aux.anterior
		aux.anterior.siguiente = nil 
		aux.anterior = nil 
		fmt.Println("Se ha eliminado un elemento de la lista")
	//Si el nodo encontrado esta entre los elementos de la lista pero no al principio ni al final
	}else{
		aux.anterior.siguiente = aux.siguiente
		aux.siguiente.anterior = aux.anterior
		aux.anterior = nil 
		aux.siguiente = nil
		fmt.Println("Se ha eliminado un elemento de la lista")
	}
	//Restar un elemento del contador del tamaño de la lista
	l.tam --
}