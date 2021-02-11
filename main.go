package main

import (
	"fmt"
	"./Tiendas"
	"./Estructura"
)

func main(){
	fmt.Println("Proyecto de Estructura de Datos, Fase 1")
	//Declarar la variables de tipo ListaDoble
	var lista *Estructura.Lista
	//Declarar las variables de tipo Tienda 
	var tienda1 *Tiendas.Tienda
	var tienda2 *Tiendas.Tienda
	var tienda3 *Tiendas.Tienda
	var tienda4 *Tiendas.Tienda
	var tienda5 *Tiendas.Tienda

	//Crear la lista 
	lista = Estructura.Nueva_Lista()
	//Crear Tiendas
	tienda1 = Tiendas.Nueva_Tienda(1,"Tienda 1", "Tienda de zapatos", "215189461", 5)
	tienda2 = Tiendas.Nueva_Tienda(2,"Tienda 2", "Tienda de Ropa", "21564857", 4)
	tienda3 = Tiendas.Nueva_Tienda(3,"Tienda 3", "Tienda de Animales", "786451864", 3)
	tienda4 = Tiendas.Nueva_Tienda(4,"Tienda 4", "Tienda de Electronicos", "6549861", 2)
	tienda5 = Tiendas.Nueva_Tienda(5,"Tienda 5", "Tienda de Mochilas", "5623189", 1)

	lista.Insertar(tienda1)
	lista.Insertar(tienda2)
	lista.Insertar(tienda3)
	lista.Insertar(tienda4)
	lista.Insertar(tienda5)

	lista.Imprimir()
	lista.Buscar(1)
	lista.Eliminar(2)
	lista.Imprimir()
	

}