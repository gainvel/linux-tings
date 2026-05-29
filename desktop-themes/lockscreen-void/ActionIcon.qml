import QtQuick
import org.kde.kirigami as Kirigami

Item {
    id: root

    property string iconName
    property bool hovered: mouseArea.containsMouse
    property bool pressed: mouseArea.pressed

    signal clicked()

    implicitWidth: 28
    implicitHeight: 28

    Kirigami.Icon {
        anchors.centerIn: parent
        width: 22
        height: 22
        source: root.iconName
        color: "#ffffff"
        isMask: true
        opacity: root.pressed ? 1.0 : root.hovered ? 0.9 : 0.6

        Behavior on opacity {
            NumberAnimation { duration: 150; easing.type: Easing.InOutQuad }
        }
    }

    MouseArea {
        id: mouseArea
        anchors.fill: parent
        anchors.margins: -10
        hoverEnabled: true
        cursorShape: Qt.PointingHandCursor
        onClicked: root.clicked()
    }
}
