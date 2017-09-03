package target

import reactor "github.com/draganm/go-reactor"

var display = &reactor.DisplayModel{ID: "mainGrid", Element: "bs.Grid", Attributes: map[string]interface {
}{"fluid": true}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Row", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Col", Attributes: map[string]interface {
}{"md": 8, "mdOffset": 0}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{ID: "panel", Element: "bs.Panel", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Table", Attributes: map[string]interface {
}{"bordered": true, "condensed": true, "responsive": false}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "thead", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "th", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Name"}}}, &reactor.DisplayModel{Element: "th", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Value"}}}}}, &reactor.DisplayModel{Element: "tbody", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "tr", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "td", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Image"}}}, &reactor.DisplayModel{ID: "image", Element: "td", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "a"}}}}}, &reactor.DisplayModel{Element: "tr", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "td", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Command"}}}, &reactor.DisplayModel{ID: "command", Element: "td", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "a"}}}}}, &reactor.DisplayModel{Element: "tr", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "td", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Interval"}}}, &reactor.DisplayModel{ID: "interval", Element: "td", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "a"}}}}}}}}}, &reactor.DisplayModel{ID: "deleteButton", Element: "bs.Button", Attributes: map[string]interface {
}{"bsStyle": "danger"}, ReportEvents: []reactor.ReportEvent{reactor.ReportEvent{Name: "click", PreventDefault: true}}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Delete"}}}}}}}}}}}
