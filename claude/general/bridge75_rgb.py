#!/usr/bin/env python3
"""Disable RGB on Bridge75 keyboard via VIA v3 raw HID protocol."""

import sys
import hid

VID = 0xFFFE
PID = 0x003E
REPORT_SIZE = 32
VIA_UNHANDLED = 0xFF

CMD_CUSTOM_SET_VALUE = 0x07
CMD_CUSTOM_GET_VALUE = 0x08
CMD_CUSTOM_SAVE = 0x09

CH_RGB_MATRIX = 3
VID_BRIGHTNESS = 0
VID_EFFECT = 1
VID_SPEED = 2


def make_report(data):
    buf = [0x00] + list(data)
    buf += [0x00] * (REPORT_SIZE + 1 - len(buf))
    return bytes(buf)


def send_recv(dev, data):
    dev.write(make_report(data))
    return dev.read(REPORT_SIZE + 1, timeout_ms=2000)


def get_value(dev, channel, value_id):
    resp = send_recv(dev, [CMD_CUSTOM_GET_VALUE, channel, value_id])
    if resp and resp[0] != VIA_UNHANDLED:
        return resp[3]
    return None


def set_value(dev, channel, value_id, value):
    resp = send_recv(dev, [CMD_CUSTOM_SET_VALUE, channel, value_id, value])
    return resp and resp[0] != VIA_UNHANDLED


def save(dev, channel):
    resp = send_recv(dev, [CMD_CUSTOM_SAVE, channel])
    return resp and resp[0] != VIA_UNHANDLED


def main():
    path = None
    for info in hid.enumerate(VID, PID):
        if info["usage_page"] == 0xFF60:
            path = info["path"]

    dev = hid.device()
    dev.open_path(path or b"/dev/hidraw4")
    dev.set_nonblocking(False)

    # Read current state
    brightness = get_value(dev, CH_RGB_MATRIX, VID_BRIGHTNESS)
    effect = get_value(dev, CH_RGB_MATRIX, VID_EFFECT)
    speed = get_value(dev, CH_RGB_MATRIX, VID_SPEED)
    print(f"Current: brightness={brightness}, effect={effect}, speed={speed}")

    # Set effect to 0 (RGB_MATRIX_NONE = off) and brightness to 0
    print("Setting effect to 0 (off)...")
    ok1 = set_value(dev, CH_RGB_MATRIX, VID_EFFECT, 0)
    print(f"  Effect set: {'OK' if ok1 else 'FAILED'}")

    print("Setting brightness to 0...")
    ok2 = set_value(dev, CH_RGB_MATRIX, VID_BRIGHTNESS, 0)
    print(f"  Brightness set: {'OK' if ok2 else 'FAILED'}")

    print("Saving to EEPROM...")
    ok3 = save(dev, CH_RGB_MATRIX)
    print(f"  Save: {'OK' if ok3 else 'FAILED'}")

    # Verify
    brightness = get_value(dev, CH_RGB_MATRIX, VID_BRIGHTNESS)
    effect = get_value(dev, CH_RGB_MATRIX, VID_EFFECT)
    print(f"After: brightness={brightness}, effect={effect}")

    dev.close()
    print("\nDone. Check if RGB is off and caps lock indicator still works.")


if __name__ == "__main__":
    main()
