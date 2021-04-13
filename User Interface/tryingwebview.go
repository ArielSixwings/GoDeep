package main

import "github.com/webview/webview"

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("GoDeep!")
	w.SetSize(1920, 1200, webview.HintNone)
	w.Navigate("file:///home/ariel/Documents/project/User%20Interface/godeepUI.html")
	w.Run()
}