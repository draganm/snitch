var targets = []

function mount() {
	setTitle("Snitch: Main")
	listen(dbpath("targetsOrdered"), function(r) {
		targets = JSON.parse(r.Data(dbpath()))
		render()
	});
}

function onUserEvent(evt) {

}


function unmount() {

}

function render() {
	var dm = indexDisplay.DeepCopy()
	// console.log("read targets",targets)

	for (var i = 0; i<targets.length; i++) {
		var it = indexTarget.DeepCopy()
		it.SetElementText("item", targets[i])
		dm.AppendChild("targets", it)
	}
	updateScreen(withNavigation(dm));
}

"/";
