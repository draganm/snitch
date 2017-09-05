package navigation

//go:generate kickback-generator -p navigation

import reactor "github.com/draganm/go-reactor"

type AlertLine struct {
	Text string
	Typ  string
}

func WithNavigation(dm *reactor.DisplayModel, alerts []AlertLine) *reactor.DisplayModel {
	var navCopy = navigation.DeepCopy()

	for _, a := range alerts {
		alertElement := alert.DeepCopy()
		alertElement.SetElementText("text", a.Text)
		alertElement.SetElementAttribute("alert", "bsStyle", a.Typ)
		navCopy.AppendChild("alerts", alertElement)
	}

	navCopy.ReplaceChild("content", dm)
	return navCopy
}
