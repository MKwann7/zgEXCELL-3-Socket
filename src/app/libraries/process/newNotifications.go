package process

import (
	"encoding/json"
	"github.com/MKwann7/zgEXCELL-Socket/src/app/dtos"
	"strconv"
)

func CheckForNewNotifications(user *dtos.User) []byte {

	notifications := dtos.Notifications{}
	collection, err := notifications.GetWhere("user_id = '"+strconv.Itoa(user.UserId)+"' AND status = 'pending'", "ASC", 1)

	if err != nil {
		return nil
	}

	res1B, _ := json.Marshal(collection)

	return res1B
}
