package clock

import (
	"html/template"
	"log"
	
	"github.com/jfyne/live"
)

func NewHandler() *live.BaseHandler {
	handler := live.NewHandler(withRenderConfig())
	handler.HandleMount(onMount)
	handler.HandleSelf(tick, onTick)
	
	return handler
}

func withRenderConfig() live.HandlerConfig {
	t, err := template.ParseFiles("root.html", "pkg/apps/clock/view.html")
	if err != nil {
		log.Fatal(err)
	}
	
	return live.WithTemplateRenderer(t)
}
