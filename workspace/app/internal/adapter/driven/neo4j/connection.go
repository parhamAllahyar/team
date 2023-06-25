package neo4j

import (
	"fmt"
	"workspace/configs"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func Connection(config configs.NEOConfig) (neo4j.DriverWithContext, error) {
	uri := "bolt://localhost:7687"
	username := "neo4j"
	password := "123456789"

	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		panic(err)
	}
	fmt.Println("jediujdejideijdeijdejidejidejiejijijidejide", driver)
	return driver, nil

}
