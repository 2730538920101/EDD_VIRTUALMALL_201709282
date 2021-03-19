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
	"strings"
	
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

//variable para cargar inventarios
var Inventario Tiendas.InvInit

//variable para cargar pedidos
var Pedidos *Tiendas.PedInit
//variables para los productos de prueba
var producto1 *Tiendas.Producto
var producto2 *Tiendas.Producto
var producto3 *Tiendas.Producto
var ListaS *Estructura.ListaSimple
var tienda *Tiendas.Tienda 

var Prod *Tiendas.Carrito
var ProdAvl []*Estructura.Arbol
var tind map[string]interface{}

func main(){
	fmt.Println("Proyecto de Estructura de Datos, Fase 1")
	//Crear el enrutador
	
	ListaS = Estructura.Nueva_ListaS()
	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/api/cargartienda", CargarTiendas).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", TiendaEspecifica).Methods("POST")
	router.HandleFunc("/api/vertiendas", ConsultarTiendas).Methods("GET")
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/id/{Id}/", Buscar_Posicion).Methods("GET")
	router.HandleFunc("/Guardar", Guardar).Methods("GET")
	router.HandleFunc("/Eliminar", Elim_Tienda).Methods("DELETE")
	router.HandleFunc("/CargarInventario",CargarInventario).Methods("POST")
	router.HandleFunc("/api/verInventario/{id}", verInventario).Methods("GET")
	router.HandleFunc("/api/verArbol",verArbol).Methods("GET")
	router.HandleFunc("/CargarPedido",CargarPedido).Methods("POST")
	router.HandleFunc("/cargar", Prueba).Methods("POST")
	router.HandleFunc("/api/ver",verCargar).Methods("GET")
	router.HandleFunc("/InsertarCarrito",InsertarCarrito).Methods("PUT")
	
	log.Fatal(http.ListenAndServe(":3000", router))
	
}
//Definir un metodo para insertar en el carrito
func InsertarCarrito(w http.ResponseWriter, r *http.Request){
	reqBody,err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No ha ingresado la informacion correctamente")
	}
	json.Unmarshal(reqBody, &Prod)
	var ind int = 0
	tiendanom := Prod.Tienda
	for a, indice := range index{
		b := []byte(tiendanom)
		indi := string(b[0])
		if indice == indi{
			ind = a
			
		}
	}
	dep := Prod.Dep
	cal := Prod.Calificacion-1
	pos := ((dep*(tamind)+ind)*(5))+cal
	cod := Prod.Codigo
	nom := Prod.Nombre
	cant := Prod.Cantidad
	productoFinal :=vector[pos].BuscarProd(tiendanom, cod, nom, cant)
	if productoFinal !=nil{
		ListaS.Insertar(productoFinal)
		ListaS.Imprimir()
	}else{
		fmt.Println("NO HAY PRODUCTO DISPONIBLE")
	}
	

	_= json.NewDecoder(r.Body).Decode(&Prod)	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&Prod)
}





//Definir un metodo post para cargar los pedidos
func CargarPedido(w http.ResponseWriter, r *http.Request){
	reqBody,err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No ha ingresado la informacion correctamente")
	}
	json.Unmarshal(reqBody, &Pedidos)
	fmt.Println(&Pedidos)
	_= json.NewDecoder(r.Body).Decode(&Pedidos)	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&Pedidos)
}
//Definir una funcion para graficar los arboles
func verArbol(w http.ResponseWriter, r *http.Request){
	Estructura.GraficarAvl(ProdAvl)
	fmt.Fprintf(w,"SE HA CREADO EL GRAFICO DE LOS ARBOLES AVL")
}

//Definir un metodo post para cargar los inventarios

func verInventario(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	
	TiendaId, err := strconv.Atoi(vars["id"])
	if err != nil{
		fmt.Fprintf(w, "EL ID DE LA TIENDA ES INVALIDO")
		return
	}
	fmt.Println("LA POSICION INGRESADA ES: ", TiendaId)
	for i:=0; i<len(vector); i++{
		busc := vector[i].BuscarId(TiendaId)
		if busc != nil{
			fmt.Println("PRODUCTOS: ", busc.NodoTienda.Productos)
				
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(busc.NodoTienda)
					
		}
	}

}



func CargarInventario(w http.ResponseWriter, r *http.Request){
	
	reqBody,err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No ha ingresado la informacion correctamente")
	}
	json.Unmarshal(reqBody, &Inventario)
	
		//debo leer prod
	for i:=0; i<len(Inventario.Inventarios);i++{
		Tiendanom := Inventario.Inventarios[i].Tienda
		Tiendadep := Inventario.Inventarios[i].Departamento
		Tiendacal := Inventario.Inventarios[i].Calificacion
		Tiendaprod := Inventario.Inventarios[i].Productos
		var ind int = 0
		var dep int = 0
		var cal int = 0
		for a, indice := range index{
			b := []byte(Tiendanom)
			indi := string(b[0])
			if indice == indi{
				ind = a
				
			}
		}
		for p, depa := range departamentos{
			if depa == Tiendadep{
				dep = p
				
			}
		}
		cal = Tiendacal -1
		fmt.Println("IND: ", ind)
		fmt.Println("DEP: ", dep)
		fmt.Println("CAL: ", cal)
		fmt.Println("TAMIND: ", tamind)
		pos := ((dep*(tamind)+ind)*(5))+cal
		fmt.Println("POS: ", pos)
		arbol := Estructura.NewArbol()
		for j:=0; j<len(Tiendaprod);j++{
			arbol.InsertarNodoAVL(&Tiendaprod[j])
			
		}
		ProdAvl = append(ProdAvl, arbol)
		fmt.Println("Se agregaron productos en: ", pos)
		vector[pos].AgregarProd(Tiendaprod,Tiendanom, Tiendacal)
	}

	for _, v:= range vector{
		v.Imprimir()
	}
	_= json.NewDecoder(r.Body).Decode(&Inventario)	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&Inventario)
}


//Definir una funcion http GET para poder consultar el servidor y ver todas las tiendas
func ConsultarTiendas(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&tind)
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
	depli = nil
	tind = nil
	for _, li:= range vector{
		salida := li.Decodificar(w,r)
		if salida !=""{
			depli = append(depli, salida )	
		}	
	}
	
	stjs := strings.Join(depli,",")
	stjs = "{\n\"tiendas\":[\n"+stjs+"]\n}"
	sal := []byte(stjs)
	err2 := json.Unmarshal(sal, &tind)
	if err2 != nil {
		log.Print(err2)
	}
	fmt.Println("BANDERA")
	fmt.Println()
	fmt.Printf("%-v\n", tind)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&tind)
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
func verCargar(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tienda)
}

func Prueba(w http.ResponseWriter, r *http.Request) {
	
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Ingrese una tienda valida")
	}
	json.Unmarshal(reqBody, &tienda)
	fmt.Println(tienda)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&tienda)
}

var depli []string
//var grjson []string
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
	vector = make([]*Estructura.Lista, cantindex*cantdep*5)
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
				tienda.Dep=j
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
				pos := ((j*(cantindex)+i)*(5))+k
				vector[pos] = matrix[i][j].List[k]
			}
		}
	}
	departamentos = RemoveDuplicatesFromSlice(departamentos)

	for _, li:= range vector{
		salida := li.Decodificar(w,r)
		if salida !=""{
			depli = append(depli, salida )	
		}	
	}
	
	stjs := strings.Join(depli,",")
	stjs = "{\n\"tiendas\":[\n"+stjs+"]\n}"
	sal := []byte(stjs)
	err2 := json.Unmarshal(sal, &tind)
	if err2 != nil {
		log.Print(err2)
	}
	fmt.Printf("%-v\n", tind)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&tind)
	tamind = len(index)
	
	
	
	

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

