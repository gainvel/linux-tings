import QtQuick
import QtQuick.Layouts
import QtQuick.Controls as QQC2

FocusScope {
    id: root

    property string notificationMessage: ""
    property alias text: passwordField.text

    signal loginRequest(string password)

    implicitWidth: 320
    implicitHeight: column.implicitHeight

    function clear() {
        passwordField.text = ""
    }

    function focusPassword() {
        passwordField.forceActiveFocus()
    }

    SequentialAnimation {
        id: rejectAnimation

        NumberAnimation { target: column; property: "x"; to: -14; duration: 50; easing.type: Easing.OutQuad }
        NumberAnimation { target: column; property: "x"; to: 14; duration: 100; easing.type: Easing.OutQuad }
        NumberAnimation { target: column; property: "x"; to: -8; duration: 100; easing.type: Easing.OutQuad }
        NumberAnimation { target: column; property: "x"; to: 8; duration: 100; easing.type: Easing.OutQuad }
        NumberAnimation { target: column; property: "x"; to: 0; duration: 80; easing.type: Easing.OutQuad }
    }

    function reject() {
        rejectAnimation.start()
    }

    Column {
        id: column
        width: parent.width
        spacing: 14

        Column {
            width: parent.width
            spacing: 6

            Text {
                text: "Password"
                color: "#ffffff"
                opacity: 0.5
                font.pixelSize: 11
                font.letterSpacing: 0.5
                renderType: Text.CurveRendering
            }

            Item {
                width: parent.width
                height: passwordField.implicitHeight

                QQC2.TextField {
                    id: passwordField
                    width: parent.width
                    echoMode: TextInput.Password
                    font.pixelSize: 14
                    horizontalAlignment: Text.AlignLeft
                    color: "#ffffff"
                    selectionColor: Qt.rgba(1, 1, 1, 0.3)
                    selectedTextColor: "#ffffff"
                    focus: true
                    leftPadding: 0
                    rightPadding: 0

                    placeholderText: ""

                    background: Item {
                        Rectangle {
                            anchors.bottom: parent.bottom
                            width: parent.width
                            height: 1
                            color: passwordField.activeFocus
                                ? Qt.rgba(1, 1, 1, 0.7)
                                : Qt.rgba(1, 1, 1, 0.3)

                            Behavior on color {
                                ColorAnimation { duration: 200 }
                            }
                        }
                    }

                    Keys.onReturnPressed: root.loginRequest(passwordField.text)
                    Keys.onEnterPressed: root.loginRequest(passwordField.text)
                }
            }
        }

        Text {
            id: notificationLabel
            text: root.notificationMessage
            color: Qt.rgba(1, 0.55, 0.55, 1.0)
            font.pixelSize: 11
            visible: text !== ""
            renderType: Text.CurveRendering
        }
    }
}
