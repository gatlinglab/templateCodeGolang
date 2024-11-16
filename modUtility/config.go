/*
	APPID: defined when this program start writing;
	G_AppToken: defined in config; if not defined. then random one;

*/

package modUtility

import (
	"strconv"

	"github.com/gatlinglab/libgatlingconfig"
)

const APPID = "WJTEMPLATE1"
const APPTOKEN = "APPTOKEN"

// {{$HTTPSERVER}}
const C_Key_HttpPort = "HTTPPORT"

// {{$HTTPSERVER_END}}

const C_Key_LogUrl = "LOGURL"
const C_Key_LogToken = "LOGTOKEN"

var G_AppToken = ""

// {{$HTTPSERVER}}
var G_HttpPort = 10000

// {{$HTTPSERVER_END}}

func GetSystemID() string {
	return ""
}

func config_initialize() error {
	err := libgatlingconfig.GetSingleGatlingConfig().Initialize(APPID)
	if err != nil {
		return err
	}

	G_AppToken = Config_Read(APPTOKEN)

	// {{$HTTPSERVER}}
	strServerPort := Config_Read("HTTPPORT")
	if strServerPort != "" {
		iRet, err := strconv.Atoi(strServerPort)
		if err == nil && iRet > 10 && iRet < 65535 {
			G_HttpPort = iRet
		}
	}
	// {{$HTTPSERVER_END}}

	return nil
}

func Config_Read(key string) string {
	return libgatlingconfig.GetSingleGatlingConfig().Get(key)
}
