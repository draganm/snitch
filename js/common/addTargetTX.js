function addTarget(name, image, command, interval) {
  var targetUUID = uuidv4()

  tx(function(w) {
    w.CreateMap(dbpath("targets", targetUUID))

    w.CreateData(
      dbpath("targets", targetUUID, "config"),
      JSON.stringify(
        {
          name: name,
          image: image,
          command: command,
          interval: interval,
        }
      )
    )

    w.CreateArray(dbpath("targets", targetUUID, "log"))

    var targets = JSON.parse(w.Data(dbpath("status")))
    targets.push({name:name, id: targetUUID, status: "unknown"})
    w.CreateData(dbpath("status"), JSON.stringify(targets))
  })
}
