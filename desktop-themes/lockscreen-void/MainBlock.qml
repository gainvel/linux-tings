import QtQuick
import QtQuick.Layouts
import QtQuick.Controls as QQC2

import org.kde.plasma.components as PlasmaComponents3
import org.kde.kirigami as Kirigami

FocusScope {
    id: mainBlockRoot

    property alias mainPasswordBox: passwordBox
    property bool lockScreenUiVisible: false
    property string notificationMessage: ""

    signal passwordResult(string password)

    implicitHeight: column.implicitHeight

    function startLogin() {
        var password = passwordBox.text
        passwordResult(password)
    }

    Column {
        id: column
        anchors.horizontalCenter: parent.horizontalCenter
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
                height: passwordBox.implicitHeight

                QQC2.TextField {
                    id: passwordBox
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

                    property bool showPassword: false
                    cursorVisible: visible

                    placeholderText: ""

                    background: Item {
                        Rectangle {
                            anchors.bottom: parent.bottom
                            width: parent.width
                            height: 1
                            color: passwordBox.activeFocus
                                ? Qt.rgba(1, 1, 1, 0.7)
                                : Qt.rgba(1, 1, 1, 0.3)

                            Behavior on color {
                                ColorAnimation { duration: 200 }
                            }
                        }
                    }

                    onAccepted: {
                        if (mainBlockRoot.lockScreenUiVisible) {
                            mainBlockRoot.startLogin()
                        }
                    }

                    Keys.onPressed: event => {
                        if (event.key === Qt.Key_Escape) {
                            passwordBox.text = ""
                            event.accepted = true
                        }
                    }

                    Component.onCompleted: {
                        if (typeof PasswordSync !== "undefined") {
                            passwordBox.text = Qt.binding(() => PasswordSync.password)
                        }
                    }
                }
            }

            Binding {
                target: typeof PasswordSync !== "undefined" ? PasswordSync : null
                property: "password"
                value: passwordBox.text
                when: typeof PasswordSync !== "undefined"
            }
        }

        Text {
            id: notificationLabel
            text: mainBlockRoot.notificationMessage
            color: Qt.rgba(1, 0.55, 0.55, 1.0)
            font.pixelSize: 11
            visible: text !== ""
            wrapMode: Text.WordWrap
            width: parent.width
            renderType: Text.CurveRendering
        }

    }

    Connections {
        target: root
        function onClearPassword() {
            passwordBox.forceActiveFocus()
            passwordBox.text = ""
            if (typeof PasswordSync !== "undefined") {
                passwordBox.text = Qt.binding(() => PasswordSync.password)
            }
        }
    }
}
