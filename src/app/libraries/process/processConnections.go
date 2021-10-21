package process

import (
	"github.com/MKwann7/zgEXCELL-Socket/src/app/libraries/errorManagement"
	"github.com/MKwann7/zgEXCELL-Socket/src/app/libraries/helper"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func HandleConnections(responseWriter http.ResponseWriter, webRequest *http.Request) {

	user, validationError := ValidateConnection(webRequest)

	if validationError != nil {
		errorManagement.HandleErr(responseWriter, validationError, http.StatusBadRequest)
		log.Println(validationError.Error())
		return
	}

	upgradeConnection := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	socketConnection, upgradeError := upgradeConnection.Upgrade(responseWriter, webRequest, nil)

	if upgradeError != nil {
		errorManagement.HandleErr(responseWriter, upgradeError, http.StatusInternalServerError)
		return
	}

	defer socketConnection.Close()

	// Send initialization model to catch the user up on their command center/hub.

	for {

		notificationCheckResult := CheckForNewNotifications(user)

		if notificationCheckResult != nil {
			writeNotificationError := socketConnection.WriteMessage(1, notificationCheckResult)

			if writeNotificationError != nil {
				errorManagement.HandleErr(responseWriter, writeNotificationError, http.StatusInternalServerError)
				break
			}
		}

		inboundMessageResult := HandleInboundMessage(socketConnection, responseWriter)

		if inboundMessageResult != nil {
			writeNotificationError := socketConnection.WriteMessage(1, []byte(inboundMessageResult))

			if writeNotificationError != nil {
				errorManagement.HandleErr(responseWriter, writeNotificationError, http.StatusInternalServerError)
				break
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func HandleHealthCheck(responseWriter http.ResponseWriter, webRequest *http.Request) {

	healthCheck := helper.TransactionBool{Success: true}
	helper.JsonReturn(healthCheck, responseWriter)
}
