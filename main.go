package main

import (
	"flag"
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	"log"
)

// Constants
const htmlAbout = `Welcome `

// Vars
var (
	AppName            string
	BuiltAt            string
	VersionAstilectron string
	VersionElectron    string
)

var (
	debug = flag.Bool("d", false, "enables debug")
	w     *astilectron.Window
)

func main() {
	// Init
	flag.Parse()

	// Create logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	l.Printf("Running bootstrap at %s\n", BuiltAt)
	// Run bootstrap
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
			SingleInstance:     true,
			VersionAstilectron: VersionAstilectron,
			VersionElectron:    VersionElectron,
		},
		Debug:  *debug,
		Logger: l,

		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astilectron.StrPtr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Label: astilectron.StrPtr("About")},
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		OnWait: func(_ *astilectron.Astilectron, iw *astilectron.Window, _ astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = iw
			return nil
		},
		WindowOptions: &astilectron.WindowOptions{
			BackgroundColor: astilectron.StrPtr("#333"),
			Center:          astilectron.PtrBool(true),
			Height:          astilectron.PtrInt(700),
			Width:           astilectron.PtrInt(700),
		},
	}); err != nil {
		fmt.Println("running bootstrap failed")
	}
}
