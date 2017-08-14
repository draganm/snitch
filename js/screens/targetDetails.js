var id

function mount() {
  id = ctx.Params["targetID"]
	setTitle("Snitch: Target Details")
	listen(dbpath("targets",id), function(r) {
		var config = JSON.parse(r.Data(dbpath("config")))
		render(config)
	});
}

function onUserEvent(evt) {

}


function unmount() {

}


function render(config) {
	var dm = targetDetailsDisplay.DeepCopy()
  dm.SetElementAttribute("panel","header","Target "+config.name+" ("+id+")")
	// console.log("read targets",targets)


	// for (var i = 0; i<targets.length; i++) {
	// 	var it = indexTarget.DeepCopy()
	// 	var target = targets[i]
	// 	it.SetElementAttribute("item","bsStyle",statusStyle(target.status))
	// 	it.SetElementAttribute("item","href","#/targets/"+target.id)
	// 	it.SetElementText("item", target.name)
	// 	dm.AppendChild("targets", it)
	// }
	updateScreen(withNavigation(dm));
}

"/targets/:targetID";
