import QtQuick 2.11
import QtQuick.Controls 2.4
import QtQuick.Window 2.2
import QtQuick.Dialogs 1.2
import QtQuick.Controls.Styles 1.1
import QtQuick.Layouts 1.3
import "../src/extractstrategy" as Extractor

ApplicationWindow{
	id: home
		visible: true
		width: 1080
		height: 720

		FileDialog{
			id: openFileDialog
			title:"Open file"
			folder: shortcuts.home
			selectExisting: true
			selectFolder: true
			nameFilters: ["All files (*)"]
			onAccepted:{ 
				fileNameTextField.text=fileUrl 
				console.log(fileNameTextField.text) 
				home.sendBool(true)
			}
			onRejected:{ 
				fileNameTextField.text="CANCELED" 
				console.log("Invalid file")
				home.sendBool(false)
			}
		}

		header: TabBar{
			id: bar
			TabButton{
				text: 'Home'
				font.pixelSize: 30
				icon.name: "IconHome"
				icon.source:"./ImagesUI/HomeIcon.png"
			}
			TabButton{
				text: 'Images'
				font.pixelSize: 30
				icon.name: "IconImage"
				icon.source: "./ImagesUI/Images.png"
			}
			TabButton{
				text:'Genetics'
				font.pixelSize: 30
				icon.name: "IconGenetic"
				icon.source: "./ImagesUI/GeneticIcon.png"
			}
			TabButton{
				text:'Statistics'
				font.pixelSize: 30
				icon.name: "StaticIcon"
				icon.source: "./ImagesUI/StaticIcon.png"
			}
			TabButton{
				text:'Help'
				font.pixelSize: 30
				icon.name: "HelpIcon"
				icon.source: "./ImagesUI/Help.png"
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
					source: "./ImagesUI/gopher.png"
					clip: false
					fillMode: Image.PreserveAspectFit
				}
				Button{
					id: aboutUs
					x: 400
					y : 400
					text:qsTr("About Us")
					font.pixelSize: 50
					palette { button: "transparent" }
					icon.name: "IconAbout"
					icon.source: "./ImagesUI/IconAbout.png"
					onClicked: { load_page("page 1"); }
				}
			}
			Item{
				id: image
				Rectangle{
					color:'gray'
					anchors.fill:parent
					Rectangle {
						id: rectangleImage2
						x: 8
						y: -214
						width: 390
						height: 770
						color: "purple"
						border.color: "purple"
						transform: Rotation { origin.x: 25; origin.y: 25; angle: 30}
					}
					Text{
						id:text1Image
						x:80
						y:85
						width:100
						height:100
						font.pixelSize:30
						color:"black"
						text:' Data set origin:'  
					}
					Rectangle {
						id:rectangle3Image
						x:80+13
						y:85+80
						width: 540-50
						height: 55
						color: "transparent"
						border.color: "black"
						border.width:6
					}
					Text{
						id: text2Images
						x:540+120
						y:85
						width:100
						height:100
						text:'Preview'
						color:"black"
						font.pixelSize:30
					}
					Rectangle{
						id:rectangle4Image
						x:540+120
						y:85+80
						width:400
						height:400
						color:'transparent'
						border.color:'black'
						border.width:8
					}
					Button{
						id:buttonPreviousImages
						x:540+125
						y:85+80+400+20
						width:40
						height:40
						icon.name: "IconPrevious"
						icon.source:"./ImagesUI/ButtonPrevious.png"
						palette{ button:"transparent" }
					}
					Button{
						id:buttonNextImages
						x:540+125+60
						y:85+80+400+20
						width:40
						height:40
						icon.name: "IconNext"
						icon.source:"./ImagesUI/ButtonNext.png"
						palette{ button:"transparent" }
					}
					Button{
						id:buttonAddImage
						text: 'Add File'
						icon.name: "IconAdd"
						icon.source:"./ImagesUI/IconAdd.png"
						icon.color:"black"
						x:900
						y:85*2+35
						width: 100
						height: 40
						onClicked:{ openFileDialog.open();}
					}
				}
			}
			Item{
				id: genetic
				Rectangle{
					color:'gray'
					anchors.fill:parent
					Rectangle {
						id: rectangle2
						x: 8
						y: -214
						width: 390
						height: 770
						color: "green"
						border.color: "green"
						transform: Rotation { origin.x: 25; origin.y: 25; angle: 30}
					}
					Button{
						id:buttonConfirmar1
						icon.name: "IconConfirm"
						icon.source:"./ImagesUI/IconConfirm.png"
						icon.color:"black"
						
						palette{ button:"transparent" }
						x:840
						y:85*2+35
						width: 40
						height: 40
					}
					Button{
						id:buttonAdd1
						text: 'Add File'
						icon.name: "IconAdd"
						icon.source:"./ImagesUI/IconAdd.png"
						icon.color:"black"
						x:900
						y:85*2+35
						width: 100
						height: 40
						onClicked:{ openFileDialog.open(); }
					}
					Rectangle{
						id:rectangleGenetic1
						color:'transparent'
						border.color: "black"
						border.width: 6
						x:40
						y:85
						width: 840
						height: 90
						Text {
							id:textGenetic
							y:24
							width: 588
							height: 108
							font.pixelSize: 30
							text:qsTr("    Insert tape or data set of tapes from fathers")
						}
					}
					TextField {
						id:textFielGenetic1
						x:40
						y:85*2+25
						width: 760
						height: 65
						font.pixelSize: 30
						placeholderText: qsTr("		Enter Sequence:")
					}
					TextField{
						id:fileNameTextField
						x:850
						y:85*2+35+50
						width: 150
						height: 40
					}
					Button{
						id:buttonConfirmar2
						icon.name: "IconConfirm2"
						icon.source:"./ImagesUI/IconConfirm.png"
						icon.color:"black"
						palette{ button:"transparent" }
						x:840
						y:320+90+35
						width: 40
						height: 40
					}
					Button{
						id:buttonAdd2
						text: 'Add File'
						icon.name: "IconAdd2"
						icon.source:"./ImagesUI/IconAdd.png"
						icon.color:"black"
						x:900
						y:320+90+35
						width: 100
						height: 40
					}
					Rectangle{
						id:rectangleGenetic2
						color:'transparent'
						border.color: "black"
						border.width: 6
						x:40
						y:320
						width: 840
						height: 90
						Text {
							id:textGenetic2
							y:24
							width: 588
							height: 108
							font.pixelSize: 30
							text:qsTr("     Insert tape or data set of tapes from Childs")
						}
					}
					TextField {
						id:textFielGenetic2
						x:40
						y:320+90+25
						width: 760
						height: 65
						font.pixelSize: 30
						placeholderText: qsTr("		Enter Sequence:")
					}
					Button{
						id:verifySequence
						text: 'Verify Sequence'
						font.pixelSize: 25
						x:40
						y:320+90+25+120
						width: 230
						height: 75
					}
					Button{
						id:analyzeResult
						text: 'Result analysis'
						font.pixelSize: 25
						x:40+420
						y:320+90+25+120
						width: 230
						height: 75
					}
				}
			}
			Item{
				id: statistic
				Rectangle{
					color:'gray'
					anchors.fill:parent
					Rectangle {
						id: rectangle2Statistic
						x: 8
						y: -214
						width: 390
						height: 770
						color: "#057fb6"
						border.color: "#057fb6"
						transform: Rotation { origin.x: 25; origin.y: 25; angle: 30}
					}
					Text{
						id:text1Statistic
						x:80
						y:85
						width:100
						height:100
						font.pixelSize:30
						color:"black"
						text:' Data set origin:'	
					}
					Rectangle {
						id:rectangle3Statistic
						x:80+13
						y:85+80
						width: 540-50
						height: 55
						color: "transparent"
						border.color: "black"
						border.width:6
					}
					Button{
						id:buttonAddDirStatic
						x:80+250
						y:85
						width:150
						height:55
						text:'Add file:'
						font.pixelSize:25
					}
					Text{
						id: text2Statistic
						x:540+120
						y:85
						width:100
						height:100
						text:'Preview'
						color:"black"
						font.pixelSize:30
					}
					Rectangle{
						id:rectangle4Statistic
						x:540+120
						y:85+80
						width:400
						height:400
						color:'transparent'
						border.color:'black'
						border.width:8
					}
					Button{
						id:button2Statistic
						x:80+13+80
						y:450+85+40
						text:'Analyse results'
						font.pixelSize:30
						width:220
						height:55
					}
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
