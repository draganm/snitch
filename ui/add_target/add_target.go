package add_target

//go:generate kickback-generator -p add_target

import (
	"fmt"
	"regexp"
	"strconv"

	reactor "github.com/draganm/go-reactor"
	"github.com/draganm/kickback"
	"github.com/draganm/snitch/executor"
	"github.com/draganm/snitch/tx"
	. "github.com/draganm/snitch/ui/navigation"
)

func style(valid bool) string {
	if valid {
		return "success"
	}
	return "error"
}

var nameRegexp = regexp.MustCompile(`^.+$`)

var imageRegexp = regexp.MustCompile(`^[a-zA-Z0-9./_:%\-]+$`)

var commandRegexp = regexp.MustCompile(`^.+$`)

var intervalRegexp = regexp.MustCompile(`^\d+$`)

func init() {
	kickback.AddScreen("/add_target", func(ctx *kickback.Context) {

		alerts := []AlertLine{}

		name := ""
		image := ""
		command := ""
		interval := ""

		var isNameValid = func() bool {
			return nameRegexp.MatchString(name)
		}

		var isImageValid = func() bool {
			return imageRegexp.MatchString(image)
		}

		var isCommandValid = func() bool {
			return commandRegexp.MatchString(command)
		}

		var isIntervalValid = func() bool {
			return intervalRegexp.MatchString(interval)
		}

		var canSubmit = func() bool {
			return isNameValid() && isImageValid() && isCommandValid() && isIntervalValid()
		}

		submitEnabled := canSubmit()
		nameValid := isNameValid()
		imageValid := isImageValid()
		commandValid := isCommandValid()
		intervalValid := isIntervalValid()

		var render = func() {
			dm := display.DeepCopy()
			dm.SetElementAttribute("submitButton", "disabled", !submitEnabled)
			dm.SetElementAttribute("submitButton", "bsStyle", style(submitEnabled))
			dm.SetElementAttribute("nameFormGroup", "validationState", style(nameValid))
			dm.SetElementAttribute("imageFormGroup", "validationState", style(imageValid))
			dm.SetElementAttribute("commandFormGroup", "validationState", style(commandValid))
			dm.SetElementAttribute("intervalFormGroup", "validationState", style(intervalValid))
			ctx.ScreenContext.UpdateScreen(&reactor.DisplayUpdate{
				Model: WithNavigation(dm, alerts),
			})
		}

		render()

		ctx.OnUserEventFunc = func(evt *reactor.UserEvent) {

			if evt.Type == "change" {
				switch evt.ElementID {
				case "name":
					name = evt.Value
					if nameValid != isNameValid() {
						nameValid = isNameValid()
						render()
					}
				case "image":
					image = evt.Value
					if imageValid != isImageValid() {
						imageValid = isImageValid()
						render()
					}
				case "command":
					command = evt.Value
					if commandValid != isCommandValid() {
						commandValid = isCommandValid()
						render()
					}
				case "interval":
					interval = evt.Value
					if intervalValid != isIntervalValid() {
						intervalValid = isIntervalValid()
						render()
					}
				default:
					return
				}
				if submitEnabled != canSubmit() {
					submitEnabled = canSubmit()
					render()
				}
			}
			if evt.Type == "submit" {

				intervalInt, err := strconv.Atoi(interval)
				if err != nil {
					alerts = append(alerts, AlertLine{
						Text: fmt.Sprintf("Error parsing interval: %s", err),
						Typ:  "danger",
					})
					render()
					return
				}

				c := &executor.Config{
					Name:     name,
					Image:    image,
					Command:  command,
					Interval: intervalInt,
				}

				id, err := tx.TX{ctx.DB}.AddTarget(c)
				if err != nil {
					alerts = append(alerts, AlertLine{
						Text: fmt.Sprintf("Error parsing interval: %s", err),
						Typ:  "danger",
					})
					render()
					return
				}
				ctx.ScreenContext.UpdateScreen(&reactor.DisplayUpdate{
					Location: fmt.Sprintf("#/targets/%s", id),
				})
			}

		}
	})
}
