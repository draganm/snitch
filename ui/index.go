package ui

import (
	"log"

	reactor "github.com/draganm/go-reactor"
	"github.com/draganm/immersadb/dbpath"
	"github.com/draganm/immersadb/modifier"
	"github.com/draganm/kickback"
	"github.com/draganm/snitch/executor"
)

func init() {

	var statusStyle = func(status string) string {
		switch status {
		case "unknown":
			return "warning"
		case "running":
			return "info"
		case "success":
			return "success"
		case "failed":
			return "danger"
		}
		return "default"
	}

	kickback.AddScreen("/", func(ctx *kickback.Context) {

		status := executor.StatusList{}

		ctx.Listen(dbpath.New("status"), func(er modifier.EntityReader) {
			if er == nil {
				status = nil
				return
			}
			err := (&status).Read(er)
			if err != nil {
				// TODO use alert
				log.Println(err)
			}
			log.Println("status", status)
		})

		var render = func() {

			d := indexDisplay.DeepCopy()

			if len(status) > 0 {
				d.DeleteChild("noTargets")
			}

			for _, s := range status {
				it := indexTarget.DeepCopy()
				it.SetElementAttribute("item", "bsStyle", statusStyle(s.Status))
				it.SetElementAttribute("item", "href", "#/targets/"+s.ID)
				it.SetElementText("item", s.Name)
				d.AppendChild("targets", it)
			}

			ctx.ScreenContext.UpdateScreen(&reactor.DisplayUpdate{
				Model: withNavigation(d, nil),
			})
		}

		render()

	})
}
