package setting

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config 博客配置
type Config struct {
	RunMode      string // 运行模式
	SearchEngine string `yaml:"searchEngine"` // 搜索引擎
	SearchIndex  string `yaml:"searchIndex"`
	Database     struct {
		Type        string `yaml:"type"`
		User        string `yaml:"user"`
		Password    string `yaml:"password"`
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		Name        string `yaml:"name"`
		DBName      string `yaml:"db_name"`
		FilePath    string `yaml:"file_path"`
		TablePrefix string
	} `yaml:"db"`
	Meilisearch struct {
		Host   string `yaml:"host"`
		APIKey string `yaml:"apiKey"`
	} `yaml:"meilisearch"`
	SMTP struct {
		Host       string
		Port       int
		Encryption bool
	} `yaml:"smtp"`
	DefaultSetting struct {
		Domain  string `yaml:"domain"`
		LogPath string `yaml:"log_path"`
	} `yaml:"default_setting"`
	Account struct {
		Username string `yaml:"username"`
		Password string
		Email    string
		Phone    string
		Address  string
	} `yaml:"account"`
}

const (
	// 该模式会输出 debug 等信息
	DEV  = "dev"
	PROD = "prod" // 该模式用于生产环境
)

var (
	Conf    = &Config{}
	BlackIP = make(map[string]bool)
)

func IsInTests() bool {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.run") {
			return true
		}
	}
	return false
}

func Init() {
	// 初始化配置
	var data []byte
	var err error
	if IsInTests() == true {
		fmt.Println("go test 模式运行")
		//flag不为空,则说明存在test所拥有的参数,是在 go test 模式
		return
	} else {
		fmt.Println("正常模式运行 run")
	}
	currentMode := os.Getenv("MODE")
	if currentMode == "dev" {
		data, err = ioutil.ReadFile("config/app.dev.yml")
	} else {
		data, err = ioutil.ReadFile("config/app.yml")
	}
	if err != nil {
		log.Fatalf("读取配置失败:%s", err)
	}
	err = yaml.Unmarshal(data, Conf)
	if err != nil {
		log.Fatalf("解析配置:%s", err)
	}
}
