//整体配置

package sys

import (
	"../utility"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
	"io"
	"net/http"
	"os"
	"time"
	"strings"
	"github.com/labstack/echo/middleware"
)

/*

Debug: 模式

Logger : 日志输出格式的配置，日志等级，输出位置

Startup Banner: 开启后的横幅信息

listen : 监听地址与相关配置
ConfigInfo := &http.Server{
	Addr:         "0.0.0.0:61234",
	ReadTimeout:  20 * time.Minute,
	WriteTimeout: 20 * time.Minute,
}

body limit: 请求的最大字节数（2M）

Disable HTTP/2 : HTTP/2 是否支持

CORS : 跨域访问控制

CSRF : 跨站请求伪造

Cube : labstack 提供的组建，https://labstack.com/signup 注册配置到Cube中，即可使用 /cule 可以查看收集的程序统计日志

Gzip : http response 数据压缩等级配置

*/

type ConfigMgr struct {
	Data ConfigInfo
}

type CORSConfig struct {
	AllowOrigins []string `yaml:"allowOrigins,flow"`
	AllowMethods []string `yaml:"allowMethods,flow"`
}

type BaseConfig struct {
	Debug        bool   `yaml:"debug"`
	LogLevel     string `yaml:"logLevel"`
	HideBanner   bool   `yaml:"hideBanner"`
	DisableHTTP2 bool   `yaml:"disableHTTP2"`
	BodyLimit    string `yaml:"bodyLimit"`
	CORS CORSConfig `yaml:"CORS"`
	CSRFTokenLookup string `yaml:"CSRFTokenLookup"`
	GzipLevel    int   `yaml:"gzipLevel"`
}

type ServerConfig struct {
	ListenAddr   string `yaml:"listenAddr"`
	ReadTimeout  int    `yaml:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout"`
}

type DbConfig struct {
	Driver string `yaml:"driver"`
	Addr   string `yaml:"addr"`
	Name   string `yaml:"name"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
}

type ConfigInfo struct {
	Base   BaseConfig   `yaml:"base"`
	Server ServerConfig `yaml:"server"`
	DB     DbConfig     `yaml:"db"`
}

var ConfigMgrInstance *ConfigMgr = new(ConfigMgr)

var configFileName = "config.yaml"

func (ptr *ConfigMgr) init() (*ConfigInfo, error) {
	fileDir, err := utility.CurrentProcessInstance.GetCurrentDir()

	if err != nil {
		fmt.Printf("ConfigInfo init failed, msg:%s, file:%s", err, fileDir)
		return nil, err
	}

	configFilePath := fileDir + "/" + configFileName

	file, err := os.Open(configFilePath)

	if err != nil {
		fmt.Printf("ConfigInfo init failed, file:%s, msg:%s", configFilePath, err)
		return nil, err
	}

	fileInfo, _ := file.Stat()

	buf := make([]byte, fileInfo.Size())
	readSize, err := file.Read(buf)

	//fmt.Printf("ConfigInfo string:\r\n%s \r\n", string(buf))

	if err != nil && err != io.EOF {
		fmt.Printf("ConfigInfo init failed, read failed, read size:%d, msg:%s \r\n", readSize, err)
		return nil, err
	}

	configInfo := ConfigInfo{}
	err = yaml.Unmarshal(buf, &configInfo)

	if err != nil {
		fmt.Printf("ConfigInfo load failed, Unmarshal failed, msg:%v \r\n", err)
		return nil, err
	}

	//fmt.Printf("ConfigInfo value:\r\n%v \r\n", configInfo)

	return &configInfo, nil
}

func (ptr *ConfigMgr) Load() error {
	info, err := ptr.init()
	if err != nil {
		fmt.Printf("config load failed, msg:%v \r\n", err)
		return errors.New("config load failed")
	}

	ptr.Data = *info

	return nil
}

func (ptr *ConfigMgr) ReLoad() {
	//reload 只能重载部分字段，因为在系统运行的时候，有些配置重载不会起任何作用
}

func (ptr *ConfigMgr) Build(app *echo.Echo) {
	ptr.initLog(app)
	app.Debug = ptr.Data.Base.Debug
	app.HideBanner = ptr.Data.Base.HideBanner
	app.DisableHTTP2 = ptr.Data.Base.DisableHTTP2

	app.Use(middleware.BodyLimit(ptr.Data.Base.BodyLimit))

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: ptr.Data.Base.CORS.AllowOrigins,
		AllowHeaders: ptr.Data.Base.CORS.AllowMethods,
	}))

	if ptr.Data.Base.CSRFTokenLookup != "" {
		app.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
			TokenLookup: ptr.Data.Base.CSRFTokenLookup,
		}))
	}

	app.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: ptr.Data.Base.GzipLevel,
	}))
}

func (ptr *ConfigMgr) initLog(app *echo.Echo) {
	switch strings.ToUpper(ptr.Data.Base.LogLevel) {
	case "DEBUG":
		app.Logger.SetLevel(log.DEBUG)
	case "INFO":
		app.Logger.SetLevel(log.INFO)
	case "WARN":
		app.Logger.SetLevel(log.WARN)
	case "ERROR":
		app.Logger.SetLevel(log.ERROR)
	case "OFF":
		app.Logger.SetLevel(log.OFF)
	}
}

func (ptr *ConfigMgr) GetServerConfig() *http.Server {
	res := http.Server{Addr: ptr.Data.Server.ListenAddr}
	res.ReadTimeout = time.Duration(ptr.Data.Server.ReadTimeout) * time.Minute
	res.WriteTimeout = time.Duration(ptr.Data.Server.WriteTimeout) * time.Minute
	return &res
}
