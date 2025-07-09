package main

import (
	_case "P_channel/case"
	"os"
	"os/signal"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//_case.Communication()
	//_case.ConcurrentSync()
	_case.NotifyAndMultiplex()
	ch := make(chan os.Signal, 0)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
	//让chan阻塞，等待信号
	//主协程

}
