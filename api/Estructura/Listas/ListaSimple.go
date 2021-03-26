//Se define una lista simple para el manejo de los pedidos en el carrito de compras
package Listas


import (
	"../Tiendas"
	"fmt"
	"encoding/json"
	"log"
)

//Definir el nodo del producto en en carrito de compras
type NodoProd struct{
	siguiente *NodoProd
	anterior *NodoProd
	producto *Tiendas.Producto
}

//Estructura que nos permite almacenar los nodos
type ListaSimple struct{
	inicio *NodoProd
	ultimo *NodoProd
	tam int
}

//Funcion para crear una nueva lista simple
func Nueva_ListaS() *ListaSimple{
	return &ListaSimple{nil, nil, 0}
}

//Insertar un nodo en la lista
func (l *ListaSimple) Insertar(prod *Tiendas.Producto){
	nuevo := &NodoProd{nil,nil, prod}
	
	if l.inicio == nil{
		l.inicio = nuevo
		l.ultimo = nuevo
		l.tam++
	}else{
		l.ultimo.siguiente = nuevo
		l.ultimo = nuevo
		l.tam++
	}
	
}

//Imptimir la lista
func (l *ListaSimple) Imprimir(){
	aux := l.inicio
	for aux != nil{
		fmt.Println(aux.producto)
		aux = aux.siguiente
	}
	fmt.Println("Tamaño del pedido en el carrito: ", l.tam)
	
}
//Decodificar Lista
func (l *ListaSimple)DecodificarLs()[]string{
	aux := l.inicio
	salida:=""
	var sal[] string = nil 
	for aux != nil{
		s, err := json.Marshal(aux.producto)
		if err != nil{
			log.Print(err)
		}
		salida = string(s)+"\n"
		sal = append(sal, salida)
		aux = aux.siguiente
	}
	return sal
}

//Buscar Elemento dentro de lista
func (l *ListaSimple) Buscar_Simple(codigo int) *NodoProd{
	aux := l.inicio
	for aux != nil {
		if aux.producto.Codigo == codigo {
			fmt.Println("Si se encontro el nodo")
			return aux
		}
		aux = aux.siguiente
	}
	fmt.Println("NO se encontro el nodo")
	return aux
}

/*
//Definir una funcion que nos devuelva el nodo en la posicion buscada
func (l *ListaSimple)GetAt(pos int)*NodoProd{
	aux := l.inicio
	if pos < 0 {
		return aux
	}
	if pos > (l.tam -1){
		return nil
	}
	for i := 0; i < pos; i++{
		aux = aux.siguiente
	}
	return aux 
}

//Eliminar nodo de la lista
func (l *ListaSimple) Eliminar_Simple(prod *Tiendas.Producto){
	borrar := l.Buscar_Simple(prod.Codigo)
	
	aux := l.inicio
	if l.tam == 0{
		fmt.Println("EL CARRITO DE COMPRAS ESTA VACIO")
	}
	for i := 0; i < l.tam; i++{
		if aux == borrar{
			if i > 0 {

				anterior := l.GetAt(i-1)
				anterior.siguiente =borrar.siguiente
			}else{
				l.inicio = aux.siguiente
			}
			l.tam--
			
		}
		aux = aux.siguiente
	}
	fmt.Println("No se encontro el nodo")

}*/


func (l *ListaSimple) Eliminar_Simple(prod *Tiendas.Producto){
	//Declarar un nodo auxiliar que busque el nodo y si existe para poder hacer las operaciones de comparacion
	aux := l.Buscar_Simple(prod.Codigo)
	//Si el nodo encontrado esta al inicio de la lista	
	if aux != nil{ 
		if l.inicio == l.ultimo{
			l.inicio = nil 
			l.ultimo = nil
		
		}else if l.inicio == aux{
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
	l.Imprimir()
}