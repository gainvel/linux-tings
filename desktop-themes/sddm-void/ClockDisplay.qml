import QtQuick

Item {
    id: root

    property alias fontLoader: laroFont
    property real timeSize: Math.round(parent.height * 0.12)
    property real dateSize: Math.round(parent.height * 0.025)

    implicitWidth: column.implicitWidth
    implicitHeight: column.implicitHeight

    FontLoader {
        id: laroFont
        source: "fonts/LaroSoftExtraBold.ttf"
    }

    Timer {
        id: clockTimer
        interval: 1000
        running: true
        repeat: true
        onTriggered: {
            timeLabel.text = Qt.formatTime(new Date(), "HH:mm")
            dateLabel.text = Qt.formatDate(new Date(), "dddd, d MMMM")
        }
    }

    Column {
        id: column
        anchors.horizontalCenter: parent.horizontalCenter
        spacing: Math.round(root.dateSize * 0.3)

        Text {
            id: timeLabel
            anchors.horizontalCenter: parent.horizontalCenter
            text: Qt.formatTime(new Date(), "HH:mm")
            color: "#ffffff"
            font.family: laroFont.name
            font.pixelSize: root.timeSize
            font.letterSpacing: -2.0
            renderType: Text.CurveRendering
        }

        Text {
            id: dateLabel
            anchors.horizontalCenter: parent.horizontalCenter
            text: Qt.formatDate(new Date(), "dddd, d MMMM")
            color: "#ffffff"
            opacity: 0.7
            font.family: laroFont.name
            font.pixelSize: root.dateSize
            font.letterSpacing: 0.5
            renderType: Text.CurveRendering
        }
    }
}
