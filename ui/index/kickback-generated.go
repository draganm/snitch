package index

import reactor "github.com/draganm/go-reactor"

var indexDisplay = &reactor.DisplayModel{Element: "bs.Grid", Attributes: map[string]interface {
}{"fluid": true}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Row", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Col", Attributes: map[string]interface {
}{"md": 10, "mdOffset": 1}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Panel", Attributes: map[string]interface {
}{"header": "Targets"}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{ID: "noTargets", Element: "bs.Well", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "No Targets"}}}, &reactor.DisplayModel{ID: "targets", Element: "bs.ListGroup", Attributes: map[string]interface {
}{}}, &reactor.DisplayModel{Element: "bs.Button", Attributes: map[string]interface {
}{"href": "/#/add_target"}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Add Target"}}}}}}}}}}}
var indexTarget = &reactor.DisplayModel{ID: "item", Element: "bs.ListGroupItem", Attributes: map[string]interface {
}{}}
