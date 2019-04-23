package bamboohr

import (
	"fmt"

	"github.com/wtfutil/wtf/wtf"
)

const APIURI = "https://api.bamboohr.com/api/gateway.php"

type Widget struct {
	wtf.TextWidget

	settings *Settings
}

func NewWidget(refreshChan chan<- string, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(refreshChan, settings.common, false),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	client := NewClient(
		APIURI,
		widget.settings.apiKey,
		widget.settings.subdomain,
	)

	todayItems := client.Away(
		"timeOff",
		wtf.Now().Format(wtf.DateFormat),
		wtf.Now().Format(wtf.DateFormat),
	)

	widget.View.SetTitle(widget.ContextualTitle(widget.Name()))

	widget.View.SetText(widget.contentFrom(todayItems))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(items []Item) string {
	if len(items) == 0 {
		return fmt.Sprintf("\n\n\n\n\n\n\n\n%s", wtf.CenterText("[grey]no one[white]", 50))
	}

	str := ""
	for _, item := range items {
		str = str + widget.format(item)
	}

	return str
}

func (widget *Widget) format(item Item) string {
	var str string

	if item.IsOneDay() {
		str = fmt.Sprintf(" [green]%s[white]\n %s\n\n", item.Name(), item.PrettyEnd())
	} else {
		str = fmt.Sprintf(" [green]%s[white]\n %s - %s\n\n", item.Name(), item.PrettyStart(), item.PrettyEnd())
	}

	return str
}
