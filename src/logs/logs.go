package logs

import (
	"github.com/cihub/seelog"
	"fmt"
)

var Logger seelog.LoggerInterface

func loadAppConfig() {
	appConfig := `
<seelog minlevel="warn">
    <outputs formatid="common">
        <rollingfile filename="/Users/berryjam/logs/roll.log" maxrolls="5" maxsize="100000" type="size"/>
        <filter levels="critical">
            <file formatid="critical" path="c/Users/berryjam/logs/critical.log"/>
        </filter>
    </outputs>
    <formats>
        <format format="%Date/%Time [%LEV] %Msg%n" id="common"/>
        <format format="%File %FullPath %Func %Msg%n" id="critical"/>
        <format format="Critical error on our server!\n    %Time %Date %RelFile %Func %Msg \nSent by Seelog" id="criticalemail"/>
    </formats>
</seelog>`
	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

// UseLogger uses a specified seelog.LoggerInterface to output library log.
// Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}

func init() {
	DisableLog()
	loadAppConfig()
}

// DisableLog disables all library log output
func DisableLog() {
	Logger = seelog.Disabled
}
