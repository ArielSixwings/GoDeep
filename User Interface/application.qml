/****************************************************************************
**
** Copyright (C) 2015 The Qt Company Ltd.
** Contact: http://www.qt.io/licensing/
**
** This file is part of the examples of the Qt Toolkit.
**
** $QT_BEGIN_LICENSE:BSD$
** You may use this file under the terms of the BSD license as follows:
**
** "Redistribution and use in source and binary forms, with or without
** modification, are permitted provided that the following conditions are
** met:
**   * Redistributions of source code must retain the above copyright
**     notice, this list of conditions and the following disclaimer.
**   * Redistributions in binary form must reproduce the above copyright
**     notice, this list of conditions and the following disclaimer in
**     the documentation and/or other materials provided with the
**     distribution.
**   * Neither the name of The Qt Company Ltd nor the names of its
**     contributors may be used to endorse or promote products derived
**     from this software without specific prior written permission.
**
**
** THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
** "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
** LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
** A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
** OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
** SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
** LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
** DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
** THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
** (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
** OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE."
**
** $QT_END_LICENSE$
**
****************************************************************************/

//Slightly edited the original code for a scrollable TextArea and Qt Quick 2 controls

import QtQuick 2.2
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.3

ApplicationWindow {
	visible: true
	title: "GoDeep"
	property int margin: 11
	minimumWidth: 600
	minimumHeight: 450

	ColumnLayout {
		id: mainLayout
		anchors.fill: parent
		anchors.margins: margin
		GroupBox {
			id: rowBox
			title: "GoDeep Graphic Interface"
			Layout.fillWidth: true

			RowLayout {
				id: rowLayout
				anchors.fill: parent
				TextField {
					placeholderText: "testing editing stuff"
					Layout.fillWidth: true
				}
				Button {
					text: "Choose Application"
					onClicked: { 
						stackLayout.advance() 
					} 
				}
			}
		}

		GroupBox {
		  id: gridBox
		  title: "Select files to read"
		  Layout.fillWidth: true

		  GridLayout {
			  id: gridLayout
			  rows: 3
			  flow: GridLayout.TopToBottom
			  anchors.fill: parent

			  Label { text: "File 1" }
			  Label { text: "File 2" }
			  Label { text: "File 3" }

			  TextField { }
			  TextField { id: textField }
			  TextField { }

			  Flickable {
				  anchors {
					  top: parent.top
					  left: textField.right
					  right: parent.right
					  bottom: parent.bottom
				  }

				  contentHeight: textid.width
				  contentWidth: textid.height

				  TextArea.flickable: TextArea {
					id: textid
					text: "Those files will be the base to all futher process.\n"
						+ "If the name of the file is significative, you can use the 'get label' option to set the labels."
					wrapMode: TextArea.Wrap
				  }

				  ScrollBar.vertical: ScrollBar { }
			  }
		  }
		}
		TextArea {
			id: t3
			text: "This fills the whole cell"
			Layout.minimumHeight: 30
			Layout.fillHeight: true
			Layout.fillWidth: true
		}
		GroupBox {
			id: stackBox
			title: "Stack layout"
			implicitWidth: 200
			implicitHeight: 60
			Layout.fillWidth: true
			Layout.fillHeight: true
			StackLayout {
				id: stackLayout
				anchors.fill: parent

				function advance() { currentIndex = (currentIndex + 1) % count }

				Repeater {
					id: stackRepeater
					model: 5
					Rectangle {
						color: Qt.hsla((0.5 + index)/stackRepeater.count, 0.3, 0.7, 1)
						Button { 
							anchors.centerIn: parent; text: "Application " + (index + 1); 
							onClicked: { 
								stackLayout.advance() 
							} 
						}
					}
				}
			}
		}
	}
}