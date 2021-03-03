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
	"strconv"
	
	
)


//Declarar una variable de tipo arreglo para almacenar los datos
var departamentos[] string
var index[] string
var tamdep int
var tamind int

var data Tiendas.Inicio
var vector []*Estructura.Lista
//Declarar la variables de tipo ListaDoble
var Tienda_Esp *Tiendas.Busc_Esp
var Tienda_Elim *Tiendas.Eliminar_Esp

func main(){
	fmt.Println("Proyecto de Estructura de Datos, Fase 1")
	//Crear el enrutador
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/cargartienda", CargarTiendas).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", TiendaEspecifica).Methods("POST")
	router.HandleFunc("/vertiendas", ConsultarTiendas).Methods("GET")
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/id/{Id}/", Buscar_Posicion).Methods("GET")
	router.HandleFunc("/Guardar", Guardar).Methods("GET")
	router.HandleFunc("/Eliminar", Elim_Tienda).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", router))
	

}
//Definir una funcion http GET para poder consultar el servidor y ver todas las tiendas
func ConsultarTiendas(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(data)
}

//Definir una ruta de inicio
func indexRoute(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Proyecto fase 1 desde el servidor")
}

//Definir una funcion para cargar la imagen del vector
func getArreglo(w http.ResponseWriter, r *http.Request){
	Estructura.Graph(vector)
	fmt.Fprintf(w, "Se ha generado exitosamente el grafico")
}
//Definir una funcion para decodificar el vector en formato json
func Guardar(w http.ResponseWriter, r *http.Request){
	for _, list := range vector{
		list.Decodificar(w,r)
	}
}
//Definir una funcion para Mostrar todas las tiendas en una posicion del vector cargada desde el servidor
func Buscar_Posicion(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	k, err := strconv.Atoi(vars["Id"])
	if err != nil{
		fmt.Fprintf(w, "La posicion ingresada no es valida")
	}
	fmt.Println("LA POSICION INGRESADA ES: ", k)
	
	
	if vector != nil{
		if vector[k] != nil{
			//Recorrer el vector para encontrar la tienda dentro de las listas
			vector[k].Decodificar(w, r)
		}else{
			fmt.Fprintf(w, "NO SE HA ENCONTRADO LA TIENDA SOLICITADA")
		}
		
	
	}else{
		fmt.Fprintf(w, "Aun no ha cargado la informacion al vector")
	}
}
//Definir una funcion para Buscar por tienda especifica
func Elim_Tienda(w http.ResponseWriter, r *http.Request){
	reqBody,err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No ha ingresado la informacion correctamente")
	}
	json.Unmarshal(reqBody, &Tienda_Elim)
	var ind int
	var dep int
	var cal = Tienda_Elim.Calificacion-1
	//Obtener todos los indices y guardar la posicion donde se encontro coincidencia
	for a, indice := range index{
		b := []byte(Tienda_Elim.Nombre)
		indi := string(b[0])
		if indice == indi{
			ind = a
			
		}
	}
	//Obtener todos los departamentos y guardar las posiciones donde hay coincidencias
	for p, depa := range departamentos{
		if depa == Tienda_Elim.Categoria{
			dep = p
			
		}
	}
	fmt.Println(ind)
	fmt.Println(dep)
	fmt.Println(cal)
	//Definir la ecuacion para encontrar elementos en el vector
	pos := ((dep*(tamind)+ind)*(5))+cal
	if vector != nil{
		//Recorrer el vector para encontrar la tienda dentro de las listas
		if vector[pos] != nil{
			vector[pos].Eliminar(Tienda_Elim.Nombre)
			fmt.Fprintf(w, "EL REGISTRO DE LA TIENDA HA SIDO ELIMINADO DE LA LISTA")
		}else{
			fmt.Fprintf(w, "EL REGISTRO QUE SOLICITA ELIMINAR NO EXISTE")
		}
	
	}else{
		fmt.Fprintf(w, "Aun no ha cargado la informacion al vector")
	}
	
}


//Definir una funcion para Buscar por tienda especifica
func TiendaEspecifica(w http.ResponseWriter, r *http.Request){
	reqBody,err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No ha ingresado la informacion correctamente")
	}
	json.Unmarshal(reqBody, &Tienda_Esp)
	var ind int
	var dep int
	var cal = Tienda_Esp.Calificacion-1
	//Obtener todos los indices y guardar la posicion donde se encontro coincidencia
	for a, indice := range index{
		b := []byte(Tienda_Esp.Nombre)
		indi := string(b[0])
		if indice == indi{
			ind = a
			
		}
	}
	//Obtener todos los departamentos y guardar las posiciones donde hay coincidencias
	for p, depa := range departamentos{
		if depa == Tienda_Esp.Departamento{
			dep = p
			
		}
	}
	fmt.Println(ind)
	fmt.Println(dep)
	fmt.Println(cal)
	tamind = len(index)
	//Definir la ecuacion para encontrar elementos en el vector
	k := ((dep*(tamind)+ind)*(5))+cal
	fmt.Println(k)
	
	
	
	if vector != nil{
		//Recorrer el vector para encontrar la tienda dentro de las listas
		info := vector[k].Buscar(Tienda_Esp.Nombre)
		if info != nil{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(info)
		}
	
	}else{
		fmt.Fprintf(w, "Aun no ha cargado la informacion al vector")
	}
	
}

//Definir una funcion para Cargar las tiendas
func CargarTiendas(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Ingrese una tienda valida")
	}

	json.Unmarshal(reqBody, &data)
	cantindex := len(data.Data)
	cantdep := len(data.Data[0].Departamentos)
	matrix := make([][]*Estructura.NodoM, cantindex)
	for i:= 0; i < cantindex; i++{
		matrix[i] = make([]*Estructura.NodoM, cantdep)
		letra := string((data.Data[i].Indice)[0])
		index = append(index, letra)
		
		for j := 0; j < cantdep; j++ {
			departamento := data.Data[i].Departamentos[j].Nombre
			departamentos = append(departamentos, departamento)
			espacio := *Estructura.NuevoNM(letra, departamento)
			cant_t := len(data.Data[i].Departamentos[j].Tiendas)
			
			for z := 0; z < cant_t; z++{
				tienda := data.Data[i].Departamentos[j].Tiendas[z]
				nombre := tienda.Nombre
				data.Data[i].Departamentos[j].Tiendas[z].Id= data.Data[i].Departamentos[j].Tiendas[z].GenerarId(nombre)
				tienda.Id= data.Data[i].Departamentos[j].Tiendas[z].GenerarId(nombre)
				//desc := tienda.Descripcion
				//contact := tienda.Contacto
				cali := tienda.Calificacion
				
				
				if cali == 1{
					espacio.List[0].Insertar(&tienda)
					espacio.List[0].Ordenar()
				}else if cali == 2{
					espacio.List[1].Insertar(&tienda)
					espacio.List[1].Ordenar()
				}else if cali == 3{
					espacio.List[2].Insertar(&tienda)
					espacio.List[2].Ordenar()
				}else if cali == 4{
					espacio.List[3].Insertar(&tienda)
					espacio.List[3].Ordenar()
				}else if cali == 5{
					espacio.List[4].Insertar(&tienda)
					espacio.List[4].Ordenar()
				}
				
			}
			matrix[i][j] = &espacio 	
			
		}
	}
	for i:=0; i< cantindex; i++{
		for j:=0; j< cantdep; j++{
			for k:=0; k<5; k++{
				vector = append(vector, matrix[i][j].List[k])
			}
		}
	}
	departamentos = RemoveDuplicatesFromSlice(departamentos)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
	for cont, l := range vector{
		l.Imprimir()
		fmt.Println("POS:", cont)
	}
	
}
//Definir una funcion que elimine los repetidos en el array de departamentos
func RemoveDuplicatesFromSlice(s []string) []string {
	m := make(map[string]bool)
	for _, item := range s {
			if _, ok := m[item]; ok {
					// duplicate item
					fmt.Println(item, "indice duplicado")
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

