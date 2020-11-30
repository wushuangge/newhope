package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type (
	envConfig struct {
		Dev        bool
		LocalDev   bool
		LogConsole bool

		LogDir        string
		SensorData    string
		SensorProject string

		HttpApi string
		WsApi   string

		//redis config
		LoginRedisAddr     string
		LoginRedisPassword string

		CoinRedisAddr     string
		CoinRedisPassword string

		MarryRedisAddr     string
		MarryRedisPassword string

		//mysql config
		MainMysqlDsn string
		LogMysqlDsn  string
		CoinMysqlDsn string
		//supervise

		TcpGoApi          string
		FixRoomApi        string
		ProxyIp           string
		ProxyPort         string
		DirtyFilterServer string

		//nsq
		NsqAddr string

		//mongodb
		MongoUserName string
		MongoPassword string
		MongoIp       string
		MongoPort     string

		//ListenAddr
		ListenAddr string

		//EnableHttps
		EnableHttps bool

		//certificate
		CertPem string
		KeyPem  string

		//authUrl
		AuthUrl  string
		UsersUrl string
	}
)

var (
	EnvConfigFile = "config/env.toml"

	envConf envConfig
)

func InitEnvConf() {
	if _, err := toml.DecodeFile(EnvConfigFile, &envConf); err != nil {
		log.Fatal(err)
	}
}

func IsDev() bool {
	return envConf.Dev
}

func IsLocalDev() bool {
	return envConf.LocalDev
}

func GetLoginRedisAddr() string {
	return envConf.LoginRedisAddr
}

func GetLoginRedisPassword() string {
	return envConf.LoginRedisPassword
}

func GetMainMysqlDsn() string {
	return envConf.MainMysqlDsn
}

func GetLogDir() string {
	return envConf.LogDir
}

func GetNsqAddr() string {
	return envConf.NsqAddr
}

func GetMongoUserName() string {
	return envConf.MongoUserName
}

func GetMongoPassword() string {
	return envConf.MongoPassword
}

func GetMongoIp() string {
	return envConf.MongoIp
}

func GetMongoPort() string {
	return envConf.MongoPort
}

func GetListenAddr() string {
	return envConf.ListenAddr
}

func GetEnableHttps() bool {
	return envConf.EnableHttps
}

func GetCertPem() string {
	return envConf.CertPem
}

func GetKeyPem() string {
	return envConf.KeyPem
}

func GetAuthUrl() string {
	return envConf.AuthUrl
}

func GetUsersUrl() string {
	return envConf.UsersUrl
}
