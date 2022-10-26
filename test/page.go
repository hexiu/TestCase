/*
 * @Date: 2021-09-17 14:31:44
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-09-17 15:15:53
 * @FilePath: /TestCase/test/page.go
 */
package test

import (
	"math/rand"
)

func Page(link []int, page int) (List []int, next bool) {
	var FormatRespData []int
	var DefPageSize int = 20
	x := rand.Intn(DefPageSize)
	if len(link) > x {
		link = link[:len(link)-x]
	}
	FormatRespData = make([]int, len(link))
	copy(FormatRespData, link)
	if page*DefPageSize < (len(FormatRespData)) {
		next = true
		List = FormatRespData[(page-1)*DefPageSize : page*DefPageSize]
	} else if page*DefPageSize == (len(FormatRespData)) {
		next = false
		List = FormatRespData[(page-1)*DefPageSize : page*DefPageSize]
	} else {
		next = false
		if (page-1)*DefPageSize < (len(FormatRespData)) {
			List = FormatRespData[(page-1)*DefPageSize:]
		}
	}
	// fmt.Println(len(link), page, next, List)
	return
}
