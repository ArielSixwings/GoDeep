import QtQuick 2.11
import QtQuick.Controls 2.4
import QtQuick.Window 2.2
import QtQuick.Dialogs 1.2
import QtQuick.Controls.Styles 1.1
import QtQuick.Layouts 1.3

ApplicationWindow{
	id: main

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

	function load_page(page){
		switch (page){
			case 'page 1':
			mystackview.push(us_Page);
			break;
		}
	}
}