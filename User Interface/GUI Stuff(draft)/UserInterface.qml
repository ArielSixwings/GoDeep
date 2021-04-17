import QtQuick 2.7
import QtQuick.Controls 2.1

ApplicationWindow{
	id: root
	visible: true
	
	height: 400
	width: height
	color: "gray"
	Text{
		id: title
		x: 100 ; y: 20
		text: "GoDeep"
		font.family: "Helvetica"; font.pixelSize: 50	
	}
	Rectangle {
		x:75 ; y:25
		
		width: 200
		height: width/2
		
		visible: true
		
		color: "blue"
	}
	// Rectangle {
	// 	x:75 ; y:125
		
	// 	width: 200
	// 	height: width/2
		
	// 	visible: true
		
	// 	color: "green"
	// }
	// Rectangle {
	// 	x:75 ; y:225
		
	// 	width: 200
	// 	height: width/2
		
	// 	visible: true
		
	// 	color: "purple"
	// }
}