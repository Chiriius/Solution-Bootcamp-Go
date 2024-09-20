package main

import (
	"bootcamp_api/api/server"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		logrus.Panic("Error al obtener el directorio de trabajo:", err)
	}
	fmt.Println("Directorio de trabajo actual:", dir)
	entries, err := os.ReadDir("./")
	if err != nil {
		logrus.Fatal(err)
	}

	for _, e := range entries {
		logrus.Info(e.Name())
	}

	//envPath := "/home/miguel-angel-sena/Documents/golang/Solution-Bootcamp-Go/.env"
	//envPath := filepath.Join("..", ".env")
	envPath := filepath.Join(dir, ".env")
	logrus.Infof("Buscando archivo .env en: %s", envPath)

	viper.SetConfigFile(envPath)

	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.JSONFormatter{})

	if err := viper.ReadInConfig(); err != nil {
		logger.Panic("Error al leer el archivo de configuración:", err)
	}

	portHttp := viper.GetString("SERVER_PORT_HTTP")
	portGrpc := viper.GetString("SERVER_PORT_GRPC")
	dbUrl := viper.GetString("DB_URL")
	logrus.Info("ESTA ES LA DB URL", dbUrl)

	server, err := server.New(logger, portHttp, portGrpc, dbUrl)
	if err != nil {
		logger.Panic("Failed to create server:", err)
	}
	defer server.Close()

	if err := server.Start(); err != nil {
		logger.Error(err)
	}
}
