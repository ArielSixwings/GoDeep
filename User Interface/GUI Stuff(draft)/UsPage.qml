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

        /*Rectangle {
            id: rectangle1
            x: 40
            y: 88
            width: 588
            height: 108
            color: "#ffffff"
        }
    }
    
    Text {
        id:text
        x:40
        y:88
        width: 588
        height: 108
        font.pixelSize: 30
    text:qsTr("Insert tape or data set of tapes from fathers")

}
}
    /*Rectangle{
    id: retangle1
    color: "#483d8b"
    anchors.fill:parent
    /*Image{
        id: imageForm
        visible: true
        source: "background2.png"
        clip: false
        fillMode: Image.PreserveAspectFit
       //  transform: Rotation { origin.x: 50; origin.y: 25; angle: 45}
         width:  800
         height: 600
    }
    }*/
}   
}