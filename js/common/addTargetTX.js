function addTarget(name, image, command) {
  var targetUUID = uuidv4()

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
    w.CreateArray(dbpath("targets", targetUUID, log))
  })
}
