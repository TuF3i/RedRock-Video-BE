package config_template

type DanmuGatewayConfig struct {
	PodUID string
	Hertz  HertzForDanmuGateway
	Etcd   EtcdForDanmuGateway
}

type HertzForDanmuGateway struct {
	ListenAddr     string
	ListenPort     string
	MonitoringPort string
}

type EtcdForDanmuGateway struct {
	ServiceName string
	Namespace   string
	Urls        []string
}
