#include "qmlintegration.h"
#include <QDebug>

qmlintegration::qmlintegration(QObject *parent) : QObject(parent){

}
void qmlintegration::callFromQml(){
	qDebug("Recieved call from Qml.");
}