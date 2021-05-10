package main

import (
	// "fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	//"github.com/therecipe/qt/quick"
)

func soma()int{
	var soma int
	soma=10+10
	return soma
}





func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var app = qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("./qml/UserInterface.qml", 0))


	gui.QGuiApplication_Exec()

}