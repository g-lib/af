package main

import (
	"embed"
	"fmt"
	"net/http"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

//go:embed asciiflow2/*
var asciiflow2 embed.FS

func main() {
	onExit := func() {
		fmt.Println(time.Now().String(), "AF退出")
	}

	systray.Run(onReady, onExit)
}

const (
	Title   = "ASCIIFlow"
	ToolTip = "ASCIIFlow:纯文本流程图绘制工具"
	Port    = ":3000"
	AF2Path = "/asciiflow2/"
	AF3Path = "/asciiflow3/"
)

func onReady() {
	systray.SetTemplateIcon(Data, Data)
	systray.SetTitle(Title)
	systray.SetTooltip(ToolTip)
	go func() {
		http.Handle(AF2Path, http.FileServer(http.FS(asciiflow2)))
		http.Handle(AF3Path, http.FileServer(http.FS(asciiflow2)))
		fmt.Println(http.ListenAndServe(Port, nil))

	}()

	// We can manipulate the systray in other goroutines
	go func() {
		mOpen2 := systray.AddMenuItem("👀打开ASCIIFlow2", "ASCIIFlow2操作界面")
		mOpen3 := systray.AddMenuItem("👀打开ASCIIFlow2", "ASCIIFlow2操作界面")
		mOpen3.Disable()
		systray.AddMenuItem(fmt.Sprintf("Server:http://127.0.0.1%s", Port), "点击拷贝").Disable()
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("退出", "退出整个应用")
		beeep.Notify(Title, "请点击"+Title+"图标打开相应的版本", "")
		for {
			select {

			case <-mOpen2.ClickedCh:
				open.Run("http://127.0.0.1" + Port + AF2Path)
				beeep.Notify(Title, "请在浏览器中操作"+Title+"2", "")
			case <-mOpen3.ClickedCh:
				open.Run("http://127.0.0.1" + Port + AF3Path)
				beeep.Notify(Title, "请在浏览器中操作"+Title+"3", "")
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}
