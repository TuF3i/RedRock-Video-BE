package config

var (
	ETCD_ADDRS          = []string{"47.108.184.229:2181"}
	REDIS_CLUSTER_ADDRS = []string{"47.108.184.229:6379", "47.108.184.229:6380", "47.108.184.229:6381"}
	KAFKA_CLUSTER_ADDRS = []string{"47.108.184.229:9092"}
	PGSQL_ADDRS         = []string{"47.108.184.229"}
	MINIO_ADDRS         = []string{"101.36.123.131:9000"}
)
