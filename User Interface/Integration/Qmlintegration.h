#ifndef QMLINTEGRATION_H
#define QMLINTEGRATION_H

#include <QObject>

class Qmlintegration : public QObject
{
	Q_Object
public:
	explicit Qmlintegration(QObject *parent = nullptr);

signals:

public slots:
	void callFromQml();	
};

#endif //QMLINTEGRATION_H