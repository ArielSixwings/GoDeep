import QtQuick 2.11
import QtQuick.Controls 2.4
import QtQuick.Window 2.2
import QtQuick.Dialogs 1.2
import QtQuick.Controls.Styles 1.1
import QtQuick.Layouts 1.3

Item {
    property bool storagePermissionGranted
    Component.onCompleted: {
        // Check if storage permission is granted (only required for Android)
        storagePermissionGranted = Permission.checkPermission(Permission.PermissionTypeStorage) === Permission.PermissionResultGranted
    }
    PermissionDialog {
        id: permissionDialog
        openSettingsWhenDenied: true
        permission: PermissionDialog.PermissionDialogTypeStorage
        onAccepted: {
            documentDialog.open()
            console.log("Permission accepted")
        }
        onRejected: {
            console.log("Permission rejected")
        }
    }
    Button {
        text: "Select document"
        onClicked: {
            if (documentDialog.supported) {
                if (storagePermissionGranted === false){
                    permissionDialog.open()
                }
                else (storagePermissionGranted === true){
                    documentDialog.open()
                }
            }
            else {
                console.log("Not Supported")
            }
        }
    }
        DocumentDialog {
                id: documentDialog
                onAccepted: {
                        console.log("selected file URL " , fileUrl)
                }
                onRejected: {
                        if (status == DocumentDialog.DocumentDialogCancelledByUser) {
                                // Cancelled By User
                        }
                        if (status == DocumentDialog.DocumentDialogPermissionDenied) {
                                // Permission Denied
                        }
                        if (status == DocumentDialog.DocumentDialogNotSupported) {
                                // Not Supported
                        }
                        if (status == DocumentDialog.DocumentDialogFileReadError) {
                                // File Read Error
                        }
                }
        }
}