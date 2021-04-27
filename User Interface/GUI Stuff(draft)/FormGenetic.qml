import QtQuick 2.7
import QtQuick.Controls 2.12
import QtQuick.Window 2.2

ApplicationWindow{
	title:'About us'
	id: home
    visible: true
    width: 1080
    height: 720
    Rectangle {
        id: rectangle
        x: 0
        y: 0
        width: 1082
        height: 722
        color: "#483d8b"

        Text {
            id: text1
            x: 330
            y: 200
            //width: 57
            //height: 32
            text: qsTr("GoDeep")
            font.pixelSize: 120
        }
    }
}