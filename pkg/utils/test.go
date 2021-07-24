package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// CreateIfNotExist 如果不存在则创建
func CreateIfNotExist(p string, b []byte, mode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return err
	}
	os.Create(p)
	err := ioutil.WriteFile(p, b, mode)
	if err != nil {
		return err
	}
	return nil
}

func CreateForTest() {
	testCase := `db:
  type: 'postgres'
  user: postgres
  password : qqwse780
  host: db
  port: 5432
  name: blog
  db_name: owl
  tablePrefix: blog_
  # sqlite专用
  file_path: './data/owl.db?cache_size=50'
# 搜索引擎 meilisearch
searchIndex: 'blog'
searchEngine: 'meilisearch'
meilisearch:
  host: 'http://search_engine:7700'
  apiKey: 'masterKey'
smtp:
  name: '' # 发送者名
  address: '' # 发送者地址
  reply_to: '' # 回复地址
  host: '' # 服务器主机名
  port: 587 # 服务器端口
  user: '' # 用户名
  password: '' # 密码
  encryption: false # 是否启用加密
  keep_alive: '' # SMTP 连接保留时长
# 默认设置
default_setting:
  domain: 'owlnet.xyz'
  log_path: './data/logs/owl.log'
# 账号相关默认设置，可以通过博客后台修改
account:
  username: '少喝咖啡'
  password: '89898902'
  email: waterservers@163.com
  phone: ""
  address: ""
# 博客相关默认设置
blogger:
  name: 'owl'
  subtitle: ''
  # 备案号
  record_filing: ''`
	err := CreateIfNotExist("config/app.yml", []byte(testCase), 0644)
	if err != nil {
		panic(err)
	}
}

func IsInTests() bool {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.run") {
			return true
		}
	}
	return false
}
