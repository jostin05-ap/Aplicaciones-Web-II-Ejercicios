package cafeteria

import ("errors")

var(
	ErrClienteNoEncontrado = errors.New("cliente no encontrado")
	ErrProductoNoEncontrado = errors.New("producto no encontrado")
	ErrStockInsuficiente= errors.New("stock insuficiente")
	ErrSaldoInsuficiente = errors.New("saldo insuficiente del cliente")
	ErrCategoriaNoEncontrada =errors.New("Categoria no Encontrada")
)


type Cliente struct{
	ID int
	Nombre string
	Carrera string
	Saldo float64
	
}

type Producto struct {
	ID          int
	Nombre      string
	Precio      float64
	Stock       int
	CategoriaID int // referencia a Categoria por ID, NO anidación
}

type Pedido struct {
	ID int 
	Cliente	Cliente
	Producto Producto
	Cantidad int 
	Total float64
}

type RepoMemoria struct{
	clientes []Cliente
	productos []Producto
	//pedidos []Pedidos
}
type Repositorio interface {
	GuardarProducto(p Producto) error
	BuscarProductoPorID(id int) (Producto, error)
	ListarProductos() []Producto
	GuardarCliente(c Cliente) error
	BuscarClientePorID(id int) (Cliente, error)
	ListarClientes() []Cliente
	

}
type Categoria struct {
	ID     int
	Nombre string
}


var categorias = []Categoria{}
//var productos = []Producto{}


func AgregarCategoria(c Categoria) {
	categorias = append(categorias, c)
}

func BuscarCategoriaPorID(id int) (Categoria, error) {
	
	for _, c := range categorias {
		if c.ID == id {
			return c, nil 
		}
	}
	return Categoria{}, ErrCategoriaNoEncontrada
}



func (r*RepoMemoria) GuardarProducto(p Producto )error{
	r.productos = append(r.productos, p)
	return nil
}
func (r*RepoMemoria) BuscarProductoPorID(id int) (Producto, error) {
	// Mismo problema que BuscarCategoriaPorID.
	for _, p := range r.productos {
		if p.ID == id {
			return p , nil
		}
	}
	return Producto{}, ErrProductoNoEncontrado
}


func (r*RepoMemoria) ListarProductos()[]Producto {
	return r.productos
}


func (r*RepoMemoria) GuardarCliente(c Cliente )error{
	r.clientes = append(r.clientes, c)
	return nil
}

func (r*RepoMemoria) BuscarClientePorID(id int) (Cliente, error) {
	// Mismo problema que BuscarCategoriaPorID.
	for _, c := range r.clientes {
		if c.ID == id {
			return c , nil
		}
	}
	return Cliente{}, ErrClienteNoEncontrado
}


func (r*RepoMemoria) ListarClientes()[]Cliente {
	return r.clientes
}

func NewRepoMemoria() *RepoMemoria{
	return &RepoMemoria{}
}

var _ Repositorio =(*RepoMemoria)(nil)