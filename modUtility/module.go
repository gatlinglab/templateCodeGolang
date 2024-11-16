package modUtility

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func Utility_Initialize() error {
	err := config_initialize()
	if err != nil {
		return err
	}

	urlLog := Config_Read(C_Key_LogUrl)
	if len(urlLog) < 10 {
		return errors.New("url of log too short")
	}
	tokenLog := Config_Read(C_Key_LogToken)
	err = log_initialize(urlLog, tokenLog)
	if err != nil {
		return err
	}

	if len(G_AppToken) < 3 {
		LogError("apptoken empty")
		return errors.New("apptoken empty")
	}

	return nil
}

func Utility_writeStartLog() error {
	instid := utility_checkInstID() // read APPTOKEN which used in libgatlingconfig
	strTimeNow := time.Now().Format("2006-01-02 15:04:05")

	strLog := fmt.Sprintf("[%s-%s] start, local time: %s", APPID, instid, strTimeNow)

	return LogInfo(strLog)
}

func utility_checkInstID() string {
	instid := Config_Read(APPTOKEN) // this key defined in libgatlingconfig
	if instid != "" {
		return instid
	}
	instid = strconv.Itoa(int(time.Now().Unix()))

	return instid
}
