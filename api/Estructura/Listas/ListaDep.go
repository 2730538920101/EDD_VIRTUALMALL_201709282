package Listas

import (
	"fmt"
	"../Matriz"
)

// Lugar donde almacenaremos la informacion
type nodomes struct {
	anterior  *nodomes
	siguiente *nodomes
	mes int
	Mat	*Matriz.Matriz

}

// Estructura para almacenar nodos de informacion
type ListaD struct {
	inicio *nodomes
	ultimo *nodomes
	tamd    int
}

// crear una nueva lista
func NewListaD() *ListaD {
	return &ListaD{nil, nil, 0}
}

//insertar un nodo

func (m *ListaD) InsertarD(valor int, cont int) {
	mat := Matriz.NewMatriz(cont) 
	nuevo := &nodomes{nil, nil, valor, mat}

	if m.inicio == nil {
		m.inicio = nuevo
		m.ultimo = nuevo
	} else {
		m.ultimo.siguiente = nuevo
		nuevo.anterior = m.ultimo
		m.ultimo = nuevo
	}
	m.tamd++
}

// imprimir la lista
func (m *ListaD) ImprimirD() {
	aux := m.inicio
	if m.tamd == 0 {
		fmt.Println("lista vacia")
	}else{
		for aux != nil {
			fmt.Print("<-[", aux.mes, "]->")
			aux = aux.siguiente
		}
		fmt.Println()
		fmt.Println("Tamaño lista = ", m.tamd)
	}

}

//Buscar Elemento dentro de lista
func (m *ListaD) BuscarD(valor int) *nodomes{
	aux := m.inicio
	for aux != nil {
		if aux.mes == valor{
			fmt.Println("Si se encontro el nodo")
			return aux
		}
		aux = aux.siguiente
	}
	fmt.Println("NO se encontro el nodo")
	return aux
}

//Eliminar nodo de la lista
func (m *ListaD) EliminarD(valor int) {
	aux := m.BuscarD(valor)
	if m.tamd == 0 {
		fmt.Println("lista vacia")
	}else{
		if m.tamd == 1{
			m.inicio = nil
			m.ultimo = nil
			fmt.Println("el nodo es unico")
			m.tamd--
		}else{
			if m.inicio == aux {

				m.inicio = aux.siguiente
				aux.siguiente.anterior = nil
				aux.siguiente = nil
				fmt.Println("el nodo es el inicio")
				m.tamd--
			} else if m.ultimo == aux {
					m.ultimo = aux.anterior
					aux.anterior.siguiente = nil
					aux.anterior = nil
					fmt.Println("el nodo es el final")
					m.tamd--
			} else {
					aux.anterior.siguiente = aux.siguiente
					aux.siguiente.anterior = aux.anterior
					aux.anterior = nil
					aux.siguiente = nil
					fmt.Println("el nodo es enmedio")
					m.tamd--
			}

		}


	}



}

//Definir una funcion para ordenar las listas por Id
func (l *ListaD)OrdenarD(){
	aux := l.inicio
	var temp int
	var temp2 *Matriz.Matriz 
	for aux != nil{
		aux2 := aux.siguiente
		for aux2 != nil{
			if aux2.mes < aux.mes{
				temp = aux.mes
				temp2 = aux.Mat
				aux.mes = aux2.mes
				aux.Mat =aux2.Mat
				aux2.mes = temp
				aux2.Mat =temp2 
			}
			

			aux2 = aux2.siguiente
		}
		aux = aux.siguiente
	}
}