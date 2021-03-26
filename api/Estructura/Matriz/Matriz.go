package Matriz

import (
	"fmt"
	"../Tiendas"
	"os"
	"log"
	"strconv"
	"os/exec"
)

type nodoInterno struct{ //nodo para guardar la informacion
	Valor *colaP
	X int
	Y int
	SiguienteX *nodoInterno
	AnteriorX *nodoInterno

	SiguienteY *nodoInterno
	AnteriorY *nodoInterno
}

type listainterna struct{ //lista que tendra cada cabecera de la matriz
	Primero * nodoInterno
}

type  nodoCabecera struct{ //nodo de la cabecesa de la condenada x o y
	Valor int
	Siguiente *nodoCabecera
	Anterior *nodoCabecera
	Lista *listainterna
}

type listaCabecera struct{ //lista para las cabeceras, se usa la misma para Y y X
	Primero *nodoCabecera
	Ultimo *nodoCabecera
}

type Matriz struct{ //nodo central de la matriz 
	Capa int
	CabecerasX *listaCabecera
	CabecerasY *listaCabecera
}

//Funciones para creacion de los elementos de la matriz

//Nueva Matriz
func NewMatriz(Valor int) * Matriz{
	cabeceraX := NewCabecera()
	cabeceraY := NewCabecera()
	fmt.Println("se creo una nueva matriz")
	return &Matriz{Valor, cabeceraX, cabeceraY}
}

//Nueva lista cabecera
func NewCabecera() * listaCabecera{
	return &listaCabecera{nil,nil}
}

//Nueva lista interna
func NewLista() * listainterna{
	return &listainterna{nil}
}

//Funciones para el manejo de la matriz

//Buscar cabecera en matriz
func (m *listaCabecera) buscar(pos int) *nodoCabecera {
	aux := m.Primero

	for aux != nil {
		if aux.Valor == pos {
			return aux
		}
		aux = aux.Siguiente
	}

	return nil
}

//Insertar en matriz
func (m *Matriz) Insertar(posx int, posy int, valor *colaP){
	//crear nuevo nodo
	nuevo := &nodoInterno{valor,posx,posy,nil,nil,nil,nil} //los 4 apuntadores se inician como nulos

	//buscar cabeceras de X en la matriz
	cabecerax := m.CabecerasX.buscar(posx) 
	//si las cabeceras no existen se deben crear
	if cabecerax == nil {
		//se crea la lista interna de las cabeceras
		lista := &listainterna{nil}
		cabecerax = &nodoCabecera{posx,nil,nil,lista}
		m.CabecerasX.Insertarcabecera(cabecerax)
	}

	//buscar cabeceras de Y en la matriz
	cabeceray := m.CabecerasY.buscar(posy)
	if (cabeceray == nil){
		listay := &listainterna{nil}
		cabeceray =  &nodoCabecera{posy,nil,nil,listay}
		m.CabecerasY.Insertarcabecera(cabeceray)
	}


	//Insertar el nodo a las cabeceras
	//Insertar en X
	listaX := cabecerax.Lista
	
	
	//**** SE DEBE INSERTAR ORDENADO en X comparando los valores en y
	if (listaX.Primero == nil){ //si es nulo solo se agrega
		listaX.Primero = nuevo
	}else{
		if (nuevo.Y < listaX.Primero.Y){ //el nodo debe ir al inicio de la lista
			nuevo.SiguienteX = listaX.Primero
			listaX.Primero.AnteriorX = nuevo
			listaX.Primero = nuevo
			//return
		}else { //se recorre la lista para insertar ordenado
			pivote := listaX.Primero

			for pivote != nil {
				if (nuevo.Y < pivote.Y){
					nuevo.SiguienteX = pivote
					nuevo.AnteriorX = pivote.AnteriorX
					pivote.AnteriorX.SiguienteX = nuevo
				}else if (nuevo.Y == pivote.Y && nuevo.X == pivote.X ){//comparacion para saber si no se ha insertado una mis posicion
					fmt.Println("Ya existe un nodo es estas coordenadas")
					break
				} else{ //else el y del nuevo es mayor al del pivote 
					if (pivote.SiguienteX == nil){ //se valida si se llego al ultimo 
						pivote.SiguienteX = nuevo //si el siginete es nil 
						nuevo.AnteriorX = pivote
						break
					}else{
						pivote = pivote.SiguienteX //si no es el ultimo nos pasamos al siguiente y vuelve a iterar el ciclo
					}
				}
			}
		}
	}

	//Insertar el nodo a las cabeceras
	//Insertar en Y
	listaY := cabeceray.Lista

	//**** SE DEBE INSERTAR ORDENADO en Y comparando los valores en X
	if (listaY.Primero == nil){ //si es nulo solo se agrega
		listaY.Primero = nuevo
	}else{
		fmt.Println("entro aqui ****")
		if (nuevo.X < listaY.Primero.X){ //el nodo debe ir al inicio de la lista
			fmt.Println("entro aqui *")
			nuevo.SiguienteY = listaY.Primero
			listaY.Primero.AnteriorY = nuevo
			listaY.Primero = nuevo
			//return
		}else { //se recorre la lista para insertar ordenado
			pivote := listaY.Primero

			for pivote != nil {
				fmt.Println("entro aqui")
				if (nuevo.X < pivote.X){
					nuevo.SiguienteY = pivote
					nuevo.AnteriorY = pivote.AnteriorY
					pivote.AnteriorY.SiguienteY = nuevo
				}else if (nuevo.Y == pivote.Y && nuevo.X == pivote.X ){//comparacion para saber si no se ha insertado una mis posicion
					fmt.Println("Ya existe un nodo es estas coordenadas")
					break
				} else{ //else el y del nuevo es mayor al del pivote 
					if (pivote.SiguienteY == nil){ //se valida si se llego al ultimo 
						pivote.SiguienteY = nuevo //si el siginete es nil 
						nuevo.AnteriorY = pivote
						break
					}else{
						pivote = pivote.SiguienteY //si no es el ultimo nos pasamos al siguiente y vuelve a iterar el ciclo
					}
				}
			}
		}
	}
	
}


//Insertar cabeceras
func (m * listaCabecera) Insertarcabecera(nuevo * nodoCabecera) {
	
	if m.Primero == nil {
		m.Primero = nuevo
		m.Ultimo = nuevo
	}else{
		if m.Primero == m.Ultimo { //solo hay un dato
			if nuevo.Valor > m.Primero.Valor {
				m.Primero.Siguiente = nuevo
				nuevo.Anterior = m.Primero
				m.Ultimo = nuevo
			}else if nuevo.Valor < m.Primero.Valor{
				nuevo.Siguiente = m.Primero
				m.Primero.Anterior = nuevo
				m.Primero = nuevo
			}
		}else { //hay mas de un dato
			if nuevo.Valor < m.Primero.Valor { //es menor al Primero 
				nuevo.Siguiente = m.Primero
				m.Primero.Anterior = nuevo
				m.Primero = nuevo
			}else if nuevo.Valor > m.Ultimo.Valor { // es mayor al Ultimo
				m.Ultimo.Siguiente = nuevo
				nuevo.Anterior = m.Ultimo
				m.Ultimo = nuevo
			}else {
				aux := m.Primero

				for aux != nil {
					if nuevo.Valor < aux.Valor {
						nuevo.Siguiente = aux
						nuevo.Anterior = aux.Anterior
						aux.Anterior.Siguiente = nuevo
						aux.Anterior = nuevo
						break
					}
				}
			}
		}
	}
}

func (m *Matriz) Comprobar(){
	listaAux := m.CabecerasX
	pivote := listaAux.Primero
	fmt.Println("**** recorrer cabeceras x ****")
	for pivote != nil {
		fmt.Println("cabecera: ",pivote.Valor)
		listaInterna := pivote.Lista
		pivoteInterno := listaInterna.Primero
		if pivoteInterno != nil {
			for pivoteInterno != nil {
				fmt.Println(pivoteInterno.Valor)
				pivoteInterno = pivoteInterno.SiguienteX
			}
			
		}
		pivote = pivote.Siguiente
	}
	
}

func (m *Matriz) MostarCabecerasX(){
	listaAux := m.CabecerasX
	aux := listaAux.Primero
	fmt.Println("**** Cabeceras en  X ****")
	if (aux == nil){
		fmt.Println("No existen cabeceras creadas X")
		return
	}

	for aux != nil {
		fmt.Println(aux.Valor)
		aux = aux.Siguiente
	}
}

func (m *Matriz) MostarCabecerasY(){
	listaAux := m.CabecerasY
	auxy := listaAux.Primero
	fmt.Println("**** Cabeceras en  Y ****")
	if (auxy == nil){
		fmt.Println("No existen cabeceras creadas Y")
		return
	}

	for auxy != nil {
		fmt.Println(auxy.Valor)
		auxy = auxy.Siguiente
	}
}

//Definir un metodo para graficar la matriz
func (m *Matriz)GraphM(){
	os.Create("Estructura/GraficaMatriz"+ strconv.Itoa(m.Capa)+".dot")
	graphdot := getFileM("Estructura/GraficaMatriz"+strconv.Itoa(m.Capa)+".dot")
	fmt.Fprintf(graphdot,"digraph G {\n")
	fmt.Fprintf(graphdot,"rankdir = TB; \n")
	fmt.Fprintf(graphdot, "node[shape = box, width=0.7, height=0.7, fillcolor=\"azure2\" color=\"white\" style= \"filled\"];\n")
	fmt.Fprintf(graphdot,"edge[style = \"bold\"];\n")
	fmt.Fprintf(graphdot, "\n\t node[label = \"Capa: %d\" fillcolor = \"darkolivegreen1\" pos= \"-1,1!\"]principal;\n", m.Capa)
	//graficar los nodos cabecera

	//Nodos Cabecera X:

	listaAux := m.CabecerasX
	aux := listaAux.Primero
	posx := 0
	for aux != nil{
		fmt.Fprintf(graphdot, "\n\t node[label = \"X: %d\" fillcolor= \"azure3\" pos= \"%d,1!\" shape = box]x%d;\n", aux.Valor, posx, aux.Valor)
		aux = aux.Siguiente
		posx++
	}
	aux = listaAux.Primero
	for aux.Siguiente != nil{
		fmt.Fprintf(graphdot, "x%d -> x%d; \n", aux.Valor, aux.Siguiente.Valor)
		fmt.Fprintf(graphdot, "x%d -> x%d; \n", aux.Siguiente.Valor, aux.Valor)
		aux = aux.Siguiente
	}
	fmt.Fprintf(graphdot, "principal -> x%d;\n", listaAux.Primero.Valor)

	//Nodos Cabecera Y:

	listaAuxY := m.CabecerasY
	auxy := listaAuxY.Primero
	posy := 0
	for auxy != nil{
		fmt.Fprintf(graphdot, "\n\t node[label = \"Y: %d\" fillcolor= \"azure3\" pos= \"-1,-%d!\" shape = box]y%d;\n", auxy.Valor, posy, auxy.Valor)
		auxy = auxy.Siguiente
		posy++
	}
	auxy = listaAuxY.Primero
	for auxy.Siguiente != nil{
		fmt.Fprintf(graphdot, "y%d -> y%d; \n", auxy.Valor, auxy.Siguiente.Valor)
		fmt.Fprintf(graphdot, "y%d -> y%d; \n", auxy.Siguiente.Valor, auxy.Valor)
		auxy = auxy.Siguiente
	}
	fmt.Fprintf(graphdot, "principal -> y%d;\n", listaAuxY.Primero.Valor)

	//Nodos Internos:
	
	pivoteX := listaAux.Primero
	posx = 0
	for pivoteX != nil{
		pivoteInterno := pivoteX.Lista.Primero
		for pivoteInterno != nil{
			pivoteY := listaAuxY.Primero
			posyI := 0
			for pivoteY != nil{
				if pivoteY.Valor == pivoteInterno.Y{
					break
				}
				posyI++
				pivoteY = pivoteY.Siguiente
			}
			fmt.Fprintf(graphdot, "\n\t node[label = \" %d, %d\n Cola tam: %d \" fillcolor= \"azure2\" pos = \" %d, -%d! \" shape = box]\"i%d-%d\";", pivoteInterno.X, pivoteInterno.Y, pivoteInterno.Valor.Size(), posx, posyI,pivoteInterno.X, pivoteInterno.Y )
			pivoteInterno = pivoteInterno.SiguienteX
		}
		pivoteInterno = pivoteX.Lista.Primero
		for pivoteInterno != nil{
			if pivoteInterno.SiguienteX != nil{
				fmt.Fprintf(graphdot, "\n \"i%d-%d\" -> \"i%d-%d\";\n", pivoteInterno.X, pivoteInterno.Y, pivoteInterno.SiguienteX.X, pivoteInterno.SiguienteY.Y)
				fmt.Fprintf(graphdot, "\n \"i%d-%d\" -> \"i%d-%d\";\n", pivoteInterno.SiguienteX.X, pivoteInterno.SiguienteY.Y, pivoteInterno.X, pivoteInterno.Y)
			}
			pivoteInterno = pivoteInterno.SiguienteX
		}
		fmt.Fprintf(graphdot, "\n x%d -> \"i%d-%d\";\n", pivoteX.Valor, pivoteX.Lista.Primero.X, pivoteX.Lista.Primero.Y)
		fmt.Fprintf(graphdot, "\n \"i%d-%d\" -> x%d; \n", pivoteX.Lista.Primero.X, pivoteX.Lista.Primero.Y, pivoteX.Valor)
		pivoteX = pivoteX.Siguiente
		posx++
	}
	pivoteY := listaAuxY.Primero
	for pivoteY != nil{
		pivoteInterno := pivoteY.Lista.Primero
		for pivoteInterno != nil{
			if pivoteInterno.SiguienteY != nil{
				fmt.Fprintf(graphdot, "\n \"i%d-%d\" -> \"i%d-%d\";\n", pivoteInterno.X, pivoteInterno.Y, pivoteInterno.SiguienteY.X, pivoteInterno.SiguienteY.Y)
				fmt.Fprintf(graphdot, "\n \"i%d-%d\" -> \"i%d-%d\";\n", pivoteInterno.SiguienteY.X, pivoteInterno.SiguienteY.Y, pivoteInterno.X, pivoteInterno.Y)
			}
			pivoteInterno = pivoteInterno.SiguienteY
		}
		fmt.Fprintf(graphdot, "\n y%d -> \"i%d-%d\"\n", pivoteY.Valor, pivoteY.Lista.Primero.X, pivoteY.Lista.Primero.Y)
		fmt.Fprintf(graphdot, "\n \"i%d-%d\" -> y%d \n", pivoteY.Lista.Primero.X, pivoteY.Lista.Primero.Y, pivoteY.Valor)
		pivoteY = pivoteY.Siguiente
	}
	fmt.Fprintf(graphdot, "}\n")
	exec.Command("neato", "-Tpng", "Estructura/GraficaMatriz"+ strconv.Itoa(m.Capa)+".dot", "-o", "Estructura/GraficaMatriz"+ strconv.Itoa(m.Capa)+".png").Output()
	graphdot.Close()



}


//Definir una funcion para crear un archivo
func getFileM(path string) *os.File{
	file, err := os.OpenFile(path, os.O_RDWR,0775)
	if err != nil{
		log.Fatal(err)
	}
	return file
}






//-----------------------------------------------------------------COLA----------------------------------//

type nodocolaP struct {
	siguiente *nodocolaP
	pedido     *Tiendas.Producto
}

type colaP struct {
	frente *nodocolaP
	fondo  *nodocolaP
	tam    int
}

// crear cola
func NewCola() *colaP {
	return &colaP{nil, nil, 0}
}

//Metodo para verificar si la cola esta vacía
func (m *colaP) ColaVacia() bool {
	return m.frente == nil
}

//Metodo para agregar datos a la cola
func (m *colaP) Encolar(pedido *Tiendas.Producto) {
	aux := &nodocolaP{nil, pedido}
	if m.ColaVacia() {
		m.frente = aux
		m.fondo = aux
	} else {
		m.fondo.siguiente = aux
		m.fondo = aux
	}
	m.tam++
}

//funcion para eliminar un dato de la cola
func (m *colaP) Desencolar() {
	if !m.ColaVacia() {
		aux := m.frente
		m.frente = aux.siguiente
		aux = nil
		m.tam--
	} else {
		fmt.Println("la cola esta vacia")
	}
}

//mostrar cola de datos
func (m *colaP) Mostrar() {
	if !m.ColaVacia() {
		aux := m.frente
		for aux != nil {
			fmt.Print("[", aux.pedido, "]->")
			aux = aux.siguiente
		}
		fmt.Println()
	} else {
		fmt.Println("la cola esta vacia")
	}
}

//Metodo para ver el frente de la cola
func (m *colaP) Front() *nodocolaP{
	fmt.Println("el frente es: ", m.frente)
	return m.frente
}

//Metodo para mostrar el fondo
func (m *colaP) Rear() *nodocolaP{
	fmt.Println("el ultimo nodo es: ", m.fondo.pedido)
	return m.fondo
}

//Metodo para vaciar la cola
func (m *colaP) Vaciar() {
	if !m.ColaVacia() {
		for !m.ColaVacia() {
			m.Desencolar()
		}
	} else {
		fmt.Println("la cola esta vacia")
	}

}

//Metodo para retornar el tamaño de la cola
func (m *colaP) Size() int {
	fmt.Println("el tamaño de la cola es: ", m.tam)
	return m.tam
}