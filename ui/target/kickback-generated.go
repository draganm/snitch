package target

import reactor "github.com/draganm/go-reactor"

var confirmDeleteModal = &reactor.DisplayModel{ID: "deleteConfirmModal", Element: "bs.Modal", Attributes: map[string]interface {
}{"show": true}, ReportEvents: []reactor.ReportEvent{reactor.ReportEvent{Name: "hide"}}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Modal.Header", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Modal.Title", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Confirm deleting target"}}}}}, &reactor.DisplayModel{Element: "bs.Modal.Body", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "\n      You are about to delete target \""}, &reactor.DisplayModel{ID: "targetName", Element: "strong", Attributes: map[string]interface {
}{}}, &reactor.DisplayModel{Text: "\". Are you sure?\n    "}}}, &reactor.DisplayModel{Element: "bs.Modal.Footer", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{ID: "deleteConfirmButton", Element: "bs.Button", Attributes: map[string]interface {
}{"bsStyle": "danger"}, ReportEvents: []reactor.ReportEvent{reactor.ReportEvent{Name: "click", PreventDefault: true}}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Delete"}}}, &reactor.DisplayModel{ID: "deleteCancelButton", Element: "bs.Button", Attributes: map[string]interface {
}{"bsStyle": "primary"}, ReportEvents: []reactor.ReportEvent{reactor.ReportEvent{Name: "click", PreventDefault: true}}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Cancel"}}}}}}}
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
