package configs

type ElasticConf struct {
	API_KEY string `envconfig:"ELASTICSEARCH_API_KEY"`
}
