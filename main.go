package main

import (
	"embed"

	"github.com/aioceans/aiotools/internal/app"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	myApp := app.NewApp()
	packageInsight := app.PackageInsight{}
	urlParser := app.UrlParser{}
	structTransform := app.StructTransformRequest{}
	// Create application with options
	err := wails.Run(&options.App{
		Title:            "AIO Tools",
		Width:            1024,
		Height:           768,
		MinWidth:         800,
		MinHeight:        500,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        myApp.Startup,
		Bind: []interface{}{
			myApp,
			&packageInsight,
			&urlParser,
			&structTransform,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				HideToolbarSeparator: true,
			},
			Appearance:          mac.NSAppearanceNameDarkAqua,
			WindowIsTranslucent: true,
			About: &mac.AboutInfo{
				Title:   "My Application",
				Message: "© 2021 Me",
			},
		},
	})

	if err != nil {
		println("Error:", err)
	}
}
