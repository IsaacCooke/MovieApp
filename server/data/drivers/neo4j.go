package drivers

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"io"
	"log"
	"os"
	"strings"
)

type Neo4jConfiguration struct {
	Url      string
	Username string
	Password string
	Database string
}

func (nc *Neo4jConfiguration) NewDriver() (neo4j.DriverWithContext, error) {
	return neo4j.NewDriverWithContext(nc.Url, neo4j.BasicAuth(nc.Username, nc.Password, ""))
}

func ParseConfiguration() *Neo4jConfiguration {
	database := lookupEnvOrGetDefault("NEO4J_DATABASE", "neo4j")
	if !strings.HasPrefix(lookupEnvOrGetDefault("NEO4J_VERSION", "5"), "3") {
		database = ""
	}
	return &Neo4jConfiguration{
		Url:      lookupEnvOrGetDefault("NEO4J_URI", "neo4j://localhost:7687"),
		Username: lookupEnvOrGetDefault("NEO4J_USER", "neo4j"),
		Password: lookupEnvOrGetDefault("NEO4J_PASSWORD", "password"),
		Database: database,
	}
}

func lookupEnvOrGetDefault(key string, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}

func UnsafeClose(closable io.Closer) {
	if err := closable.Close(); err != nil {
		log.Fatal(fmt.Errorf("could not close resource: %w", err))
	}
}

func ToStringSlice(slice []interface{}) []string {
	var result []string
	for _, e := range slice {
		result = append(result, e.(string))
	}
	return result
}
