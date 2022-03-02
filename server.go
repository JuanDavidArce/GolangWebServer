package main // Conoce todo lo de este paquete
import "net/http"

// Puerto del servidor
type Server struct {
	port string
}

// Instanciar servidor y escuchar conexciones
func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Listen() error {
	// Seguno parametro es un handler
	err := http.ListenAndServe(s.port, nil)

	return err
}
