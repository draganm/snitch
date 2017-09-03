package ui

import reactor "github.com/draganm/go-reactor"

type alertLine struct {
	text string
	typ  string
}

func withNavigation(dm *reactor.DisplayModel, alerts []alertLine) *reactor.DisplayModel {
	var navCopy = navigation.DeepCopy()

	for _, a := range alerts {
		alertElement := alert.DeepCopy()
		alertElement.SetElementText("text", a.text)
		alertElement.SetElementAttribute("alert", "bsStyle", a.typ)
	}

	navCopy.ReplaceChild("content", dm)
	return navCopy
}
