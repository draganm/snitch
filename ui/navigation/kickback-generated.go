package navigation

import reactor "github.com/draganm/go-reactor"

var alert = &reactor.DisplayModel{ID: "alert", Element: "bs.Alert", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{ID: "text", Element: "strong", Attributes: map[string]interface {
}{}}}}
var navigation = &reactor.DisplayModel{Element: "div", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Navbar", Attributes: map[string]interface {
}{"collapseOnSelect": true, "fluid": true}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Navbar.Header", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Navbar.Brand", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "a", Attributes: map[string]interface {
}{"className": "navbar-brand", "href": "#"}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Snitch"}}}}}}}, &reactor.DisplayModel{Element: "bs.Navbar.Collapse", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Nav", Attributes: map[string]interface {
}{}}}}}}, &reactor.DisplayModel{Element: "bs.Grid", Attributes: map[string]interface {
}{"fluid": true}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Row", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Col", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{ID: "alerts", Element: "div", Attributes: map[string]interface {
}{}}}}}}, &reactor.DisplayModel{Element: "bs.Row", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Col", Attributes: map[string]interface {
}{"xs": 12}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{ID: "content", Element: "div", Attributes: map[string]interface {
}{"className": "container"}}}}}}}}}}
