console.log("setup done!")

function withNavigation(dm) {
  var navCopy = navigation.DeepCopy();
  navCopy.ReplaceChild("content", dm);
  return navCopy;
}


tx(function(w) {
  if (!w.Exists(dbpath("targets"))) {
    w.CreateMap(dbpath("targets"))
  }
  if (!w.Exists(dbpath("targetsOrdered"))) {
    w.CreateMap(dbpath("targetsOrdered"))
  }
});
