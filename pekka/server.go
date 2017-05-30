package pekka

import (
	"fmt"

	"google.golang.org/grpc"

	"github.com/jinzhu/gorm"
	pekka "github.com/koodinikkarit/pekka/db"
)

type PekkaServer struct {
	db            *gorm.DB
	server        *grpc.Server
	pekkaEndpoint *PekkaEndpoint
}

func (p *PekkaServer) Start() {
	fmt.Println("Kaynnistetty")
}

func CreatePekkaServer(
	dbUser string,
	dbPass string,
	dbIp string,
	dbPort string,
	dbName string,
	servicePort string,
	pekkaPort string,
) *PekkaServer {
	db := pekka.CreateDb(dbUser, dbPass, dbIp, dbPort, dbName)
	pekkaEndpoint := CreatePekkaEndpoint(pekkaPort, db)
	server := CreateService(db, servicePort)

	return &PekkaServer{
		db:            db,
		server:        server,
		pekkaEndpoint: pekkaEndpoint,
	}
}
