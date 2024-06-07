#
# Licensed under 3-Clause BSD license available in the License file. Copyright (c) 2021-2023 iRobot Corporation. All rights reserved.
#

from irobot_edu_sdk.backend.bluetooth import Bluetooth
from irobot_edu_sdk.robots import event, hand_over, Color, Robot, Root, Create3
from irobot_edu_sdk.music import Note

import asyncio

print('gay robot')

async def main():
    robot = Create3(Bluetooth())
    #await robot.play_note(Note.A4, .1)
    #
    methods = [
        "get_cliff_sensors",
        "get_cliff_sensors_cached",
        "get_docking_values",
        "get_enabled_events",
        "get_ipv4_address",
        "get_ir_proximity",
        "get_name",
        "get_packed_ir_proximity",
        "get_position",
        "get_serial_number",
        "get_sku",
        "get_touch_sensors",
        "get_touch_sensors_cached",
        "get_version_string",
        "get_versions",
        "get_6x_ir_proximity",
        "get_7x_ir_proximity",
        "get_accelerometer",
        "get_battery_level",
        "get_bumpers",
        "get_bumpers_cached",
    ]

    for method in methods:
        try:
            result = await getattr(robot, method)()
            print(f"{method}: {result}")
        except Exception as e:
            print(f"Error calling {method}: {e}")

asyncio.run(main())

# robot = Create3(Bluetooth())

# duration = 0.15


# @event(robot.when_touched, [True, False])  # (.) button.
# async def touched(robot):
#     await robot.set_lights_on_rgb(255, 0, 0)
#     await robot.play_note(Note.A4, duration)


# @event(robot.when_touched, [False, True])  # (..) button.
# async def touched(robot):
#     await robot.set_lights_on_rgb(0, 255, 0)
#     await robot.play_note(Note.C5_SHARP, duration)


# @event(robot.when_touched, [True, True])
# async def touched(robot):
#     print('ANY sensor touched')


# @event(robot.when_play)
# async def play(robot):
#     await robot.play_note(Note.C5_SHARP, duration)
#     await robot.play_note(Note.C5_SHARP, duration)
#     await robot.play_note(Note.C5_SHARP, duration)
#     await robot.play_note(Note.C5_SHARP, duration)
#     await robot.play_note(Note.A5, duration)

# robot.play()
