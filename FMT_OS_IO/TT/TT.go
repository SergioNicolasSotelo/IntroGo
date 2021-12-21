package main

func main() {

}

type usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

func (u *usuario) cambiarNombre(nuevoNombre string) {
	u.Nombre = nuevoNombre
}
func (u *usuario) cambiarApellido(nuevoApellido string) {
	u.Apellido = nuevoApellido
}
func (u *usuario) cambiarCorreo(nuevoCorreo string) {
	u.Correo = nuevoCorreo
}
func (u *usuario) cambiarContrasena(nuevoContrasena string) {
	u.Contrasena = nuevoContrasena
}

type user struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Product
}
type Product struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func nuevoProducto(nombre string, precio float64) Product {
	nuevoProducto := Product{Nombre: nombre, Precio: precio}
	return nuevoProducto
}
func agregarProducto(usuario user, producto *Product, cantidad int) {
	producto.Cantidad = cantidad
	usuario.Productos = append(usuario.Productos, *producto)
}
func borrarProductos(usuario user) {
	usuario.Productos = nil
}
