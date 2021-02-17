//En este paquete definir una Lista Doblemente enlazada para su implementacion
package Estructura
//Importar los paquetes necesarios
import (
	"fmt"
	"../Tiendas"
	"log"
	"os"
	"os/exec"
	"strconv"
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
//Metodo que verifica si esta vacia
func (l *Lista)Es_Vacia() bool{
	if l.inicio == nil && l.ultimo == nil{
		return true
	}
	return false
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
func (l *Lista)Buscar(Nombre string) *Nodo{
	//definir un nodo auxiliar para recorrer la lista
	aux := l.inicio
	for aux != nil{
		//Verificar si el nodo en su propiedad de Id es igual al Id ingresado
		if aux.NodoTienda.Nombre == Nombre{
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
func (l *Lista) Eliminar(Nombre string){
	//Declarar un nodo auxiliar que busque el nodo y si existe para poder hacer las operaciones de comparacion
	aux := l.Buscar(Nombre)
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

//Definir una funcion para ordenar las listas por Id
func (l *Lista)Ordenar(){
	aux := l.inicio
	var temp *Tiendas.Tienda 
	for aux != nil{
		aux2 := aux.siguiente
		for aux2 != nil{
			if aux2.NodoTienda.Id < aux.NodoTienda.Id{
				temp = aux.NodoTienda
				aux.NodoTienda = aux2.NodoTienda
				aux2.NodoTienda = temp
			}
			

			aux2 = aux2.siguiente
		}
		aux = aux.siguiente
	}
}



//Definir una funcion para graficar el vector
func Graph(listas[] *Lista){
	os.Create("Estructura/GraficaPila.dot")
	graphdot := getFile("Estructura/GraficaPila.dot")
	fmt.Fprintf(graphdot,"digraph G {\n")
	fmt.Fprintf(graphdot,"rankdir = LR; \n")
	fmt.Fprintf(graphdot,"\tnode [shape=record, color=black]; \n")
	fmt.Fprintf(graphdot,"label = \"Estructura\";\n")
	fmt.Fprintf(graphdot,"color=black;\n")
	var text_aux string = ""
	var cont = 0
	var contador = 0
	
	for _, lis := range listas{
		if lis.Es_Vacia(){
			text_aux = "\t\tn_" + strconv.Itoa(cont) + "[label = \"NO HAY LISTA\"];\n"
			fmt.Fprintf(graphdot, text_aux)
			cont ++
		}else{
			text_aux = "\t\tn_" + strconv.Itoa(cont) + "[label = \"LISTA CON:"+ strconv.Itoa(lis.tam) +"ELEMENTOS\"];\n"
			fmt.Fprintf(graphdot, text_aux)
			
			aux := lis.inicio
			text_aux = "subgraph Lista_"+strconv.Itoa(contador)+"{\n"
			fmt.Fprintf(graphdot, text_aux)
			fmt.Fprintf(graphdot,"rankdir = UD; \n")
			fmt.Fprintf(graphdot, "\tnode [shape=record, fillcolor =\"blue\", style =\"filled\", color=black]; \n")
			for aux != nil{
				if aux.siguiente != nil{
					fmt.Fprintf(graphdot, "\""+ aux.NodoTienda.Nombre+"\"->\""+ aux.siguiente.NodoTienda.Nombre +"\";\n")
				}
				
				aux = aux.siguiente
				
			}
			contador ++
			text_aux = "n_" + strconv.Itoa(cont) + "->Lista_" + strconv.Itoa(contador-1)+ ";\n"
			fmt.Fprintf(graphdot, text_aux)
			text_aux = "Lista_" + strconv.Itoa(contador -1) + "->" + "\""+ lis.inicio.NodoTienda.Nombre+"\";\n"
			fmt.Fprintf(graphdot, text_aux)
			fmt.Fprintf(graphdot, "}\n")
			
			
			
			cont ++
		}
	}
	for i:= 0; i< cont-1; i++{
		text_aux = "n_" + strconv.Itoa(i) + "->n_" + strconv.Itoa(i+1)+ ";\n"
		fmt.Fprintf(graphdot, text_aux)
		
	}
	fmt.Fprintf(graphdot, "}\n")
	exec.Command("dot", "-Tpng", "Estructura/GraficaPila.dot", "-o", "Estructura/GraficaPila.png").Output()
	graphdot.Close()
}

//Definir una funcion para crear un archivo
func getFile(path string) *os.File{
	file, err := os.OpenFile(path, os.O_RDWR,0775)
	if err != nil{
		log.Fatal(err)
	}
	return file
}