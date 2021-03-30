package bcrypt

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestCheckPwd(t *testing.T) {
	flag := CheckPw("123456", "$2b$12$u5ZBnqIpuoCCu31Ai0f3uugdJVD/FqwaJUDSiqaHvgeZY5Qy0XVMq")
	assert.Equal(t, true, flag)
}

func TestGenPwd(t *testing.T) {
	pwd := HashPw("123456", GenSalt("$2b", 12))
	flag := CheckPw("123456", pwd)
	assert.Equal(t, true, flag)
	pwd2 := HashPw("123456", GenSalt("$2b", 12))
	flag2 := CheckPw("654321", pwd2)
	assert.Equal(t, false, flag2)

}
