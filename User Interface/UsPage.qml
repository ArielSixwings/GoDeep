import QtQuick 2.11
import QtQuick.Controls 2.4
import QtQuick.Window 2.2
import QtQuick.Dialogs 1.2
import QtQuick.Controls.Styles 1.1
import QtQuick.Layouts 1.3
ApplicationWindow{
	id: us
	visible: true
	width: 800
	height: 600
	 Rectangle {
		id: rectangle
		x: 0
		y: 0
		width: 800
		height: 600
		anchors.fill:parent
		color: "#483d8b"
		border.color: "#483d8b"
	}
}