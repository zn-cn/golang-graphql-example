package util

import (
	"regexp"
)

const (
	// 用户昵称的正则匹配, 合法的字符有 0-9, A-Z, a-z, _,  [空格],汉字
	// 字符 '_' 只能出现在中间且不能重复, 如 "__"
	nicknamePattern = `^[a-z0-9A-Z \p{Han}]+(_[a-z0-9A-Z\p{Han}]+)*$`
	// 电子邮箱的正则匹配, 考虑到各个网站的 mail 要求不一样, 这里匹配比较宽松
	// 邮箱用户名可以包含 0-9, A-Z, a-z, -, _, .
	// 开头字母不能是 -, _, .
	// 结尾字母不能是 -, _, .
	// -, _, . 这三个连接字母任意两个不能连续, 如不能出现 --, __, .., -_, -., _.
	// 邮箱的域名可以包含 0-9, A-Z, a-z, -
	// 连接字符 - 只能出现在中间, 不能连续, 如不能 --
	// 支持多级域名, x@y.z, x@y.z.w, x@x.y.z.w.e
	mailPattern = `^[a-z0-9A-Z]+([\-_\.][a-z0-9A-Z]+)*@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)*\.)+[a-zA-Z]+$`
)

// CheckEmail util
func CheckEmail(email string) (b bool) {
	if ok, _ := regexp.MatchString(mailPattern, email); !ok {
		return false
	}
	return true
}

// CheckNickName util
func CheckNickName(nickname string) bool {
	if ok, _ := regexp.MatchString(nicknamePattern, nickname); !ok {
		return false
	}
	return true
}
