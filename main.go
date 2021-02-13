package main

import (
	"fmt"
	"./Tiendas"
	"./Estructura"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	
	
)
//Declarar una variable de tipo arreglo para almacenar los indices
var index [] string
//Declarar una variable de tipo arreglo para almacenar los departamentos
var departamentos [] string
//Declarar una variable de tipo arreglo para almacenar las tiendas
var tiendas [] Tiendas.Tienda
//Declarar una variable de tipo arreglo para almacenar los datos
var datos [] Tiendas.Inicio
var data Tiendas.Inicio
//Declarar la variables de tipo ListaDoble
var lista *Estructura.Lista

//var tienda1 *Tiendas.Tienda
//var tienda2 *Tiendas.Tienda
//var tienda3 *Tiendas.Tienda
//var tienda4 *Tiendas.Tienda
//var tienda5 *Tiendas.Tienda


func main(){
	fmt.Println("Proyecto de Estructura de Datos, Fase 1")
	

	//Crear la lista 
	//lista = Estructura.Nueva_Lista()
	//Crear Tiendas
	//tienda1 = Tiendas.Nueva_Tienda(1,"Tienda 1", "Tienda de zapatos", "215189461", 5)
	//tienda2 = Tiendas.Nueva_Tienda(2,"Tienda 2", "Tienda de Ropa", "21564857", 4)
	//tienda3 = Tiendas.Nueva_Tienda(3,"Tienda 3", "Tienda de Animales", "786451864", 3)
	//tienda4 = Tiendas.Nueva_Tienda(4,"Tienda 4", "Tienda de Electronicos", "6549861", 2)
	//tienda5 = Tiendas.Nueva_Tienda(5,"Tienda 5", "Tienda de Mochilas", "5623189", 1)
	
	//lista.Insertar(tienda1)
	//lista.Insertar(tienda2)
	//lista.Insertar(tienda3)
	//lista.Insertar(tienda4)
	//lista.Insertar(tienda5)

	//lista.Imprimir()
	//lista.Buscar(1)
	//lista.Eliminar(2)
	//lista.Imprimir()

	//Crear el enrutador
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/cargartienda", CargarTiendas).Methods("POST")
	router.HandleFunc("/vertiendas", ConsultarTiendas).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
	

}
//Definir una funcion http GET para poder consultar el servidor y ver todas las tiendas
func ConsultarTiendas(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(datos)
}

//Definir una ruta de inicio
func indexRoute(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Proyecto fase 1 desde el servidor")
}

//Definir una funcion para Cargar las tiendas
func CargarTiendas(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Ingrese una tienda valida")
	}

	json.Unmarshal(reqBody, &data)

	var c_datos int = 0
	var c_dep int = 0
	var c_tiendas int = 0

	for _, d := range data.Data {
		index = append(index, d.Indice)
		for _, dep := range d.Departamentos {
			departamentos = append(departamentos, d.Departamentos[c_dep].Nombre)
			for _, t := range dep.Tiendas {
				t.Id = t.Id + t.GenerarId(t.Nombre)
				data.Data[c_datos].Departamentos[c_dep].Tiendas[c_tiendas].Id = t.GenerarId(t.Nombre)
				tiendas = append(tiendas, d.Departamentos[c_dep].Tiendas[c_tiendas])
				c_tiendas++
			}
			c_tiendas = 0
			c_dep++
		}
		c_dep = 0
		c_datos++
	}
	c_datos = 0

	//A fmt.Println(data)
	datos = append(datos, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(datos)

	
	for _, indice := range index{
		fmt.Println("INDICE: ", indice)
	}
	for _, depa := range departamentos{
		fmt.Println("DEPARTAMENTO: ", depa)
	}
	for _, tienda := range tiendas{
		fmt.Println("TIENDA: ", tienda)
	}

	depfinal := RemoveDuplicatesFromSlice(departamentos)
	departamentos = depfinal
	fmt.Println(departamentos)
	
}

//Definir una funcion que elimine los repetidos en el array de departamentos
func RemoveDuplicatesFromSlice(s []string) []string {
	m := make(map[string]bool)
	for _, item := range s {
			if _, ok := m[item]; ok {
					// duplicate item
					fmt.Println(item, "departamento duplicado")
			} else {
					m[item] = true
			}
	}

	var result []string
	for item, _ := range m {
			result = append(result, item)
	}
	return result
}
