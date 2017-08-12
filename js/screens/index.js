function mount() {
	setTitle("Snitch: Main")
	listen(dbpath("targets"), function(r) {
		render();
	});
}

function onUserEvent(evt) {
	if (evt.ElementID === "counter") {
		tx(function(w) {
			var counter = JSON.parse(w.Data(dbpath("counter")))
			counter++
			w.CreateData(dbpath("counter"), JSON.stringify(counter))
		});
	}
}

function unmount() {

}

function render() {
	var dm = indexDisplay.DeepCopy();
	// console.Println("Counter", dm)
	// dm.SetElementText("counter", "Counter: "+counter)
	updateScreen(withNavigation(dm));
}

"/";
