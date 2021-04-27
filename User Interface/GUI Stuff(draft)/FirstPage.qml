import QtQuick 2.7
import QtQuick.Controls 2.12
import QtQuick.Window 2.2
import QtQuick.Controls.Styles 1.1
import QtQuick.Layouts 1.3
ApplicationWindow{
	id: home
    visible: true
    width: 1080
    height: 720
  		header: TabBar{
        id: bar
        TabButton{
            text: 'Home'
            font.pixelSize: 30
            icon.name: "IconHome"
            icon.source:"HomeIcon.png"
        }
        TabButton{
            text: 'Images'
            font.pixelSize: 30
            icon.name: "IconImage"
        icon.source: "Images.png"
        }
        TabButton{
            text:'Genetics'
            font.pixelSize: 30
            icon.name: "IconGenetic"
            icon.source: "GeneticIcon.png"
        }
        TabButton{
            text:'Statistics'
            font.pixelSize: 30
            icon.name: "StaticIcon"
            icon.source: "StaticIcon.png"
        }
        TabButton{
            text:'Help'
            font.pixelSize: 30
            icon.name: "HelpIcon"
            icon.source: "Help.png"
        }

    }
    StackLayout{
        anchors.fill:parent
        currentIndex: bar.currentIndex
        Item{
            id: paginaHome
            Rectangle{
                color: "#483d8b"
                anchors.fill:parent
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
            id: gopher
            x: 800
            y: 200
            width: 172
            height: 187
            visible: true
            source: "gopher.png"
            clip: false
            fillMode: Image.PreserveAspectFit
        }
        Button{
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
        load_page();
        }
}
        }
        Item{
            id: image
            Rectangle{
                color: 'black'
                anchors.fill:parent
            }
        }
        Item{
            id: genetic
            Rectangle{
                color:'black'
                anchors.fill:parent
            }
        }
        Item{
            id: statistic
            Rectangle{
                color:'black'
                anchors.fill:parent
            }
        }
        Item{
            id: help
            Rectangle{
                color:'black'
                anchors.fill:parent
            }
        }
    }
	}
