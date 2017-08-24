var alerts = []

function withNavigation(dm) {
  var navCopy = navigation.DeepCopy()
  for (var i=0; i<alerts.length; i++) {
    var al= alerts[i]
    var alEl = alert.DeepCopy()
    alEl.SetElementText("text", al.text)
    alEl.SetElementAttribute("alert","bsStyle", al.type)
    navCopy.AppendChild("alerts",alEl)
  }
  navCopy.ReplaceChild("content", dm);
  return navCopy;
}
