var targets = []

function mount() {
	setTitle("Snitch: Main")
	listen(dbpath("status"), function(r) {
		targets = JSON.parse(r.Data(dbpath()))
		render()
	});
}

function onUserEvent(evt) {

}


function unmount() {

}

function statusStyle(status) {
	switch (status) {
		case "unknown":
			return "warning"
	}

}

function render() {
	var dm = indexDisplay.DeepCopy()
	// console.log("read targets",targets)


	if (targets.length > 0) {
		dm.DeleteChild("noTargets")
	}

	for (var i = 0; i<targets.length; i++) {
		var it = indexTarget.DeepCopy()
		var target = targets[i]
		it.SetElementAttribute("item","bsStyle",statusStyle(target.status))
		it.SetElementAttribute("item","href","#/targets/"+target.id)
		it.SetElementText("item", target.name)
		dm.AppendChild("targets", it)
	}
	updateScreen(withNavigation(dm));
}

"/";
