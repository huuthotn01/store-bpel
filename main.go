package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"os"
	"strings"
)

func main() {
	serviceName := os.Args[1]
	serviceIdentifier := getServiceIdentifier(serviceName)

	// create service folder
	err := os.Mkdir(serviceName, 0755)
	if err != nil {
		panic(err)
	}

	// create config.txt file
	err = os.WriteFile(serviceName+"/config.txt", []byte(fmt.Sprintf(configTxtFileTemplate, serviceName)), 0755)
	if err != nil {
		panic(err)
	}

	// create Makefile
	err = os.WriteFile(serviceName+"/Makefile", []byte(fmt.Sprintf(makeFileTemplate, serviceName, serviceName)), 0755)
	if err != nil {
		panic(err)
	}

	// create service sub folders
	folderLists := []string{"adapter", "config", "controller", "main", "migration", "repository", "schema"}
	for _, folder := range folderLists {
		err := os.Mkdir(serviceName+"/"+folder, 0755)
		if err != nil {
			panic(err)
		}
	}

	// create config.go file in config folder
	err = os.WriteFile(serviceName+"/config/config.go", []byte(loadConfigFileTemplate), 0755)
	if err != nil {
		panic(err)
	}

	// create kafka_adapter.go file in adapter folder
	err = os.WriteFile(serviceName+"/adapter/kafka_adapter.go", []byte(kafkaAdapterFileTemplate), 0755)
	if err != nil {
		panic(err)
	}

	var (
		serviceNameCamelCase  = strcase.ToCamel(serviceName)
		serviceNameLowerCamel = strcase.ToLowerCamel(serviceName)
	)

	// create controller.go file in controller folder
	err = os.WriteFile(serviceName+"/controller/controller.go",
		[]byte(fmt.Sprintf(controllerFileTemplate, serviceName, serviceName, serviceName,
			serviceName, serviceNameCamelCase, serviceNameLowerCamel, serviceNameCamelCase,
			serviceNameCamelCase, serviceNameLowerCamel)),
		0755)
	if err != nil {
		panic(err)
	}

	// create database.go in main folder
	err = os.WriteFile(serviceName+"/main/database.go", []byte(fmt.Sprintf(databaseFileTemplate, serviceName)), 0755)
	if err != nil {
		panic(err)
	}

	// create main.go in main folder
	err = os.WriteFile(serviceName+"/main/main.go", []byte(fmt.Sprintf(mainFileTemplate, serviceName, serviceName,
		serviceNameCamelCase, strcase.ToCamel(serviceIdentifier), "%d", strcase.ToCamel(serviceIdentifier), "%d")), 0755)
	if err != nil {
		panic(err)
	}

	// create model.go in repository folder
	err = os.WriteFile(serviceName+"/repository/model.go", []byte(fmt.Sprintf(modelFileTemplate, serviceNameLowerCamel)), 0755)
	if err != nil {
		panic(err)
	}

	// create repository.go in repository folder
	err = os.WriteFile(serviceName+"/repository/repository.go", []byte(fmt.Sprintf(repositoryFileTemplate, serviceNameCamelCase,
		serviceNameCamelCase, serviceNameLowerCamel)), 0755)
	if err != nil {
		panic(err)
	}

	// create response.go in schema folder
	err = os.WriteFile(serviceName+"/schema/response.go", []byte(responseSchemaFileTemplate), 0755)
	if err != nil {
		panic(err)
	}
}

func getServiceIdentifier(serviceName string) string {
	return strings.Split(serviceName, "_")[0]
}

const configTxtFileTemplate = `HTTP_PORT: 14000

MYSQL:
  HOST: localhost
  PORT: 3306
  USERNAME: admin
  PASSWORD: 'admin'
  DATABASE: %s
`

const makeFileTemplate = `start:
	cd main && go run .

migrate:
	migrate create -ext sql -dir migration -seq $(name)

migrate-up:
	migrate -path migration -database "mysql://admin:admin@tcp(localhost:3306)/%s" -verbose up

migrate-down:
	migrate -path migration -database "mysql://admin:admin@tcp(localhost:3306)/%s" -verbose down
`

const loadConfigFileTemplate = `package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	HttpPort int ` + "`json:\"http_port\" mapstructure:\"http_port\"`" + `
	MySQL *MySQLConfig ` + "`json:\"mysql\" mapstructure:\"mysql\"`" + `

	StaffServicePort int ` + "`json:\"staff_service_port\" mapstructure:\"staff_service_port\"`" + `
}

type MySQLConfig struct {
	Host string ` + "`json:\"host\" mapstructure:\"host\"`" + `
	Port int ` + "`json:\"port\" mapstructure:\"port\"`" + `
	Username string ` + "`json:\"username\" mapstructure:\"username\"`" + `
	Password string ` + "`json:\"password\" mapstructure:\"password\"`" + `
	Database string ` + "`json:\"database\" mapstructure:\"database\"`" + `
}

func Load() (config *Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}
`

const controllerFileTemplate = `package controller

import (
	"gorm.io/gorm"
	"store-bpel/%s/adapter"
	"store-bpel/%s/config"
	repo "store-bpel/%s/repository"
	"store-bpel/%s/schema"
)

type I%sController interface {
	
}

type %sController struct{
	cfg *config.Config
	repository repo.I%sRepository

	kafkaAdapter adapter.IKafkaAdapter
}

func NewController(cfg *config.Config, db *gorm.DB) I%sController {
	// init repository
	repository := repo.NewRepository(db)

	// init kafka adapter
	kafkaAdapter := adapter.NewKafkaAdapter()

	return &%sController{
		cfg: cfg,
		repository: repository,
		kafkaAdapter: kafkaAdapter,
	}
}
`

const kafkaAdapterFileTemplate = `package adapter

type IKafkaAdapter interface {

}

type kafkaAdapter struct {

}

func NewKafkaAdapter() IKafkaAdapter {
	return &kafkaAdapter{}
}
`

const databaseFileTemplate = `package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"store-bpel/%s/config"
)

func dsn(dbConfig *config.MySQLConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
}

func DbConnect(dbConfig *config.MySQLConfig) (*gorm.DB, error) {
	dbDsn := dsn(dbConfig)
	db, err := gorm.Open(mysql.Open(dbDsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// test db connection
	err = db.Raw("SELECT 1").Error
	if err != nil {
		return nil, err
	}

	return db, nil
}
`

const mainFileTemplate = `package main

import (
	"github.com/spf13/cast"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"store-bpel/%s/config"
	"store-bpel/%s/controller"
)

var ctrl controller.I%sController

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("%s Service server started at port %s", cfg.HttpPort)

	db, err := DbConnect(cfg.MySQL)
	if err != nil {
		panic(err)
	}

	ctrl = controller.NewController(cfg, db)

	r := mux.NewRouter()
	registerEndpoint(r)

	if err = http.ListenAndServe(":" + cast.ToString(cfg.HttpPort), r); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s Service initialized successfully at port %s", cfg.HttpPort)
}

func registerEndpoint(r *mux.Router) {
	// r.HandleFunc({api}, {handleFunc})
}
`

const modelFileTemplate = `package repository

import (
	"gorm.io/gorm"
)

type %sRepository struct {
	db *gorm.DB
}
`

const repositoryFileTemplate = `package repository

import (
	"gorm.io/gorm"
)

type I%sRepository interface {
	
}

func NewRepository(db *gorm.DB) I%sRepository {
	return &%sRepository{
		db: db,
	}
}
`

const responseSchemaFileTemplate = `package schema

type UpdateResponse struct {
	StatusCode int
	Message string
}
`
