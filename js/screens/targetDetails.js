var id
var showDeleteConfirm = false
var config = {}
function mount() {
  id = ctx.Params["targetID"]
	setTitle("Snitch: Target Details")
	listen(dbpath("targets",id), function(r) {
		config = JSON.parse(r.Data(dbpath("config")))
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
  dm.SetElementAttribute("panel","header","Target "+config.name+" ("+id+")")
  dm.SetElementText("image", config.image)
  dm.SetElementText("command", config.command)
  if (showDeleteConfirm) {
    dm.ReplaceChild("deleteButton", confirmDeleteModal)
  }
	updateScreen(withNavigation(dm));
}

"/targets/:targetID";
