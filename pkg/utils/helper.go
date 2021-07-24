package utils

import (
	"bytes"
	"crypto/rand"
	"errors"
	"io"
	"math/big"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"go-template/pkg/log"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/html"
)

// What's bcrypt? https://en.wikipedia.org/wiki/Bcrypt
// Golang bcrypt doc: https://godoc.org/golang.org/x/crypto/bcrypt
// You can change the value in bcrypt.DefaultCost to adjust the security index.
// 	err := userModel.setPassword("password0")
func SetPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	result := string(passwordHash)
	return result, nil
}

// Database will only save the hashed string, you should check it by util function.
// 	if err := serModel.checkPassword("password0"); err != nil { password error }
func CheckPassword(password string, pwdHash string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(pwdHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// IgnoreHTMLTag 去掉 html tag
func IgnoreHTMLTag(src string) string {
	// 去除所有尖括号内的HTML代码
	re, _ := regexp.Compile(`<[\S\s]+?>`)
	// re, _ := regexp.Compile(`<[^>]*>`)
	src = re.ReplaceAllString(src, "")

	// 去除换行符
	re, _ = regexp.Compile(`\s+`)
	return re.ReplaceAllString(src, "")
}

// SumDayFromNow 计算与当前相差天数
func SumDayFromNow(d time.Time) int64 {
	now := time.Now().Unix()
	day := now - d.Unix()
	return day / 86400
}

// FuzzyContent 查找模糊内容
func FuzzyContent(query string, content string, num int) string {
	pos := strings.Index(content, query)
	if pos == -1 {
		log.Warnf("未查找搜索内容，关键词：%s", query)
	}
	strLen := len(content)
	left := 0
	if pos-num/2 > 0 {
		left = pos - num/2
	}
	right := strLen
	if pos+num/2 < strLen {
		right = pos + num/2
	}
	var result string
	for !utf8.RuneStart(content[left]) && left < pos {
		left++
	}
	for right < strLen && right > pos && !utf8.RuneStart(content[right]) {
		right--
	}
	result = content[left:right]
	if left > 0 {
		result = "..." + result
	}
	if right < strLen {
		result = result + "..."
	}
	return result
}

func extract(node *html.Node, buff *bytes.Buffer) {
	if node.Type == html.TextNode {
		data := strings.Trim(node.Data, "\r\n ")
		if data != "" {
			buff.WriteString("\n")
			buff.WriteString(data)
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extract(c, buff)
	}
}

// ExtractText 提取html中的text
func ExtractText(reader io.Reader) string {
	var buffer bytes.Buffer
	doc, err := html.Parse(reader)
	if err != nil {
		log.Error("ExtractText error")
	}
	extract(doc, &buffer)
	re, _ := regexp.Compile(`\s+`)
	return re.ReplaceAllString(buffer.String(), "")
}

// GenerateRandomString 生辰随机token
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
