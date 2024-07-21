package main

import (
	"BoostTool/Core/Bot"
	"BoostTool/Core/Discord"
	"BoostTool/Core/Utils"
	"time"

	title "github.com/lxi1400/GoTitle"

	KeyAuthApp "BoostTool/Core/KeyAuth"
	"fmt"
)

func Input(message string) string {
	fmt.Print(message)

	var input string
	fmt.Scanln(&input)
	return input
}

func main() {
	KeyAuthApp.Api(
		"DeadXFetcher", // Application name
		"ZueBudVOib",   // Owner ID
		"76cf001702a1ad91fba8710859a90d0d1167f4cb699f3012fd0b335bcd2f2ce4", // Application Secret
		"1.0", // Application Version
		"",    // Token Path (PUT NULL OR LEAVE BLANK IF YOU DO NOT WANT TO USE THE TOKEN VALIDATION SYSTEM! MUST DISABLE VIA APP SETTINGS)
	)

	config, _ := Utils.LoadConfig() // Ignoring the error for simplicity
	license := config.Bot.License

	Utils.LogInfo("Checking your license...", "", "")
	KeyAuthApp.License(license)

	// Continue with the rest of your program
	Utils.ClearScreen()

	time.Sleep(time.Second * 1)
	Utils.LogInfo("Logging in.", "", "")

	title.SetTitle("Boost Bot | .gg/b11 | ")
	Utils.ClearScreen()
	Utils.PrintASCII()

	if config.CustomPersonalization.Onliner {
		Utils.LogInfo("Token Onliner: Enabled", "", "")
		go Discord.Websocket()
	} else {
		Utils.LogInfo("Token Onliner: Disabled", "", "")
	}

	go Bot.Automation()
	Bot.StartBot()
}
