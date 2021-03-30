package utils

import (
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
)

func TestCheckPwd(t *testing.T) {
	flag := CheckPw("123456", "$2b$12$u5ZBnqIpuoCCu31Ai0f3uugdJVD/FqwaJUDSiqaHvgeZY5Qy0XVMq")
	assert.Equal(t, true, flag)
}

func TestGenPwd(t *testing.T) {
	salt := GenSalt("$2b", 12)
	pwd := HashPw("123456", salt)
	fmt.Println(pwd)
	flag := CheckPw("123456", pwd)
	assert.Equal(t, true, flag)

}
