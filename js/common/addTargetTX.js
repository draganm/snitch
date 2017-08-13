function addTarget(name, image, command) {
  console.log("done")
  var targetUUID = uuidv4()
  console.log("done")

  tx(function(w) {
    w.CreateMap(dbpath("targets", targetUUID))

    w.CreateData(
      dbpath("targets", targetUUID,"config"),
      JSON.stringify(
        {
          name: name,
          image: image,
          command: command,
        }
      )
    )

    w.CreateData(
      dbpath("targets", targetUUID,"status"),
      JSON.stringify(
        {
          lastUpdate: newDate(),
          status: "unknown",
        }
      )
    )

    w.CreateArray(dbpath("targets", targetUUID, "log"))
    var targets = JSON.parse(w.Data(dbpath("targetsOrdered")))
    targets.push(targetUUID)
    w.CreateData(dbpath("targetsOrdered"), JSON.stringify(targets))
  })
}
