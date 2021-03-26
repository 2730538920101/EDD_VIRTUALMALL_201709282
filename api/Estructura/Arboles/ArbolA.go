package Arboles

import (
	"fmt"
	"../Listas"
)

// nodo del arbol
type NodoAVL2 struct {
	Valor2 int
	Meses *Listas.ListaD 
	izq2   *NodoAVL2
	der2   *NodoAVL2
	feq2   int
}

//arbol binario
type Arbol2 struct {
	raiz2 *NodoAVL2
	tam2  int
}

//creando nuevo arbol
func NewArbol2() *Arbol2 {
	return &Arbol2{nil, 0}
}

//Variable estatica que ayuda a verificar si el nodo existe o no
var Yaesta2 bool

//metodo para insertar un nuevo nodo
func (ar *Arbol2) InsertarNodoAVL2(valor int) bool {
	Lista := Listas.NewListaD()
	nuevo := &NodoAVL2{valor,Lista, nil, nil,0}
	Yaesta2 = false
	if ar.raiz2 == nil {
		fmt.Println("se inserto la raiz")
		ar.raiz2 = nuevo
		ar.tam2++
	} else {
		// LLamar metodo de insercion recursiva
		insertarAVL2(ar, ar.raiz2, nuevo)

		return Yaesta2
	}
	return Yaesta2
}

func insertarAVL2(ar *Arbol2, raiz *NodoAVL2, nuevo *NodoAVL2) {
	if nuevo.Valor2 > raiz.Valor2 {
		if raiz.der2 == nil {
			fmt.Println("se inserto a la derecha")
			raiz.der2 = nuevo
		} else {
			// verificar nueva insecion
			insertarAVL2(ar, raiz.der2, nuevo)
		}
	} else if nuevo.Valor2 < raiz.Valor2 {
		if raiz.izq2 == nil {
			fmt.Println("se inserto a la izquierda")
			raiz.izq2 = nuevo

		} else {
			// verificar nueva insercion
			insertarAVL2(ar, raiz.izq2, nuevo)
		}
	} else if nuevo.Valor2 == raiz.Valor2 {
		fmt.Println("el nodo ya existe")
		Yaesta2 = true
	}
	Equilibrar2(ar, raiz)
}

// Metodo para obtener un arbol de manera recursiva
func ObtenerNodo2(raiz *NodoAVL2, valor int) *NodoAVL2 {
	if raiz == nil {
		return nil
	} else if raiz.Valor2 == valor {
		fmt.Println("se encontro el nodo")
		return raiz
	} else {
		var valor1 *NodoAVL2
		if valor > raiz.Valor2 {
			valor1 = ObtenerNodo2(raiz.der2, valor)
		} else if valor < raiz.Valor2 {
			valor1 = ObtenerNodo2(raiz.izq2, valor)
		}
		return valor1
	}
}

// Metodo que obtiene el nodo como tal
func (ar *Arbol2) Obtener2(valor int) *NodoAVL2 {
	var retornado = ObtenerNodo2(ar.raiz2, valor)
	fmt.Println(retornado)
	return retornado
}




//funcion para retornar la cantidad de nodos del arbol
func (ar *Arbol2) CantidadNodos2() int {
	return ar.tam2
}

//Metodo recursivo de profundidad de un arbol
func profundidad2(raiz *NodoAVL2) int {
	if raiz == nil {
		return 0
	} else {
		var profizq = profundidad2(raiz.izq2)
		var profder = profundidad2(raiz.der2)

		if profizq > profder {

			return profizq + 1
		} else {
			return profder + 1
		}
	}
}

// Metodo para retornar el valor de la profundidad
func (ar *Arbol2) RetornarProf2() int {
	profundidad := profundidad2(ar.raiz2)
	fmt.Println("La profundidad es: ", profundidad)
	return profundidad
}

/*****************************RECORRIDOS**********************/
//Metodo recursivo para recorrer Inorden
// IZQ - RAIZ - DERECHA
func inorden2(raia *NodoAVL2) {
	if raia.izq2 != nil {
		inorden2(raia.izq2)
	}

	fmt.Print("nodo: ", raia.Valor2, "    ")

	if raia.der2 != nil {
		inorden2(raia.der2)
	}
}

//metodo que retorna el recorrido
func (ar *Arbol2) RecorridoInorden2() {
	inorden2(ar.raiz2)
	fmt.Println("Termino el recorrido")
}

//Metodo recursivo para recorre Preorder
func Preorder2(raiz *NodoAVL2) {
	fmt.Print("nodo: ", raiz.Valor2, "    ")

	if raiz.izq2 != nil {
		Preorder2(raiz.izq2)
	}

	if raiz.der2 != nil {
		Preorder2(raiz.der2)
	}
}

//metodo que retorna el recorrido
func (ar *Arbol2) RecorridoPreorden2() {
	Preorder2(ar.raiz2)
	fmt.Println("Termino el recorrido")
}

//Metodo recursivo para recorre Postorden
func postorden2(raiz *NodoAVL2) {

	if raiz.izq2 != nil {
		postorden2(raiz.izq2)
	}

	if raiz.der2 != nil {
		postorden2(raiz.der2)
	}

	fmt.Print("nodo: ", raiz.Valor2, "    ")
}

//metodo que retorna el recorrido
func (ar *Arbol2) RecorridoPostorden2() {
	postorden2(ar.raiz2)
	fmt.Println("Termino el recorrido")
}


//Metodo para obtener el padre
func obtenerPadreAVL2(raiz *NodoAVL2, valor int) *NodoAVL2{
	if valor > raiz.Valor2{
		if valor == raiz.der2.Valor2{
			return raiz
		}else{
		 return	obtenerPadreAVL2(raiz.der2, valor)
		}
	}else if valor < raiz.Valor2 {
		if valor == raiz.izq2.Valor2 {
			return raiz
		}else{
		 return	obtenerPadreAVL2(raiz.izq2, valor)
		}
	}else{
		return nil
	}
}

//Rotaciones
func  rotII2(ar *Arbol2, n *NodoAVL2, n1 *NodoAVL2){
	n.izq2 = n1.der2
	n1.der2 = n

	if n1.feq2 == -1{
		n.feq2 =0
		n1.feq2 = 0
	}else{
		n.feq2 = -1
		n1.feq2 = 0
	}
	if ar.raiz2 == n{
		n = n1
		ar.raiz2 = n1
	}else{
		temp := obtenerPadreAVL2(ar.raiz2, n.Valor2)
		if temp.izq2 == n{
			temp.izq2 = n1
		}else{
			temp.der2 = n1
		}
	}
	fmt.Println("Se realizo rotacion II ")
}

func  rotDD2( ar *Arbol2, n *NodoAVL2, n1 *NodoAVL2){
	n.der2 = n1.izq2
	n.izq2 = n

	if n1.feq2 == 1 {
		n.feq2 =0
		n1.feq2 = 0
	}else{
		n.feq2 = 1
		n1.feq2 = 0
	}

	if ar.raiz2 == n {
		n = n1
		ar.raiz2 = n1
	}else{
		temp := obtenerPadreAVL2(ar.raiz2, n.Valor2)
		if temp.izq2 == n{
			temp.izq2 = n1
		}else{
			temp.der2 = n1
		}
	}
	fmt.Println("Se realizo rotacion DD ")
}

func rotID2( ar *Arbol2, n *NodoAVL2, n1 *NodoAVL2, n2 *NodoAVL2){
	n.izq2 = n2.der2
	n2.der2 = n
	n1.der2 = n2.izq2
	n2.izq2 = n1

	if n2.feq2 == 1 {
		n1.feq2 = -1
	}else{
		n1.feq2 = 0
	}

	if n2.feq2 == -1{
		n.feq2 = 1
	}else{
		n.feq2 = 0
	}
	n2.feq2 = 0

	if ar.raiz2 == n {
		n = n2
		ar.raiz2 = n2
	}else{
		temp := obtenerPadreAVL2(ar.raiz2, n.Valor2)
		if temp.izq2 == n{
			temp.izq2 = n2
		}else{
			temp.der2 = n2
		}
	}

	fmt.Println("Se realizo rotacion ID ")
}

func  rotDI2(ar *Arbol2, n *NodoAVL2, n1 *NodoAVL2, n2 *NodoAVL2){
	n.der2 = n2.izq2
	n2.izq2 = n
	n1.izq2 = n2.der2
	n2.der2 = n1

	if n2.feq2 == 1{
		n.feq2 = -1
	}else{
		n.feq2 = 0
	}

	if n2.feq2 == -1 {
		n1.feq2 = 1
	}else{
		n1.feq2 = 0
	}
	n2.feq2 = 0

	if ar.raiz2 == n{
		n = n2
		ar.raiz2 = n2
	}else{
		temp := obtenerPadreAVL2(ar.raiz2, n.Valor2)
		if temp.izq2 == n {
			temp.izq2 = n2
		}else{
			temp.der2 = n2
		}
	}

	fmt.Println("Se realizo rotacion DI ")
}

// Metodo que sirve para equilibrar el arbol
func  Equilibrar2( ar * Arbol2, raiz *NodoAVL2){
	izq := profundidad2(raiz.izq2)
	der := profundidad2(raiz.der2)

	raiz.feq2 = der - izq

	if raiz.feq2 == -2{
		if raiz.izq2.feq2 > 0{
			rotID2(ar, raiz, raiz.izq2, raiz.izq2.der2)
		}else{
			rotII2(ar, raiz, raiz.izq2)
		}
	}else if raiz.feq2 == 2{
		if raiz.der2.feq2 < 0{
			rotDI2(ar, raiz, raiz.der2, raiz.der2.izq2)
		}else{
			rotDD2(ar, raiz, raiz.der2)
		}
	}
}