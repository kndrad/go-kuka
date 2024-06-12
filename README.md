# go-kuka

This Go project simulates the movement of a Kuka robot by generating a `.dat` file containing the position data for each part of the robot. The movement is calculated using smooth acceleration and deceleration, resulting in a realistic simulation experience.

## How it Works

The project defines a `Part` struct to represent each individual part of the robot, such as the base, body, arm, wrist, tool, and disk. Each part has a unique ID, name, and a `Theta` value representing its current position.

The `Move` and `Reset` methods on the `Part` struct calculate the new positions for a part based on the target position, acceleration, and number of steps. The `Move` method uses a sine-based easing function to smoothly transition from the current position to the target position, while the `Reset` method transitions the part back to the initial position (0 degrees).

The `Kuka` struct represents the entire robot and contains a slice of `Part` pointers. It provides methods to move each part of the robot, such as `MoveBase`, `MoveBody`, `MoveArm`, `MoveWrist`, `MoveTool`, and `MoveDisk`. These methods call the `MovePart` method, which calculates the new positions for the specified part and writes the position data for all parts to the `.dat` file.

## Usage

Clone the repository or download the source code, navigate to the project directory, then run or build using Go.

The program will generate a `Kuka.dat` file in the same directory, containing the position data for each part of the robot during the simulated movement.

## Viewing the Kuka Robot Animation

Launch the `Kuka.scn` file from the `RoboWorksDemo3.0/Examples` directory, then choose the `Kuka.dat` generated file.
Unfortunately, the software (RoboWorksDemo3.0) we've used at the university to visualize the Kuka robot animation seems to be no longer available for download from the creator's site (www.newtonium.com) because it's no longer hosted, so it's available as a trial version in this repository under the creator's SHAREWARE License.

# License
This project is licensed under the MIT License - see the LICENSE file for details.
