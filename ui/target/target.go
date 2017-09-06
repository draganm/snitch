package target

//go:generate kickback-generator -p target

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	reactor "github.com/draganm/go-reactor"
	"github.com/draganm/immersadb/dbpath"
	"github.com/draganm/immersadb/modifier"
	"github.com/draganm/kickback"
	"github.com/draganm/snitch/executor"
	"github.com/draganm/snitch/tx"
	. "github.com/draganm/snitch/ui/navigation"
)

type logEntry struct {
	Time   time.Time              `json:"time"`
	Fields map[string]interface{} `json:"fields"`
	Level  string                 `json:"level"`
}

func levelStyle(status string) string {
	switch status {
	case "info":
		return "info"
	case "error":
		return "danger"
	case "failure":
		return "danger"
	case "success":
		return "success"
	default:
		return "default"
	}
}

func init() {
	kickback.AddScreen("/targets/:targetID", func(ctx *kickback.Context) {

		id := ctx.ScreenContext.Params["targetID"]
		alerts := []AlertLine{}
		config := &executor.Config{}
		showDeleteConfirm := false
		logEntries := []logEntry{}

		var render = func() {

			dm := display.DeepCopy()
			dm.SetElementAttribute("panel", "header", fmt.Sprintf("Target %s (%s)", config.Name, id))
			dm.SetElementText("image", config.Image)
			dm.SetElementText("command", config.Command)
			dm.SetElementText("interval", fmt.Sprintf("%d", config.Interval))

			if showDeleteConfirm {
				var mod = confirmDeleteModal.DeepCopy()
				mod.SetElementText("targetName", config.Name)
				dm.ReplaceChild("deleteButton", mod)
			}

			for _, lent := range logEntries {
				le := logEvent.DeepCopy()
				// var d = new Date(log[i].time)
				le.SetElementAttribute("rowPanel", "header", lent.Time.Format(time.RFC850))
				le.SetElementAttribute("rowPanel", "bsStyle", levelStyle(lent.Level))
				for k, v := range lent.Fields {
					lep := logEventProperty.DeepCopy()
					lep.SetElementText("name", k)
					lep.SetElementText("value", fmt.Sprintf("%v", v))
					le.AppendChild("rowPanel", lep)
				}
				// for (var key in log[i].fields) {
				//   var lep = logEventProperty.DeepCopy()
				//   lep.SetElementText("name", key)
				//   lep.SetElementText("value", log[i].fields[key])
				//   le.AppendChild("rowPanel",lep)
				// }
				dm.AppendChild("mainGrid", le)
			}

			ctx.ScreenContext.UpdateScreen(&reactor.DisplayUpdate{
				Model: WithNavigation(dm, nil),
			})
		}

		ctx.Listen(dbpath.New("targets", id), func(er modifier.EntityReader) {
			if er == nil {
				alerts = []AlertLine{
					AlertLine{Text: "Target does not exist!", Typ: "danger"},
				}
				render()
				return
			}

			err := config.Read(er.EntityReaderFor(dbpath.New("config")))
			if err != nil {
				alerts = []AlertLine{
					AlertLine{Text: err.Error(), Typ: "danger"},
				}
				render()
				return
			}

			logEntries = nil
			err = er.EntityReaderFor(dbpath.New("log")).ForEachArrayElement(func(index uint64, logEr modifier.EntityReader) error {
				le := logEntry{}
				e := json.NewDecoder(logEr.Data()).Decode(&le)
				if e != nil {
					return e
				}
				logEntries = append(logEntries, le)
				return nil
			})

			if err != nil {
				alerts = []AlertLine{
					AlertLine{Text: err.Error(), Typ: "danger"},
				}
				render()
				return
			}

			render()

			//   r.ForEach(dbpath("log"), function(i, lr){
			//     log.push(JSON.parse(lr.Data(dbpath())))
			//   })
		})

		ctx.OnUserEventFunc = func(evt *reactor.UserEvent) {

			switch evt.ElementID {
			case "deleteButton":
				showDeleteConfirm = true
				render()
			case "deleteCancelButton":
				showDeleteConfirm = false
				render()
			case "deleteConfirmButton":
				err := tx.TX{ctx.DB}.DeleteTarget(id)
				if err != nil {
					log.Println(err)
					return
				}
				ctx.ScreenContext.UpdateScreen(&reactor.DisplayUpdate{
					Location: "#/",
				})
			}

		}

	})
}
