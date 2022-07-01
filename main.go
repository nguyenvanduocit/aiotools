package main

import (
	"embed"

	"github.com/aioceans/aiotools/internal/app"
	"github.com/aioceans/aiotools/internal/dnslookup"
	"github.com/aioceans/aiotools/internal/getjson"
	"github.com/aioceans/aiotools/internal/packageinsights"
	"github.com/aioceans/aiotools/internal/parseurl"
	setjson "github.com/aioceans/aiotools/internal/setJson"
	"github.com/aioceans/aiotools/internal/transformstruct"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed frontend/dist
var assets embed.FS

func main() {

	myApp := app.NewApp()
	parseUrl := parseurl.NewPlugin()
	packageInsights := packageinsights.NewPlugin()
	transformStruct := transformstruct.NewPlugin()
	getJson := getjson.NewGetJson()
	setJson := setjson.NewSetJson()
	dnsLookup := dnslookup.NewPlugin()

	err := wails.Run(&options.App{
		Title:            "AIO Tools",
		Width:            1000,
		Height:           500,
		MinWidth:         800,
		MinHeight:        500,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		Menu:             myApp.AppMenu,
		OnStartup:        myApp.Startup,
		OnShutdown:       myApp.Shutdown,
		StartHidden:      true,
		Bind: []interface{}{
			myApp,
			parseUrl,
			packageInsights,
			transformStruct,
			getJson,
			setJson,
			dnsLookup,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				HideToolbarSeparator: true,
			},
			Appearance:           mac.DefaultAppearance,
			WindowIsTranslucent:  true,
			WebviewIsTransparent: true,
			About: &mac.AboutInfo{
				Title:   "AIO Tools",
				Message: "© 2022 AI Ocean",
			},
		},
	})

	if err != nil {
		println("Error:", err)
	}
}
