import QtQuick 2.7
import QtQuick.Controls 2.1
//import QtQuick 2.2
//import QtQuick.Layouts 1.3

ApplicationWindow{
	visible: true
	width: 300 
	height: 115
	flags: Qt.FramelessWindowHint

	TextInput{
		id: textElement
		x: 50 
		y: 25
		text: "GoDeep"
		font.family: "Helvetica"; font.pixelSize: 50	
	}

	// Rectangle {
	// 	x: 50; y: 75; height: 5
	// 	width: textElement.width
	// 	color: "green"
	// }
}