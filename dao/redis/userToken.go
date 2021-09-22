package redis

import (
	"turan.com/web_demo/models"
)

func SetToken(user *models.User, token string) (err error) {
	_, err = Client.Set(user.UserName, token, 0).Result()
	if err != nil {
		return err
	}
	return nil
}
