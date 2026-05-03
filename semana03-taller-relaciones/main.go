package main
import(
	"fmt"
	"semana03-taller-relaciones/internal/cafeteria")

	func main() {

		cafeteria.AgregarCategoria(cafeteria.Categoria{ID: 1, Nombre: "Bebidas"})
		cafeteria.AgregarCategoria(cafeteria.Categoria{ID: 2, Nombre: "Snacks"})
		
		var repo cafeteria.Repositorio = cafeteria.NewRepoMemoria()
		repo.GuardarCliente(cafeteria.Cliente{ID: 1, Nombre: "jostin",Carrera: "ti", Saldo: 500.00})
		repo.GuardarCliente(cafeteria.Cliente{ID: 2, Nombre: "JOSEPH",Carrera: "ti", Saldo: 500.00})
		repo.GuardarProducto(cafeteria.Producto{ID: 1, Nombre: "agua",Precio: 0.50,Stock: 50,CategoriaID: 1})
		repo.GuardarProducto(cafeteria.Producto{ID: 2, Nombre: "tostada",Precio: 1.00,Stock: 50,CategoriaID: 2})
		repo.GuardarProducto(cafeteria.Producto{ID: 3, Nombre: "cafe",Precio: 0.70,Stock: 50,CategoriaID: 1})

		fmt.Printf("listando productos:\n")
	for _,p := range repo.ListarProductos(){
		fmt.Printf("%+v\n",p)
	}
	fmt.Printf("listando clientes:\n")
	for _,c := range repo.ListarClientes(){
		fmt.Printf("%+v\n",c)
	}

		c, err := repo.BuscarClientePorID(1)
	if err != nil{
		fmt.Printf("Error al buscar Cliente: %s\n", err.Error())
	}
	fmt.Printf("\nEncontrado: %s\n", c.Nombre)

	// Buscamos uno que NO existe — aquí está el problema que resolveremos
	fantasma,err := repo.BuscarClientePorID(999)
	if err != nil{
		fmt.Printf("Error al buscar Cliente: %s\n", err.Error())
		return
	}
	fmt.Println(fantasma)

	
		
	}
	//* ¿Tuviste que poner cliente producto pedidos en el mismo paquete porque si porque no? 
	//    Si los puse en el mismo paquete porque están relacionados y forman parte del mismo paquete cafeteria*
	// ¿Que problema aparecia si intentara separa producto en un paquete aparte cuando pedido lo tiene anidado?
	// 	me pedidiria llamar a mi paquete de cafeteria para volver usarlo y asi mismo llamarlo al main
	// ¿Comparado con el dia A (donde usamos IDs): ¿Que ventaja tiene el modelo con id para organizar el codigo en paquete?
	// Menos dependencia , y mejorea al momento de organizar paquete y codigo //