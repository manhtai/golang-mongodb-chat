package config

import (
	"math"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/gplus"
)

func init() {
	store := sessions.NewFilesystemStore(os.TempDir(), []byte("cusbot"))
	store.MaxLength(math.MaxInt64)
	gothic.Store = store
}

var providers = make(map[string]*gplus.Provider)

// CreateProvider return provider with corresponding redirect_uri
// We only init goth with Google+ for now
func CreateProvider(redirectURI string) {

	if providers[redirectURI] == nil {
		providers[redirectURI] = gplus.New(
			os.Getenv("GPLUS_KEY"),
			os.Getenv("GPLUS_SECRET"),
			redirectURI,
		)
	}

	goth.UseProviders(
		providers[redirectURI],
	)
}
