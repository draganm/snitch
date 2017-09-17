package add_target

import reactor "github.com/draganm/go-reactor"

var display = &reactor.DisplayModel{Element: "bs.Grid", Attributes: map[string]interface {
}{"fluid": true}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Row", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Col", Attributes: map[string]interface {
}{"md": 6, "mdOffset": 1}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.PageHeader", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Add Target"}}}}}}}, &reactor.DisplayModel{Element: "bs.Row", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.Col", Attributes: map[string]interface {
}{"md": 6, "mdOffset": 1}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{ID: "form", Element: "form", Attributes: map[string]interface {
}{}, ReportEvents: []reactor.ReportEvent{reactor.ReportEvent{Name: "submit", PreventDefault: true}}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{ID: "nameFormGroup", Element: "bs.FormGroup", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.ControlLabel", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Name"}}}, &reactor.DisplayModel{ID: "name", Element: "bs.FormControl", Attributes: map[string]interface {
}{"label": "Name", "placeholder": "Name", "type": "text"}, ReportEvents: []reactor.ReportEvent{reactor.ReportEvent{Name: "change"}}}, &reactor.DisplayModel{Element: "bs.HelpBlock", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Name of the target"}}}}}, &reactor.DisplayModel{ID: "imageFormGroup", Element: "bs.FormGroup", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.ControlLabel", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Image"}}}, &reactor.DisplayModel{ID: "image", Element: "bs.FormControl", Attributes: map[string]interface {
}{"placeholder": "Image", "type": "text"}, ReportEvents: []reactor.ReportEvent{reactor.ReportEvent{Name: "change"}}}, &reactor.DisplayModel{Element: "bs.HelpBlock", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Docker image used to run the command"}}}}}, &reactor.DisplayModel{ID: "commandFormGroup", Element: "bs.FormGroup", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.ControlLabel", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Command"}}}, &reactor.DisplayModel{ID: "command", Element: "bs.FormControl", Attributes: map[string]interface {
}{"placeholder": "Command", "type": "text"}, ReportEvents: []reactor.ReportEvent{reactor.ReportEvent{Name: "change"}}}, &reactor.DisplayModel{Element: "bs.HelpBlock", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Command to run with the docker image"}}}}}, &reactor.DisplayModel{ID: "intervalFormGroup", Element: "bs.FormGroup", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Element: "bs.ControlLabel", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Interval"}}}, &reactor.DisplayModel{ID: "interval", Element: "bs.FormControl", Attributes: map[string]interface {
}{"placeholder": "Interval", "type": "text"}, ReportEvents: []reactor.ReportEvent{reactor.ReportEvent{Name: "change"}}}, &reactor.DisplayModel{Element: "bs.HelpBlock", Attributes: map[string]interface {
}{}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Interval in seconds between runs"}}}}}, &reactor.DisplayModel{ID: "submitButton", Element: "bs.Button", Attributes: map[string]interface {
}{"bsStyle": "danger", "disabled": true, "type": "submit"}, Children: []*reactor.DisplayModel{&reactor.DisplayModel{Text: "Add Target"}}}}}}}}}}}
