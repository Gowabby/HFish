package ssh

import (
	"github.com/gliderlabs/ssh"
	"fmt"
)

func Start() {
	ssh.ListenAndServe(":2222", nil,
		ssh.PasswordAuth(func(s ssh.Context, password string) bool {
			fmt.Println("SSH Ip:" + s.RemoteAddr().String() + " User:" + s.User() + " Password:" + password)
			return false // false 代表 账号密码 不正确
		}),
	)
}
