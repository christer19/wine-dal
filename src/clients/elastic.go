package clients

import (
	"fmt"

	elastic "github.com/olivere/elastic/v7"
)

// ElasticClient exposes the pointer to the client to be used in queries
var ElasticClient *elastic.Client

// InitElastic sets a connection with the elastic cluster
func InitElastic() error {
	elasticServer := "http://localhost:9200"
	fmt.Println("connecting to elasticserver " + elasticServer)

	var err error
	ElasticClient, err = elastic.NewClient(
		elastic.SetURL(elasticServer),
		// autodetect new elastic nodes
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	if err != nil {
		return err
	}

	fmt.Println("connected to elasticsearch cluster")
	return nil
}
