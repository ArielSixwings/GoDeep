package main

import (
	// "fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	//"github.com/therecipe/qt/quick"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var app = qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("UserInterface.qml", 0))

	// var (
	// 	key string
	// 	namefolder string
	// )
	gui.QGuiApplication_Exec()

	// for ; true;  {
	// 	fmt.Scanf("%s",key)
	// 	fmt.Println(key)
	// 	if key == "23399784##27"{
	// 		fmt.Scanf("%s", &namefolder)
	// 		fmt.Println(namefolder)
	// 		break
	// 	} 
		
	// }

}