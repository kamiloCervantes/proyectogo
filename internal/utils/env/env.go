package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DbUser          string
	DbPassword      string
	DbServer        string
	DbPort          string
	DbName          string
	SecretJwt       string
	SslMode         string
	ServerPort      string
	DocPath         string
	DebugMode       bool
	Mode            string
	ServerPortQuery string
	ServerPortWrite string
}

func NewEnv() *Env {

	return &Env{
		// inicializar otros campos
	}
}

// NewConfig crea una nueva instancia de Config con valores cargados de las variables de entorno
func (c *Env) Env() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	c.DbUser = os.Getenv("DB_USER")
	c.DbPassword = os.Getenv("DB_PASSWORD")
	c.DbServer = os.Getenv("DB_SERVER")
	c.DbPort = os.Getenv("DB_PORT")
	c.DbName = os.Getenv("DB_NAME")
	c.SecretJwt = os.Getenv("SECRET_JWT")
	c.SslMode = os.Getenv("SSLMODE")
	c.ServerPort = os.Getenv("SERVER_PORT")
	c.DocPath = os.Getenv("DOC_PATH")
	c.Mode = os.Getenv("MODE")
	c.DebugMode = os.Getenv("DEBUG_MODE") == "true"
	c.ServerPortQuery = os.Getenv("SERVER_PORT_QUERY")
	c.ServerPortWrite = os.Getenv("SERVER_PORT_WRITE")

}
