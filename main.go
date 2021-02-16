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
var tamind int
//Declarar una variable de tipo arreglo para almacenar los departamentos
var departamentos [] string
var tamdep int

//Declarar una variable de tipo arreglo para almacenar los datos
var datos [] Tiendas.Inicio
var data Tiendas.Inicio
//Declarar la variables de tipo ListaDoble
var lista1 *Estructura.Lista
var lista2 *Estructura.Lista
var lista3 *Estructura.Lista
var lista4 *Estructura.Lista
var lista5 *Estructura.Lista

var listas [] *Estructura.Lista


var tienda1 *Tiendas.Tienda
var tienda2 *Tiendas.Tienda
var tienda3 *Tiendas.Tienda
var tienda4 *Tiendas.Tienda
var tienda5 *Tiendas.Tienda


func main(){
	fmt.Println("Proyecto de Estructura de Datos, Fase 1")
	//Crear el enrutador
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/cargartienda", CargarTiendas).Methods("POST")
	router.HandleFunc("/vertiendas", ConsultarTiendas).Methods("GET")
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
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

//Definir una funcion para cargar la imagen del vector
func getArreglo(w http.ResponseWriter, r *http.Request){
	Estructura.Graph(listas)
	fmt.Fprintf(w, "Se ha generado exitosamente el grafico")
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
			//departamentos = append(departamentos, d.Departamentos[c_dep].Nombre)
			lista1 = Estructura.Nueva_Lista()
			lista2 = Estructura.Nueva_Lista()
			lista3 = Estructura.Nueva_Lista()
			lista4 = Estructura.Nueva_Lista()
			lista5 = Estructura.Nueva_Lista()
			
			departamentos = append(departamentos, data.Data[c_datos].Departamentos[c_dep].Nombre)
			for _, t := range dep.Tiendas {
				
				data.Data[c_datos].Departamentos[c_dep].Tiendas[c_tiendas].Id = t.GenerarId(t.Nombre)
				//tiendas = append(tiendas, d.Departamentos[c_dep].Tiendas[c_tiendas])
				if data.Data[c_datos].Departamentos[c_dep].Tiendas[c_tiendas].Calificacion == 1{
					lista1.Insertar(&d.Departamentos[c_dep].Tiendas[c_tiendas])	
				}else if data.Data[c_datos].Departamentos[c_dep].Tiendas[c_tiendas].Calificacion == 2{
					lista2.Insertar(&d.Departamentos[c_dep].Tiendas[c_tiendas])
				}else if data.Data[c_datos].Departamentos[c_dep].Tiendas[c_tiendas].Calificacion == 3{
					lista3.Insertar(&d.Departamentos[c_dep].Tiendas[c_tiendas])
				}else if data.Data[c_datos].Departamentos[c_dep].Tiendas[c_tiendas].Calificacion == 4{
					lista4.Insertar(&d.Departamentos[c_dep].Tiendas[c_tiendas])
				}else if data.Data[c_datos].Departamentos[c_dep].Tiendas[c_tiendas].Calificacion == 5{
					lista5.Insertar(&d.Departamentos[c_dep].Tiendas[c_tiendas])
				}else{
					fmt.Println("El campo esta vacio o pudo haber ocurrido un error")
				}
				c_tiendas ++
			}
			
			c_tiendas = 0
			c_dep++
			lista1.Ordenar()
			lista2.Ordenar()
			lista3.Ordenar()
			lista4.Ordenar()
			lista5.Ordenar()
			listas = append(listas, lista1)
			listas = append(listas, lista2)
			listas = append(listas, lista3)
			listas = append(listas, lista4)
			listas = append(listas, lista5)

		}
		c_dep = 0
		c_datos++
	}
	c_datos = 0

	
	datos = append(datos, data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(datos)

	for pos, lis := range listas{
		lis.Imprimir()
		fmt.Println("LA POSICION DE ESTA LISTA ES: ", pos)
	}

	depfinal := RemoveDuplicatesFromSlice(departamentos)
	departamentos = depfinal
	tamdep = len(departamentos)
	tamind = len(index)
	fmt.Println("INDEICES: ",tamind,"DEPARTAMENTOS: ",tamdep)
	
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