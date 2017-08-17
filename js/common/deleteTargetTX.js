function deleteTarget(targetUUID) {

  tx(function(w) {
    w.Delete(dbpath("targets", targetUUID))
    var status = JSON.parse(w.Data(dbpath("status")))
    var targetIndex = -1
    for (var i = 0; i<status.length; i++) {
      if (status[i].id === targetUUID) {
        targetIndex = i
        break
      }
    }
    if (targetIndex >= 0) {
      status.splice(targetIndex,1)
    }
    w.CreateData(dbpath("status"), JSON.stringify(status))
  })
}
