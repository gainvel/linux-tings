import QtQuick
import QtQuick.Controls as QQC2

Item {
    id: root

    width: 1920
    height: 1080

    property bool passwordVisible: false
    property string notificationMessage: ""
    property int currentSessionIndex: sessionModel.lastIndex

    Repeater {
        model: screenModel
        Image {
            x: geometry.x
            y: geometry.y
            width: geometry.width
            height: geometry.height
            source: config.background
            fillMode: Image.PreserveAspectCrop
            smooth: true
        }
    }

    Rectangle {
        anchors.fill: parent
        gradient: Gradient {
            GradientStop { position: 0.0; color: "transparent" }
            GradientStop { position: 0.55; color: "transparent" }
            GradientStop { position: 1.0; color: Qt.rgba(0, 0, 0, 0.35) }
        }
    }

    MouseArea {
        id: inputArea
        anchors.fill: parent
        hoverEnabled: true

        onClicked: {
            root.passwordVisible = true
            passwordPrompt.focusPassword()
        }

        Keys.onPressed: event => {
            root.passwordVisible = true
            passwordPrompt.focusPassword()
            event.accepted = false
        }

        focus: true

        ClockDisplay {
            id: clock
            width: parent.width
            anchors.horizontalCenter: parent.horizontalCenter
            y: parent.height * 0.30 - height / 2
        }

        PasswordPrompt {
            id: passwordPrompt
            anchors.horizontalCenter: parent.horizontalCenter
            width: 320

            y: root.passwordVisible
                ? clock.y + clock.height + 40
                : clock.y + clock.height + 60

            opacity: root.passwordVisible ? 1.0 : 0.0
            enabled: root.passwordVisible

            notificationMessage: root.notificationMessage

            Behavior on opacity {
                NumberAnimation {
                    duration: 400
                    easing.type: Easing.OutCubic
                }
            }

            Behavior on y {
                NumberAnimation {
                    duration: 400
                    easing.type: Easing.OutCubic
                }
            }

            onLoginRequest: password => {
                root.notificationMessage = ""
                sddm.login(userModel.lastUser, password, root.currentSessionIndex)
            }
        }

        Row {
            id: bottomBar
            anchors.left: parent.left
            anchors.bottom: parent.bottom
            anchors.leftMargin: 20
            anchors.bottomMargin: 20
            spacing: 24

            ActionIcon {
                iconName: "system-suspend"
                visible: sddm.canSuspend
                onClicked: sddm.suspend()
            }

            ActionIcon {
                iconName: "system-reboot"
                visible: sddm.canReboot
                onClicked: sddm.reboot()
            }

            ActionIcon {
                iconName: "system-shutdown"
                visible: sddm.canPowerOff
                onClicked: sddm.powerOff()
            }

            Rectangle {
                width: 1
                height: 22
                color: "#ffffff"
                opacity: 0.2
                visible: sessionModel.rowCount() > 1
                anchors.verticalCenter: parent.verticalCenter
            }

            Item {
                id: sessionSwitcher
                width: sessionRow.width
                height: sessionRow.height
                visible: sessionModel.rowCount() > 1

                Row {
                    id: sessionRow
                    spacing: 8

                    ActionIcon {
                        id: sessionIcon
                        iconName: "computer-symbolic"
                        onClicked: sessionPopup.visible = !sessionPopup.visible
                    }

                    Text {
                        anchors.verticalCenter: sessionIcon.verticalCenter
                        text: sessionModel.data(sessionModel.index(root.currentSessionIndex, 0), Qt.DisplayRole) || ""
                        color: "#ffffff"
                        opacity: 0.6
                        font.pixelSize: 13
                        renderType: Text.CurveRendering
                    }
                }

                Column {
                    id: sessionPopup
                    visible: false
                    anchors.bottom: sessionRow.top
                    anchors.bottomMargin: 8
                    anchors.left: sessionRow.left
                    spacing: 2

                    Repeater {
                        model: sessionModel
                        delegate: Rectangle {
                            width: sessionText.implicitWidth + 24
                            height: 32
                            radius: 6
                            color: sessionMouseArea.containsMouse
                                ? Qt.rgba(1, 1, 1, 0.15)
                                : index === root.currentSessionIndex
                                    ? Qt.rgba(1, 1, 1, 0.08)
                                    : "transparent"

                            Text {
                                id: sessionText
                                anchors.centerIn: parent
                                text: name
                                color: "#ffffff"
                                font.pixelSize: 13
                                renderType: Text.CurveRendering
                            }

                            MouseArea {
                                id: sessionMouseArea
                                anchors.fill: parent
                                hoverEnabled: true
                                cursorShape: Qt.PointingHandCursor
                                onClicked: {
                                    root.currentSessionIndex = index
                                    sessionPopup.visible = false
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    Timer {
        id: hideTimer
        interval: 30000
        running: root.passwordVisible && passwordPrompt.text === ""
        onTriggered: root.passwordVisible = false
    }

    Timer {
        id: notificationTimer
        interval: 4000
        onTriggered: root.notificationMessage = ""
    }

    Connections {
        target: sddm
        function onLoginSucceeded() {
        }
        function onLoginFailed() {
            root.notificationMessage = "Authentication failed"
            notificationTimer.restart()
            passwordPrompt.reject()
            passwordPrompt.clear()
            passwordPrompt.focusPassword()
        }
    }
}
