import QtQml
import QtQuick
import QtQuick.Layouts

import org.kde.plasma.components as PlasmaComponents3
import org.kde.plasma.private.keyboardindicator as KeyboardIndicator
import org.kde.kirigami as Kirigami
import org.kde.kscreenlocker as ScreenLocker

import org.kde.plasma.private.sessions

Item {
    id: lockScreenUi

    readonly property bool softwareRendering: GraphicsInfo.api === GraphicsInfo.Software
    property string pendingPassword: ""

    function handleMessage(msg) {
        if (!root.notification) {
            root.notification += msg
        } else if (root.notification.includes(msg)) {
            root.notificationRepeated()
        } else {
            root.notification += "\n" + msg
        }
    }

    Connections {
        target: authenticator
        function onFailed(kind) {
            if (kind != 0) return
            lockScreenUi.handleMessage("Unlocking failed")
            graceLockTimer.restart()
            notificationRemoveTimer.restart()
            rejectAnimation.start()
        }
        function onSucceeded() {
            if (authenticator.hadPrompt) {
                Qt.quit()
            }
        }
        function onInfoMessageChanged() {
            lockScreenUi.handleMessage(authenticator.infoMessage)
        }
        function onErrorMessageChanged() {
            lockScreenUi.handleMessage(authenticator.errorMessage)
        }
        function onPromptChanged(msg) {
            lockScreenUi.handleMessage(authenticator.prompt)
        }
        function onPromptForSecretChanged(msg) {
            mainBlock.mainPasswordBox.showPassword = false
            mainBlock.mainPasswordBox.forceActiveFocus()
            if (lockScreenUi.pendingPassword !== "") {
                var pw = lockScreenUi.pendingPassword
                lockScreenUi.pendingPassword = ""
                authenticator.respond(pw)
            }
        }
    }

    SessionManagement {
        id: sessionManagement
    }

    KeyboardIndicator.KeyState {
        id: capsLockState
        key: Qt.Key_CapsLock
    }

    Connections {
        target: sessionManagement
        function onAboutToSuspend() {
            lockScreenUi.pendingPassword = ""
            root.clearPassword()
        }
    }

    SequentialAnimation {
        id: rejectAnimation

        NumberAnimation { target: mainBlock; property: "anchors.horizontalCenterOffset"; to: -14; duration: 50; easing.type: Easing.OutQuad }
        NumberAnimation { target: mainBlock; property: "anchors.horizontalCenterOffset"; to: 14; duration: 100; easing.type: Easing.OutQuad }
        NumberAnimation { target: mainBlock; property: "anchors.horizontalCenterOffset"; to: -8; duration: 100; easing.type: Easing.OutQuad }
        NumberAnimation { target: mainBlock; property: "anchors.horizontalCenterOffset"; to: 8; duration: 100; easing.type: Easing.OutQuad }
        NumberAnimation { target: mainBlock; property: "anchors.horizontalCenterOffset"; to: 0; duration: 80; easing.type: Easing.OutQuad }
    }

    MouseArea {
        id: lockScreenRoot

        property bool uiVisible: true
        property bool blockUI: mainBlock.mainPasswordBox.text.length > 0

        x: parent.x
        y: parent.y
        width: parent.width
        height: parent.height
        hoverEnabled: true

        onPressed: {
            uiVisible = true
            Window.window.requestActivate()
            mainBlock.mainPasswordBox.forceActiveFocus()
        }
        onUiVisibleChanged: {
            if (uiVisible) {
                authenticator.startAuthenticating()
            }
        }

        Keys.onPressed: event => {
            uiVisible = true
            Window.window.requestActivate()
            event.accepted = false
        }

        Timer {
            id: notificationRemoveTimer
            interval: 3000
            onTriggered: root.notification = ""
        }

        Timer {
            id: graceLockTimer
            interval: 3000
            onTriggered: {
                lockScreenUi.pendingPassword = ""
                root.clearPassword()
                authenticator.startAuthenticating()
            }
        }

        PropertyAnimation {
            id: launchAnimation
            target: lockScreenRoot
            property: "opacity"
            from: 0
            to: 1
            duration: 600
        }

        Component.onCompleted: {
            launchAnimation.start()
            authenticator.startAuthenticating()
        }

        Image {
            anchors.fill: parent
            source: "images/wallpaper.jpg"
            fillMode: Image.PreserveAspectCrop
            smooth: true
        }

        Rectangle {
            anchors.fill: parent
            gradient: Gradient {
                GradientStop { position: 0.0; color: "transparent" }
                GradientStop { position: 0.55; color: "transparent" }
                GradientStop { position: 1.0; color: Qt.rgba(0, 0, 0, 0.35) }
            }
        }

        ClockDisplay {
            id: clock
            width: parent.width
            anchors.horizontalCenter: parent.horizontalCenter
            y: parent.height * 0.30 - height / 2
        }

        MainBlock {
            id: mainBlock
            anchors.horizontalCenter: parent.horizontalCenter
            width: 320
            y: clock.y + clock.height + 40

            lockScreenUiVisible: lockScreenRoot.uiVisible
            enabled: !graceLockTimer.running

            notificationMessage: {
                var parts = []
                if (capsLockState.locked) {
                    parts.push("Caps Lock is on")
                }
                if (root.notification) {
                    parts.push(root.notification)
                }
                return parts.join(" · ")
            }

            onPasswordResult: password => {
                lockScreenUi.pendingPassword = password
                authenticator.startAuthenticating()
            }
        }

        Row {
            id: powerBar
            anchors.left: parent.left
            anchors.bottom: parent.bottom
            anchors.leftMargin: 20
            anchors.bottomMargin: 20
            spacing: 24

            ActionIcon {
                iconName: "system-suspend"
                onClicked: root.suspendToRam()
            }

            ActionIcon {
                iconName: "system-reboot"
                onClicked: sessionManagement.requestReboot(SessionManagement.Skip)
            }

            ActionIcon {
                iconName: "system-shutdown"
                onClicked: sessionManagement.requestShutdown(SessionManagement.Skip)
            }
        }

        Loader {
            z: 2
            active: root.viewVisible
            source: "LockOsd.qml"
            anchors {
                horizontalCenter: parent.horizontalCenter
                bottom: parent.bottom
                bottomMargin: 20
            }
        }
    }
}
