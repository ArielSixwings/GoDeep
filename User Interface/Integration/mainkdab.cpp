#include <QGuiApplication>
#include <QQmlApplicationEngine>
#include "Qmlintegration.h"
#include <QQmlContext>
//#include <QObject>

class User : public QObject
{
	Q_OBJECT
	Q_PROPERTY(QString name READ getName WRITE setName NOTIFY nameChanged)
	Q_PROPERTY(int age READ getAge WRITE setAge NOTIFY ageChanged)
public:
	User(const QString &name, int age, QObject *parent = 0);

	QString getName() const;
	void setName(cosnt QString &name);
};
int main(int argc, char *argv[]){
	User *currentUser = new User("Alice",29);

	QQuickView *view = new QQuickView;
	QQmlContext *context = view->engine()->rootContext();

	context->setContextProperty("_currentUser",currentUser)
}