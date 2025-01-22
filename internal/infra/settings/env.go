package settings

import (
	"github.com/caarlos0/env/v7"
	"log"
)

var envs Environment

type Environment struct {
	DatabaseEnvironment
	HTTPServerEnvironment
}
type DatabaseEnvironment struct {
	DatabaseHost        string `env:"DATABASE_HOST,required" envDefault:"localhost"`
	DatabaseUser        string `env:"DATABASE_USER,required"  envDefault:"user"`
	DatabasePassword    string `env:"DATABASE_PASSWORD,required"  envDefault:"password"`
	DatabaseName        string `env:"DATABASE_NAME,required"  envDefault:"finances"`
	DatabasePort        string `env:"DATABASE_PORT,required"  envDefault:"5433"`
	DatabaseSSL         string `env:"DATABASE_SSL,required"  envDefault:"disable"`
	DatabaseDebug       string `env:"DATABASE_DEBUG,required"  envDefault:"false"`
	DatabaseAutoMigrate string `env:"DATABASE_AUTOMIGRATE,required"  envDefault:"true"`
}

type HTTPServerEnvironment struct {
	HTTPServerPort           string `env:"HTTP_SERVER_PORT,required"  envDefault:"8080"`
	HTTPServerProtocol       string `env:"HTTP_SERVER_PROTOCOL,required"  envDefault:"http"`
	HTTPServerTLSCertificate string `env:"HTTP_SERVER_TLS_CERTIFICATE,required" envDefault:"/cert.pem"`
	HTTPServerTLSKey         string `env:"HTTP_SERVER_TLS_KEY,required"  envDefault:"/cert.key"`
}

func GetEnvs() Environment {
	return envs
}

func LoadEnvs() {
	if err := env.Parse(&envs); err != nil {
		log.Fatalln("Error loading environment variables" + err.Error())
	}
	log.Println("Loaded environment variables!")
}
