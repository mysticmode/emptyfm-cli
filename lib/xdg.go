package lib

import (
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// assign to a local var to allow mocking in unit tests
var currentUser = user.Current

// Get the current user home directory (first from the $HOME env var and
// fallback on calling getpwuid_r() from libc if $HOME is unset).
func HomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		u, e := currentUser()
		if e == nil {
			home = u.HomeDir
		} else {
			log.Fatalf("HomeDir: %s (while handling %s)", e, err)
		}
	}

	return home
}

// Replace ~ with the current user's home dir
func ExpandHome(fragments ...string) string {
	home := HomeDir()

	res := path.Join(fragments...)
	if strings.HasPrefix(res, "~/") || res == "~" {
		res = home + strings.TrimPrefix(res, "~")
	}

	return res
}

// Return a path relative to the user home config dir
func ConfigPath(paths ...string) string {
	res := filepath.Join(paths...)

	if !filepath.IsAbs(res) {
		var config string

		if runtime.GOOS == "darwin" {
			config = os.Getenv("XDG_CONFIG_HOME")
			if config == "" {
				config = ExpandHome("~/Library/Preferences")
			}
		} else {
			var err error
			config, err = os.UserConfigDir()
			if err != nil {
				config = ExpandHome("~/.config")
			}
		}

		res = filepath.Join(config, res)
	}

	return res
}
