//Se define una lista simple para el manejo de los pedidos en el carrito de compras
package Estructura


import (
	"../Tiendas"
	"fmt"
)

//Definir el nodo del producto en en carrito de compras
type NodoProd struct{
	siguiente *NodoProd
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
	nuevo := &NodoProd{nil, prod}
	if l.inicio == nil{
		l.inicio = nuevo
		l.ultimo = nuevo
	}else{
		l.ultimo.siguiente = nuevo
		l.ultimo = nuevo
	}
	l.tam++
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
func (l *ListaSimple) Eliminar_Simple(codigo int){
	borrar := l.Buscar_Simple(codigo)
	aux := l.inicio
	if l.tam == 0{
		fmt.Println("EL CARRITO DE COMPRAS NO ESTA VACIO")
	}
	for i := 0; i < l.tam; i++{
		if aux == borrar{
			if i > 0 {
				anterior := l.GetAt(i-1)
				anterior.siguiente =l.GetAt(i).siguiente
			}else{
				l.inicio = aux.siguiente
			}
			l.tam--
			
		}
		aux = aux.siguiente
	}
	fmt.Println("No se encontro el nodo")

}