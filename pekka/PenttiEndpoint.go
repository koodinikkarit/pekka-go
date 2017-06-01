package pekka

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	pekka "github.com/koodinikkarit/pekka/db"
)

type PenttiEndpoint struct {
	conn *net.Listener
	db   *gorm.DB
}

func (pe *PenttiEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pin := r.URL.Query()["pin"][0]
	//state := r.URL.Query()["state"][0]
	sourceAddr := strings.Split(r.RemoteAddr, ":")[0]

	//fmt.Println(pin, state, sourceAddr)

	pinNumber, _ := strconv.Atoi(pin)
	//buttonState, _ := strconv.Atoi(state)
	var pentti pekka.Pentti
	var button pekka.Button
	pe.db.Where(&pekka.Pentti{Ip: sourceAddr}).FirstOrCreate(&pentti)
	pe.db.Where(&pekka.Button{PenttiID: pentti.ID, ButtonNumber: pinNumber}).FirstOrCreate(&button)

	fmt.Println("pentti", pentti)
}

func CreatePenttiEndpoint(port string, db *gorm.DB) *PenttiEndpoint {
	l, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Println("Pentti endpoint create failed", port)
	}

	penttiEndpoint := &PenttiEndpoint{
		conn: &l,
		db:   db,
	}

	http.Serve(l, penttiEndpoint)

	return penttiEndpoint
}
