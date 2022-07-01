package app

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx     context.Context
	AppMenu *menu.Menu
}

func NewApp() *App {

	app := &App{}

	app.AppMenu = menu.NewMenuFromItems(
		menu.AppMenu(),
		menu.EditMenu(),
		&menu.MenuItem{
			Type:  menu.SubmenuType,
			Label: "Browser",
			SubMenu: menu.NewMenuFromItems(
				menu.Text("Refresh", keys.CmdOrCtrl("r"), app.onBrowserRefreshMenuClicked),
			),
		},
	)

	return app
}

func (app *App) onBrowserRefreshMenuClicked(_ *menu.CallbackData) {
	runtime.EventsEmit(app.ctx, "browser.refresh")
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Shutdown(ctx context.Context) {
	a.ctx = ctx
}
