//En este paquete se define la estructura de un producto
package Tiendas

//definir la estructura del producto
type Producto struct{
	Nombre string `json:"Nombre"`
	Codigo int `json:"Codigo"`
	Descripcion string `json:"Descripcion"`
	Precio int `json:"Precio"`
	Cantidad int `json:"Cantidad"`
	Imagen string `json:"Imagen"`
}

func (p *Producto) getCodigo() int{
	return p.Codigo
}

//Definir una estructura inicial para los inventarios
type InvInit struct{
	Inventarios [] Inventario `json:"Inventarios"`
}

//Definir una estructura inicial para los pedidos
type PedInit struct{
	Pedidos [] Pedido `json:"Pedidos"`
}

//Definir la estructura de los inventarios
type Inventario struct{
	Tienda string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int `json:"Calificacion"`
	Productos [] Producto `json:"Productos"`
}

//Definir la estructura de los pedidos
type Pedido struct{
	Fecha string `json:"Fecha"`
	Tienda string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int `json:"Calificacion"`
	Productos [] Codigos `json:"Productos"`
}

//Definir una estructura para los codigos de los productos en los pedidos
type Codigos struct{
	Codigo int `json:"Codigo"`
}

//Definir una funcion para crear un nuevo producto
func NuevoProducto(nomb string,cod int, desc string, precio int, cant int, img string ) *Producto{
	return &Producto{nomb, cod, desc, precio, cant, img}
}

//Definir una estructura para el carrito de compras
type Carrito struct{
	Tienda string `json: "Tienda"`
	Dep int `json: "Dep"`
	Calificacion int `json: "Calificacion"`
	Nombre string `json: "Nombre"`
	Codigo int `json: "Codigo"`
	Cantidad int `json: "Cantidad"`
}