#ifndef QMLINTEGRATION_H
#define QMLINTEGRATION_H

#include <QObject>

class qmlintegration : public QObject
{
	Q_Object
public:
	explicit qmlintegration(QObject *parent = nullptr);

signals:

public slots:
	void callFromQml();	
};

#endif //QMLINTEGRATION_H