package conf

import (
	"gin_mall/dao"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppMode  string
	HttpPort string

	Db          string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassWord  string
	DbName      string
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
	ValidEmail  string
	SmtpHost    string
	SmtpEmail   string
	SmtpPass    string
	Host        string
	ProductPath string
	AvatarPath  string
)

func Init() {
	//本地读取环境变量
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	LoadMysql(file)
	LoadRedis(file)
	LoadEmail(file)
	LoadPhotoPath(file)
	//mysql 读(8)主
	pathRead := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	//mysql 写（2）从 主从复制
	pathWrite := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	dao.Database(pathRead, pathWrite)
}

func LoadPhotoPath(file *ini.File) {
	Host = file.Section("service").Key("Host").String()
	ProductPath = file.Section("service").Key("ProductPath").String()
	AvatarPath = file.Section("service").Key("AvatarPath").String()

}

func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("service").Key("ValidEmail").String()
	SmtpHost = file.Section("service").Key("SmtpHost").String()
	SmtpEmail = file.Section("service").Key("SmtpEmail").String()
	SmtpPass = file.Section("service").Key("SmtpPass").String()
}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("service").Key("RedisDb").String()
	RedisAddr = file.Section("service").Key("RedisAddr").String()
	RedisPw = file.Section("service").Key("RedisPw").String()
	RedisDbName = file.Section("service").Key("RedisDbName").String()
}

func LoadMysql(file *ini.File) {
	Db = file.Section("service").Key("Db").String()
	DbHost = file.Section("service").Key("DbHost").String()
	DbPort = file.Section("service").Key("DbPort").String()
	DbUser = file.Section("service").Key("DbUser").String()
	DbPassWord = file.Section("service").Key("DbPassWord").String()
	DbName = file.Section("service").Key("DbName").String()
}
func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}
