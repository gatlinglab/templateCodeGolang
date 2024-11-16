package modUtility

import (
	"fmt"

	"github.com/gatlinglab/libgatlinglog"
)

func log_initialize(url, token string) error {
	logInst, err := libgatlinglog.GatlingLogLib_Initialize(url, token, APPID)

	if err != nil {
		fmt.Println("log init error: ", err)
		return err
	}
	libgatlinglog.GatlingLogLib_SetDefault(logInst)

	return nil
}

func LogInfo(log string) error {
	err := libgatlinglog.GatlingLogLib_info(log)
	if err != nil {
		println("log info failed:")
		println(err.Error())
	}

	return err
}
func LogError(log string) error {
	err := libgatlinglog.GatlingLogLib_error(log)
	if err != nil {
		println("log error failed:", err)
	}

	return err
}
