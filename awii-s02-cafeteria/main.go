package main

import (
	"errors"
	"fmt"
)

type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Saldo   float64
}
type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}
type Pedido struct {
	ID         int
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
	Fecha      string
}

func ListarClientes(clientes []Cliente) {
	fmt.Println("\n=== CLIENTES REGISTRADOS ===")
	if len(clientes) == 0 {
		fmt.Println("no hay clientes")
		return

	}
	for _, c := range clientes {
		fmt.Printf("[%d] %s | Carrera: %s | Saldo: %.2f\n", c.ID, c.Nombre, c.Carrera, c.Saldo)

	}

}

func BuscarClientePorID(clientes []Cliente, id int) int {
	for i, c := range clientes {
		if c.ID == id {
			return i
		}
	}
	return -1 // convención: -1 significa "no encontrado"
}

func agregarCliente(clientes []Cliente, cliente Cliente) []Cliente {
	return append(clientes, cliente)
}

func EliminarCliente(clientes []Cliente, id int) []Cliente {
	idx := BuscarClientePorID(clientes, id)
	if idx == -1 {
		fmt.Printf("⚠ Cliente con ID %d no existe.\n", id)
		return clientes
	}
	return append(clientes[:idx], clientes[idx+1:]...)
}

func ListarProducto(productos []Producto) {
	fmt.Println("\n=== PRODUCTOS REGISTRADOS ===")
	if len(productos) == 0 {
		fmt.Println("no hay productos")
		return

	}
	for _, c := range productos {
		fmt.Printf("[%d] PRODUCTO: %s | PRECIO: %.2F | STOCK: %.d | CATEGORIA: %s |\n", c.ID, c.Nombre, c.Precio, c.Stock, c.Categoria)

	}

}

func BuscarProductoPorID(productos []Producto, id int) int {
	for i, c := range productos {
		if c.ID == id {
			return i
		}
	}
	return -1 // convención: -1 significa "no encontrado"
}

func agregarProducto(productos []Producto, producto Producto) []Producto {
	return append(productos, producto)
}

func EliminarProducto(productos []Producto, id int) []Producto {
	idx := BuscarProductoPorID(productos, id)
	if idx == -1 {
		fmt.Printf("⚠ Producto con ID %d no existe.\n", id)
		return productos
	}
	return append(productos[:idx], productos[idx+1:]...)
}

func DescontarStock(productos *Producto, cantidad int) error {
	if cantidad <= 0 {
		return errors.New("la cantidad debe ser positiva")
	}
	if productos.Stock < cantidad {
		return fmt.Errorf("cupos insuficientes: hay %d, piden %d",
			productos.Stock, cantidad)
	}
	productos.Stock -= cantidad
	return nil
}

func DescontarDinero(clientes *Cliente, cantidad int) error {
	if cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}
	if int(clientes.Saldo) < cantidad {
		return fmt.Errorf("cupos insuficientes en %s (hay %.2f, solicita %d)",
			clientes.Nombre, clientes.Saldo, cantidad)
	}
	clientes.Saldo -= float64(cantidad)
	return nil
}

func RegistrarPedido(
	clientes []Cliente,
	productos []Producto,
	pedidos []Pedido,
	clienteID int,
	productoID int,
	cantidad int,
	fecha string,
) ([]Pedido, error) {

	// Paso 1: validar que el cliente existe
	idxT := BuscarClientePorID(clientes, clienteID)
	if idxT == -1 {
		return pedidos, errors.New("cliente no encontrado")
	}

	// Paso 2: validar que el producto existe
	idxN := BuscarProductoPorID(productos, productoID)
	if idxN == -1 {
		return pedidos, errors.New("producto no encontrado")
	}

	// Paso 3: Calcular el total
	resultado := productos[idxN].Precio * float64(cantidad)

	// Verificar Stock
	if productos[idxN].Stock < cantidad {
		return pedidos, errors.New("stock insuficiente")
	}

	err := DescontarStock(&productos[idxN], cantidad)
	if err != nil {
		return pedidos, err
	}

	err = DescontarDinero(&clientes[idxT], int(resultado))
	if err != nil {
		productos[idxN].Stock += cantidad
		return pedidos, err
	}

	// Crear pedido con ID autoincremental
	nuevoID := len(pedidos) + 1
	nuevo := Pedido{
		ID:         nuevoID,
		ClienteID:  clienteID,
		ProductoID: productoID,
		Cantidad:   cantidad,
		Total:      resultado,
		Fecha:      fecha,
	}

	//////
	//Agregar el nuevo pedido al slice:
	pedidos = append(pedidos, nuevo)
	return pedidos, nil

}

func ListarPedido(pedidos []Pedido) {
	fmt.Println("\n=== PEDIDOS REGISTRADOS ===")
	if len(pedidos) == 0 {
		fmt.Println("(no hay pedidos)")
		return
	}
	for _, n := range pedidos {
		fmt.Printf("  [%d] ClienteID:%d |ProductoID: %d | Cantidad:%d |Total: $%.2f | Fecha: %s |\n",
			n.ID, n.ClienteID, n.ProductoID, n.Cantidad, n.Total, n.Fecha)
	}
}

func main() {
	var pedidos []Pedido
	var err error

	clientes := []Cliente{
		{1, "Jostin", "Ti", 500.40},
		{2, "Joseph", "Ti", 100.40},
		{3, "Lisbeth", "Derechos", 500.40},
	}
	//agregar
	cliente := Cliente{4, "Carmen", "Odontologia", 250.00}
	clientes = agregarCliente(clientes, cliente)
	clientes = EliminarCliente(clientes, 1)
	ListarClientes(clientes)

	idx := BuscarClientePorID(clientes, 2)
	if idx == -1 {
		fmt.Println("cliente no existe")
		return
	}
	fmt.Println("Encontrado:", clientes[idx].Nombre)

	productos := []Producto{
		{1, "Cafe", 0.75, 100, "Bebidas"},
		{2, "Capuchino", 1.5, 50, "Bebidas"},
		{3, "Tostada", 1.0, 30, "Comestible"},
		{4, "Batidos", 1.50, 20, "Bebidas"},
	}

	producto := Producto{5, "Manzana", 0.25, 50, "Fruta"}
	productos = agregarProducto(productos, producto)

	productos = EliminarProducto(productos, 1)
	ListarProducto(productos)
	idx2 := BuscarProductoPorID(productos, 2)
	if idx2 == -1 {
		fmt.Println("producto no existe")
		return
	}
	fmt.Println("Encontrado:", productos[idx2].Nombre)

	pedidos, err = RegistrarPedido(clientes, productos, pedidos, 2, 3, 10, "20/09/2005")
	if err != nil {
		fmt.Println("Error:", err)
	}
	ListarPedido(pedidos)
	ListarClientes(clientes)
	ListarProducto(productos)

}
