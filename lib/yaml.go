package lib

import (
	"github.com/yancyzhou/unionsdk/JdunionSdk"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const (
	CONFPATH string = "DockerBuild/configuration/conf.yaml"
)

//start
type Yaml struct {
	ConfigKey    string       `yaml:"ConfigKey"`
	DBConf       MongoDB      `yaml:"mongodb"`
	RedisConn    string       `yaml:"redisConn"`
	JwtConf      JwtConf      `yaml:"JwtConf"`
	WeChat       WeChats      `yaml:"wechat"`
	Server       Server       `yaml:"Server"`
	AliAccessKey AliAccessKey `yaml:"AliAccessKey"`
	AliPayConf   AliPayConf   `yaml:"AliPayConf"`
	WxPayConf    WxPayConf    `yaml:"WxPayConf"`
	JDConfig     JDConfig     `yaml:"JDConfig"`
}

type AliPayConf struct {
	APPID        string `yaml:"APPID"`
	ALIPUBLICKEY string `yaml:"ALIPUBLICKEY"`
	PRIVATEKEY   string `yaml:"PRIVATEKEY"`
	IsPruduction bool   `yaml:"IsPruduction"`
}

type Server struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}
type JDConfig struct {
	AppSecret string `yaml:"appSecret"`
	AppKey    string `yaml:"appKey"`
}
type MongoDB struct {
	User         string `yaml:"db_user"`
	Host         string `yaml:"db_host"`
	Password     string `yaml:"db_pass"`
	Port         string `yaml:"db_port"`
	DatabaseName string `yaml:"db_database_name"`
	AuthDBName   string `yaml:"db_auth_name"`
	Uri          string `yaml:"url"`
}

type JwtConf struct {
	Issuer    string `yaml:"issuer"`
	Exptime   int64  `yaml:"exptime"`
	Notbefore int64  `yaml:"notbefore"`
}

type WeChats struct {
	APPID     string `yaml:"APPID"`
	APPSECRET string `yaml:"APPSECRET"`
}

type WxPayConf struct {
	APPID     string `yaml:"APPID"`
	APPSECRET string `yaml:"APPSECRET"`
	MCHID     string `yaml:"MCHID"`
	APPKEY    string `yaml:"APPKEY"`
}

type AliAccessKey struct {
	AccessKeyID     string `yaml:"AccessKeyID"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
	AliYunDomain    string `yaml:"AliYunDomain"`
	SmsVersion      string `yaml:"SmsVersion"`
	SmsTemplateCode string `yaml:"SmsTemplateCode"`
	SmsSignName     string `yaml:"SmsSignName"`
	SmsRegionId     string `yaml:"SmsRegionId"`
}

var ServerConf *Yaml
var J JdunionSdk.JdSdk
func init() {
	yamlFile, err := ioutil.ReadFile(CONFPATH)
	if err != nil {
		log.Fatal(err)
	} else {
		err = yaml.Unmarshal(yamlFile, &ServerConf)
		if err != nil {
			log.Fatal(err)
		}
	}

	J.NewContext(ServerConf.JDConfig.AppKey, ServerConf.JDConfig.AppSecret)
}
