import QtQuick 2.9
import QtQuick.Controls 2.2
import QtQuick.Window 2.2
import QtQuick.Controls.Styles 1.1
import QtQuick.Layouts 1.3
//import "UsPage.qml"
ApplicationWindow{
	id: main
    //visible: true
    title: qsTr("Home")
    width: 1080
    height: 720
    minimumWidth :1080
    minimumHeight:720
    StackView {
        id: mystackview
        initialItem:first_Page
    }
    Component{
        id:first_Page
        FirstPage{}
    }
    Component{
        id:us_Page
        UsPage{}
    }
    function load_page(){
            mystackview.push(us_Page);
        }
}

//}

/*


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
    Image {
            id: image
            x: 800
            y: 200
            width: 172
            height: 187
            visible: true
            source: "gopher.png"
            clip: false
            fillMode: Image.PreserveAspectFit
        }
    /* Image {
        id: image2
        x: 8
        y: 8
        width: 71
        height: 57
        source: "menu_3.0.png"
        fillMode: Image.PreserveAspectFit
    }*/
    
    /*Button{
        id: aboutUs
    	//icon.width:150
    	//icon.height:30
    	x: 400
    	y : 400
        //width: 150
        //height: 70
        text:qsTr("About Us")
        font.pixelSize: 50
        palette {
        button: "transparent"
    	}
    	icon.name: "IconAbout"
    	icon.source: "IconAbout.png"
    	//icon.x:200
        onClicked: {
               
        }
}

        Button{
        id: openMenuButton
    	icon.name: "IconButton"
    	icon.source: "menu_3.0.png"
    	icon.width:50
    	icon.height:60
    	icon.color: "black"
        width: 250
        height: 70
        palette {
        button: "transparent"
    	}
        onClicked: {
               menu.open();
        }
}

    /*Label{
        id :label
        anchors.centerIn:parent
        text:'Selecione uma opção do Menu'
        font.pointSize: 30
    }*/

/*
    Menu {
    	width: 250
        height: 200

        id:menu
        y: openMenuButton.height
           MenuSeparator {
        contentItem: Rectangle {
            implicitWidth: 200
            implicitHeight: 1
            color: "#483d8b"
        }
    }
        MenuItem{
        icon.name: "ImagesIcon"
    	icon.source: "images.png"
       // width: 73
        //height: 60

         text:'Images'
         font.pixelSize: 30
        /*Images{
        id: image3
        x: 0
        y: 0
        width: 71
        height: 57
        source: "Images.png"
            }*/

/*
            onClicked: {
          //  label.text='OPÇÃO 1'
            }
        }
        MenuItem{
       // width: 73
       // height: 60

       icon.name: "GeneticIcon"
       icon.source: "GeneticIcon.png"
            text:'Genetic'
            font.pixelSize: 30
            onClicked: {
            //label.text= 'OPÇÃO 2';
            }
        }
        MenuItem{
      //  width: 73
     //   height: 60
            text:'Static'
            font.pixelSize: 30
            icon.name: "StaticIcon"
       		icon.source: "StaticIcon.png"
            onClicked: {
            //label.text='OPÇÃO 3';
            }
        }
       MenuItem{
    //   width: 73
  //     height: 60
           text:'Sair'
           font.pixelSize: 30
           onClicked: {
               Qt.quit();
           }
       }
       delegate: MenuItem {
        id: menuItem
        implicitWidth: 200
        implicitHeight: 40

        arrow: Canvas {
            x: parent.width - width
            implicitWidth: 40
            implicitHeight: 40
            visible: menuItem.subMenu
            onPaint: {
                var ctx = getContext("2d")
                ctx.fillStyle = menuItem.highlighted ? "#ffffff" : "#21be2b"
                ctx.moveTo(15, 15)
                ctx.lineTo(width - 15, height / 2)
                ctx.lineTo(15, height - 15)
                ctx.closePath()
                ctx.fill()
            }
        }

        indicator: Item {
            implicitWidth: 40
            implicitHeight: 40
            Rectangle {
                width: 26
                height: 26
                anchors.centerIn: parent
                visible: menuItem.checkable
                border.color: "#21be2b"
                radius: 3
                Rectangle {
                    width: 14
                    height: 14
                    anchors.centerIn: parent
                    visible: menuItem.checked
                    color: "#21be2b"
                    radius: 2
                }
            }
        }

        contentItem: Text {
            leftPadding: menuItem.indicator.width
            rightPadding: menuItem.arrow.width
            text: menuItem.text
            font: menuItem.font
            opacity: enabled ? 1.0 : 0.3
            color: menuItem.highlighted ? "#ffffff" : "#21be2b"
            horizontalAlignment: Text.AlignLeft
            verticalAlignment: Text.AlignVCenter
            elide: Text.ElideRight
        }

        background: Rectangle {
            implicitWidth: 200
            implicitHeight: 40
            opacity: enabled ? 1 : 0.3
            color: menuItem.highlighted ? "#21be2b" : "transparent"
        }
    }

    background: Rectangle {
        implicitWidth: 200
        implicitHeight: 40
        color: "#ffffff"
        border.color: "#483d8b"
        radius: 2
    }
    }
}*/