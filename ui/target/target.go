package target

//go:generate kickback-generator -p target

import (
	"fmt"

	reactor "github.com/draganm/go-reactor"
	"github.com/draganm/immersadb/dbpath"
	"github.com/draganm/immersadb/modifier"
	"github.com/draganm/kickback"
	"github.com/draganm/snitch/executor"
	. "github.com/draganm/snitch/ui/navigation"
)

func init() {
	kickback.AddScreen("/targets/:targetID", func(ctx *kickback.Context) {

		id := ctx.ScreenContext.Params["targetID"]
		alerts := []AlertLine{}
		config := &executor.Config{}

		var render = func() {
			dm := display.DeepCopy()
			dm.SetElementAttribute("panel", "header", fmt.Sprintf("Target %s (%s)", config.Name, id))
			dm.SetElementText("image", config.Image)
			dm.SetElementText("command", config.Command)
			dm.SetElementText("interval", fmt.Sprintf("%d", config.Interval))
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

	})
}
