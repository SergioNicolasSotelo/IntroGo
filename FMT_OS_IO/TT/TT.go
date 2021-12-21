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
