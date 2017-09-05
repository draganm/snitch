package target

//go:generate kickback-generator -p target

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	reactor "github.com/draganm/go-reactor"
	"github.com/draganm/immersadb/dbpath"
	"github.com/draganm/immersadb/modifier"
	"github.com/draganm/kickback"
	"github.com/draganm/snitch/executor"
	"github.com/draganm/snitch/tx"
	. "github.com/draganm/snitch/ui/navigation"
)

func init() {
	kickback.AddScreen("/targets/:targetID", func(ctx *kickback.Context) {

		id := ctx.ScreenContext.Params["targetID"]
		alerts := []AlertLine{}
		config := &executor.Config{}
		showDeleteConfirm := false

		var render = func() {

			spew.Dump(config)

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
