/**
 * @Author: Lynn
 * @Time  : 2020-05-23 0:58
 * @File  : main.go
 * @Description: ...
 */
package main

import "go-lion/src/router"

func main() {
	r := router.Router
	router.SetRouter()
	r.Run(":2009")
}
