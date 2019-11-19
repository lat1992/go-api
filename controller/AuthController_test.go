/*
 * File: X:\go-api\controller\AuthController_test.go
 * Created At: 2019-11-19 17:01:05
 * Created By: Mauhoi WU
 * 
 * Modified At: 2019-11-19 17:29:15
 * Modified By: Mauhoi WU
 */

package controller

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	if GenerateToken(-1) != "" {
		t.Error("GenerateToken FAIL")
	}
	if GenerateToken(1) != "" {
		t.Log("GenerateToken PASS")
	}
}
