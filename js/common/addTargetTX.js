function addTarget(name, image, command, interval) {
  console.log("done")
  var targetUUID = uuidv4()
  console.log("done")

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

    var lastUpdate = new Date()

    w.CreateData(
      dbpath("targets", targetUUID,"status"),
      JSON.stringify(
        {
          lastUpdate: lastUpdate,
          status: "unknown",
        }
      )
    )

    w.CreateArray(dbpath("targets", targetUUID, "log"))


    var targets = JSON.parse(w.Data(dbpath("status")))
    targets.push({name:name, id: targetUUID, status: "unknown", lastUpdate: lastUpdate})
    w.CreateData(dbpath("status"), JSON.stringify(targets))
  })
}
