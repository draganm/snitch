var id
var showDeleteConfirm = false
var config = {}
var log = []
var exists = false
function mount() {
  id = ctx.Params["targetID"]
	setTitle("Snitch: Target Details")
	listen(dbpath("targets",id), function(r) {
    if (r) {
      exists = true
      config = JSON.parse(r.Data(dbpath("config")))
      log = []
      r.ForEach(dbpath("log"), function(i, lr){
        log.push(JSON.parse(lr.Data(dbpath())))
      })
    } else {
      exists = false
      alerts=[{type:"danger", text:"Task does not exist"}]
    }
		render()
	});
}

function onUserEvent(evt) {
  if (evt.ElementID === "deleteButton") {
    showDeleteConfirm = true
    render()
  }

  if (evt.ElementID === "deleteConfirmButton") {
    deleteTarget(id)
    setLocation("/#/")
  }

  if (evt.ElementID === "deleteCancelButton" || evt.ElementID == "deleteConfirmModal") {
    showDeleteConfirm = false
    render()
  }
}


function unmount() {

}


function render() {
	var dm = targetDetailsDisplay.DeepCopy()
  if (exists) {

    dm.SetElementAttribute("panel","header","Target "+config.name+" ("+id+")")
    dm.SetElementText("image", config.image)
    dm.SetElementText("command", config.command)
    dm.SetElementText("interval", config.interval.toString())
    if (showDeleteConfirm) {
      dm.ReplaceChild("deleteButton", confirmDeleteModal)
    }
    for (var i=0; i<log.length; i++) {
      var le = logEvent.DeepCopy()
      var d = new Date(log[i].time)
      le.SetElementAttribute("title", "header", d.toString())
      for (var key in log[i].fields) {
        var lep = logEventProperty.DeepCopy()
        lep.SetElementText("name", key)
        lep.SetElementText("value", log[i].fields[key])
        le.AppendChild("properties",lep)
      }
      dm.AppendChild("log", le)
    }
    updateScreen(withNavigation(dm));
  } else {
    updateScreen(withNavigation(parseDisplayModel("<span/>")))
  }

}

"/targets/:targetID";
