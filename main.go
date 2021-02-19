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
var datos [] Tiendas.Inicio
var data Tiendas.Inicio
//Declarar la variables de tipo ListaDoble
var lista1 *Estructura.Lista
var lista2 *Estructura.Lista
var lista3 *Estructura.Lista
var lista4 *Estructura.Lista
var lista5 *Estructura.Lista
var Tienda_Esp *Tiendas.Busc_Esp
var Tienda_Elim *Tiendas.Eliminar_Esp
var listas [] *Estructura.Lista

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
	router.HandleFunc("/Eliminar", Elim_Tienda).Methods("POST")
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
//Definir una funcion para decodificar el vector en formato json
func Guardar(w http.ResponseWriter, r *http.Request){
	for _, list := range listas{
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
	
	
	if listas != nil{
		if listas[k] != nil{
			//Recorrer el vector para encontrar la tienda dentro de las listas
			listas[k].Decodificar(w, r)
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
	if listas != nil{
		//Recorrer el vector para encontrar la tienda dentro de las listas
		if listas[pos] != nil{
			listas[pos].Eliminar(Tienda_Elim.Nombre)
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
	//Definir la ecuacion para encontrar elementos en el vector
	k := ((dep*(tamind)+ind)*(5))+cal
	
	fmt.Println(k)
	
	
	
	if listas != nil{
		//Recorrer el vector para encontrar la tienda dentro de las listas
		info := listas[k].Buscar(Tienda_Esp.Nombre)
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

	var c_datos int = 0
	var c_dep int = 0
	var c_tiendas int = 0
	
	for _, d := range data.Data[c_datos].Departamentos{
		departamentos = append(departamentos, d.Nombre)
		for _, f := range data.Data{
			index = append(index, f.Indice)
			lista1 = Estructura.Nueva_Lista()
			lista2 = Estructura.Nueva_Lista()
			lista3 = Estructura.Nueva_Lista()
			lista4 = Estructura.Nueva_Lista()
			lista5 = Estructura.Nueva_Lista()
		
			fmt.Println(f.Indice)
			fmt.Println(d.Nombre)
			for c_tiendas, t := range f.Departamentos[c_dep].Tiendas{
				f.Departamentos[c_dep].Tiendas[c_tiendas].Id = t.GenerarId(f.Departamentos[c_dep].Tiendas[c_tiendas].Nombre)
			}
			if f.Departamentos[c_dep].Tiendas[c_tiendas].Calificacion == 1{
				lista1.Insertar(&f.Departamentos[c_dep].Tiendas[c_tiendas])
			}else if f.Departamentos[c_dep].Tiendas[c_tiendas].Calificacion == 2{
				lista2.Insertar(&f.Departamentos[c_dep].Tiendas[c_tiendas])
			}else if f.Departamentos[c_dep].Tiendas[c_tiendas].Calificacion == 3{
				lista3.Insertar(&f.Departamentos[c_dep].Tiendas[c_tiendas])
			}else if f.Departamentos[c_dep].Tiendas[c_tiendas].Calificacion == 4{
				lista4.Insertar(&f.Departamentos[c_dep].Tiendas[c_tiendas])
			}else if f.Departamentos[c_dep].Tiendas[c_tiendas].Calificacion == 5{
				lista5.Insertar(&f.Departamentos[c_dep].Tiendas[c_tiendas])
			}else{
				fmt.Println("El campo esta vacio o pudo haber ocurrido un error")
			}
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
			c_datos++
		}
		
		c_tiendas = 0
		c_dep++
		
		
		
	}
		
	datos = append(datos, data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(datos)

	for pos, lis := range listas{
		lis.Imprimir()
		fmt.Println("LA POSICION DE ESTA LISTA ES: ", pos)
	}
	
	indexfinal := RemoveDuplicatesFromSlice(index)
	index = indexfinal
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

