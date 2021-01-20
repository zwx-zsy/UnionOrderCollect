package lib

import (
	"github.com/yancyzhou/JdunionSdk"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const (
	CONFPATH string = "/Users/Vincent/workspace/UnionOrderCollect/DockerBuild/configuration/conf.yaml"
)

//start
type Yaml struct {
	DBConf        MongoDB      `yaml:"Mongodb"`
	RedisConf     Redis        `yaml:"Redis"`
	JwtConf       JwtConf      `yaml:"JwtConf"`
	WeChat        WeChats      `yaml:"Wechat"`
	Server        Server       `yaml:"Server"`
	AliAccessKey  AliAccessKey `yaml:"AliAccessKey"`
	AliPayConf    AliPayConf   `yaml:"AliPayConf"`
	WXPayConf     WXPayConf    `yaml:"WXPayConf"`
	AliOssConf    AliOssConf   `yaml:"AliOss"`
	TimeLayoutStr string       `yaml:"TimeLayoutStr"`
	JDConfig      JDConfig     `yaml:"JDConfig"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
type MongoDB struct {
	User         string `yaml:"user"`
	Host         string `yaml:"host"`
	Password     string `yaml:"passWord"`
	Port         string `yaml:"port"`
	DatabaseName string `yaml:"databaseName"`
	AuthDBName   string `yaml:"authName"`
	Uri          string `yaml:"url"`
}
type Redis struct {
	Host                       string `yaml:"host"`
	PassWord                   string `yaml:"passWord"`
	Port                       string `yaml:"port"`
	DatabaseName               int    `yaml:"db"`
	VerificationRedisKeyPrefix string `json:"verificationRedisKeyPrefix"`
}

type JwtConf struct {
	IsSuer    string `yaml:"issuer"`
	ExpTime   int64  `yaml:"expTime"`
	NotBefore int64  `yaml:"notBefore"`
}

type JDConfig struct {
	AppSecret string `yaml:"appSecret"`
	AppKey    string `yaml:"appKey"`
}

//微信资源配置
type WeChats struct {
	AppID     string `yaml:"appID"`
	AppSecret string `yaml:"appSecret"`
}

type WXPayConf struct {
	AppID     string `yaml:"appID"`
	MchID     string `yaml:"mchID"`
	WXPApiKey string `yaml:"wxPApiKey"`
}

//阿里资源配置
type AliAccessKey struct {
	AccessKeyID     string `yaml:"accessKeyID"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	AliYunDomain    string `yaml:"aliYunDomain"`
	SmsVersion      string `yaml:"smsVersion"`
	SmsTemplateCode string `yaml:"smsTemplateCode"`
	SmsSignName     string `yaml:"smsSignName"`
	SmsRegionId     string `yaml:"smsRegionId"`
}

type AliPayConf struct {
	AppID        string `yaml:"appID"`
	AliPublicKey string `yaml:"aliPublicKey"`
	PrivateKey   string `yaml:"privateKey"`
	IsProduction bool   `yaml:"isProduction"`
}

type AliOssConf struct {
	AccessKeyID     string `yaml:"ossAccessKeyID"`
	AccessKeySecret string `yaml:"ossAccessKeySecret"`
	Bucket          string `yaml:"ossBucket"`
	EndPoint        string `yaml:"ossEndPoint"`
	CallBackHost    string `yaml:"ossCallBackHost"`
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
