#include "Qmlintegration.h"
#include <QDebug>

Qmlintegration::Qmlintegration(QObject *parent) : QObject(parent){

}
void Qmlintegration::callFromQml(){
	qDebug("Recieved call from Qml.");
}