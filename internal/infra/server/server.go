package server

import (
	"finances/internal/infra/settings"
	"finances/pkg/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	routes.Grouper(r)
	routes.Docs(r)

	port := settings.GetEnvs().HTTPServerPort
	protocol := settings.GetEnvs().HTTPServerProtocol
	certificate := settings.GetEnvs().HTTPServerTLSCertificate
	key := settings.GetEnvs().HTTPServerTLSKey

	switch protocol {
	case "http":
		if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
			log.Fatal("Error on server start: ", err)
		}
	case "https":
		if err := r.RunTLS(fmt.Sprintf(":%s", port), certificate, key); err != nil {
			log.Fatal("Error on server start: ", err)
		}
	}
}
