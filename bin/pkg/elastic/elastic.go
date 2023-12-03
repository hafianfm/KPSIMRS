package elastic

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

var ESCli *elasticsearch.Client

func InitElastic() (err error) {

	cfg := elasticsearch.Config{
		CloudID: "SIMRS:dXMtY2VudHJhbDEuZ2NwLmNsb3VkLmVzLmlvJDhjMmViMTcwNGEwNjQ2MzdiNDhlYmVmMjI5NjJhZDA5JDEwNjA0ZDhlNzZiNTQzNTFhNzgzMTVmZTVlMjQ4M2U4",
		APIKey:  "aE16VUtvd0JvWGM2ZGdVOUpVRlU6SVlLTDNGZlVSZy1pNGItbjFsZVh2QQ==",
	}

	
	ESCli, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		return
	}

	res, err := ESCli.Ping()

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("ElasticSearch Connected %s", res)
	return
}
