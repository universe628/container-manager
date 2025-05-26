package auth

import (
	"container-manager/internal/infra/config"
	"container-manager/internal/utils"
)

func PasswordHashing(str string) string {
	return utils.ToHex(utils.Sha512Hash(str + config.GetConfig().Salt))
}
