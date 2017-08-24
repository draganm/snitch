tx(function(w) {
  if (!w.Exists(dbpath("targets"))) {
    w.CreateMap(dbpath("targets"))
  }

  if (!w.Exists(dbpath("status"))) {
    w.CreateData(dbpath("status"), JSON.stringify([]))
  }
});
