package pekka

import (
	"fmt"
	"net"
	"net/http"

	"github.com/jinzhu/gorm"
)

type PekkaEndpoint struct {
	conn *net.Listener
	db   *gorm.DB
}

type httpserve struct {
}

func (h httpserve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request", r.URL.Query()["pin"], r.URL.Query()["state"], r.RemoteAddr)
}

func (pe *PekkaEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request", r.URL.Query()["pin"], r.URL.Query()["state"], r.RemoteAddr)

}

func CreatePekkaEndpoint(port string, db *gorm.DB) *PekkaEndpoint {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		fmt.Println("Pekka endpoint create failed")
	}

	pekkaEndpoint := &PekkaEndpoint{
		conn: &l,
		db:   db,
	}

	http.Serve(l, pekkaEndpoint)

	return pekkaEndpoint
}
