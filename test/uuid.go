/*
 * @Date: 2021-11-15 13:37:03
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-11-15 13:37:04
 * @FilePath: /TestCase/test/uuid.go
 */
package test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/twinj/uuid"
)

func UUID(businessType string) string {
	uuStr := uuid.NewV4().String()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := 0
	for index := 0; index < 10; index++ {
		random = r.Intn(100)
	}

	return fmt.Sprintf("interact_%s_%d_%s_%d", businessType, time.Now().UnixNano(), uuStr, random)

}
