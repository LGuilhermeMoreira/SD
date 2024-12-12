package user

import (
	"chamda_remota/client/proxy"
	"strings"
)

type User struct {
	proxy proxy.Proxy
}

// implementar interação com usuario

func (u *User) Run(input string) {
	action := strings.Split(input, " ")[1]
	switch action {
	case "+":

	}
}
