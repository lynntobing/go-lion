/**
 * @Author: Lynn
 * @Time  : 2020-05-23 1:45
 * @File  : strutil.go
 * @Description: ...
 */
package utils

import "bytes"

func StitchingStr(args ...string) (rstStr string) {
	var buf bytes.Buffer
	for _, v := range args {
		buf.WriteString(v)
	}
	rstStr = buf.String()
	return
}
