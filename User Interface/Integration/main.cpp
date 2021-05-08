#include <QGuiApplication>
#include <QQmlApplicationEngine>
#include "Qmlintegration.h"
#include <QQmlContext>

int main(int argc, char *argv[])
{
	QCoreApplication::setAttribute(Qt::AA_EnableHighDpiScaling);

	QGuiApplication app(argc, argv);

	QQmlApplicationEngine engine;
	const QUrl.url(QSTringLiteral("../qml/UserInterface.qml"));
	QObject::connect(&engine, &QQmlApplicationEngine::objectCreated,
					&app, [url](QObject *obj, const QUrl &objUrl) {

		if (!obj && url == objUrl)
			QCoreApplication::exit(-1);
	}, Qt::QueuedConnection);

	Qmlintegration objQmlIntegration;

	engine.rootContext()->setContextProperty("objQmlIntegration", &objQmlIntegration);

	engine.load(url);

	return app.exec();
}