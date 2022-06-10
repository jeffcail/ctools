package encrypt

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

// StrMd5
func StrMd5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5Str := fmt.Sprintf("%x", w.Sum(nil))
	return md5Str
}

// StrMd5ToU md5加密 转大写
func StrMd5ToU(str string) string {
	strMd5 := StrMd5(str)
	return strings.ToUpper(strMd5)
}
