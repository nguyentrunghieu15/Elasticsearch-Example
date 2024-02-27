package configs

type ElasticConf struct {
	API_KEY    string `envconfig:"optional,ELASTICSEARCH_API_KEY"`
	Username   string `envconfig:"ELASTICSEARCH_USERNAME"`
	Password   string `envconfig:"ELASTICSEARCH_PASSWORD"`
	CertCAPath string `envconfig:"ELASTIC_CERT_PATH"`
	Addresses  string `envconfig:"ELASTIC_ADDRESS"`
	CloudID    string `envconfig:"optional,ELASTIC_CLOUD_ID"`
}
