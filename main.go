package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"
	"wjtemplate1/modHttpServer"
	"wjtemplate1/modUtility"
)

func main() {
	catchCrash()
	err := modUtility.Utility_Initialize()
	if err != nil {
		fmt.Println("utility init error: ", err)
		return
	}
	// {{$STARTFIRSTLOG}}
	modUtility.Utility_writeStartLog()
	// {{$STARTFIRSTLOG_END}}

	// {{$HTTPSERVER}}
	err = modHttpServer.Http_Initialize()
	if err != nil {
		modUtility.LogError("http init failed: " + err.Error())
		return
	}
	// {{$HTTPSERVER_END}}

	// {{$HTTPSERVER}}
	// $OPTIONMAIN
	err = modHttpServer.Http_Start()
	if err != nil {
		modUtility.LogError("http start failed: " + err.Error())
		return
	}
	// $OPTIONGO
	// go modHttpServer.Http_Start()
	// $OPTIONEND

	// {{$HTTPSERVER_END}}

}

func catchCrash() {
	t := time.Now()
	f, err := os.Create(t.Format("crash_20060102150405.log"))
	if err != nil {
		panic(err)
	}
	debug.SetCrashOutput(f, debug.CrashOptions{})
}
