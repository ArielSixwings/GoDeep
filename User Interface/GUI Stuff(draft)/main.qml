import QtQuick 2.7
import QtQuick.Controls 2.1
//import QtQuick 2.2
//import QtQuick.Layouts 1.3

ApplicationWindow{
	visible: true
	Image{	
		id: image
		source: "https://www.nasa.gov/sites/default/files/styles/full_width_feature/public/thumbnails/image/iss064e038775.jpg"
		fillMode: Image.PreserveAspectFit
		width: 200
		height: 200
		Rectangle {
			color: "red"
			x:0 ; y:950
			height: 25
		    width: 100 * image.progress
		    visible: image.progress != 1
		}
	}
}