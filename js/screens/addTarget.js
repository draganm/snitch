var name = ""
var image = ""
var command = ""

var nameValid = false;
var imageValid = false;
var commandValid = false;


var submitEnabled = false

function mount() {
  setTitle("Snitch: add target")
  render()
  nameValid = isNameValid()
  imageValid = isImageValid()
  commandValid = isCommandValid()
  submitEnabled = canSubmit()
}


function canSubmit() {
  return isNameValid() && isImageValid() && isCommandValid()
}

function isNameValid() {
  return !!name.match('^[a-zA-Z0-9]+$')
}

function isImageValid() {
  return !!image.match('^[a-zA-Z0-9./_:%]+$')
}

function isCommandValid() {
  return command !== ""
}

function onUserEvent(evt) {
  if (evt.Type == "change") {
    switch (evt.ElementID) {
      case "name":
        name = evt.Value
        if (nameValid !== isNameValid()) {
          nameValid = isNameValid()
          render()
        }
        break
      case "image":
        image = evt.Value
        if (imageValid !== isImageValid()) {
          imageValid = isImageValid()
          render()
        }
        break
      case "command":
        command = evt.Value
        if (commandValid !== isCommandValid()) {
          commandValid = isCommandValid()
          render()
        }
        break
      default:
        console.log("Unknown Change event", JSON.stringify(evt))
    }
    if (submitEnabled !== canSubmit()) {
       submitEnabled = canSubmit()
       render()
    }
  }

  if (evt.Type == "submit") {
    addTarget(name, image, command)
    setLocation("/#/")
  }

}

function unmount() {

}

function render() {
  var dm = addTargetDisplay.DeepCopy();
  dm.SetElementAttribute("submitButton", "disabled", !submitEnabled)
  dm.SetElementAttribute("submitButton","bsStyle",submitEnabled ? "success" : "danger")
  dm.SetElementAttribute("nameFormGroup","validationState", nameValid ? "success" : "error")
  dm.SetElementAttribute("imageFormGroup","validationState", imageValid ? "success" : "error")
  dm.SetElementAttribute("commandFormGroup","validationState", commandValid ? "success" : "error")
	updateScreen(withNavigation(dm));
}

"/add_target"
