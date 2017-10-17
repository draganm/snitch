package favicon

import (
	"github.com/urfave/negroni"
	"gitlab.netice9.com/dragan/favicon/data"
)

// NegroniHandler will serve all favicon assets.
// Make sure this handler is not authenticated.
var NegroniHandler = negroni.NewStatic(data.AssetFS())
