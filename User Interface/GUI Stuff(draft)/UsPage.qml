import QtQuick 2.7
import QtQuick.Controls 2.12
import QtQuick.Window 2.2
import QtQuick.Controls.Styles 1.1
import QtQuick.Layouts 1.3
ApplicationWindow{
    id: us
    visible: true
    width: 800
    height: 600
Rectangle{
    id: retangle1
    color: "#483d8b"
    anchors.fill:parent
    Image{
        id: imageForm
        visible: true
        source: "background2.png"
        clip: false
        fillMode: Image.PreserveAspectFit
       //  transform: Rotation { origin.x: 50; origin.y: 25; angle: 45}
         width:  800
         height: 600
    }

        }
}   