package main

import "log"

func main() {
	kuka, err := NewKuka()
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}
	defer kuka.file.Close()

	// Move every part
	kuka.MoveBase(45.0, Acceleration, Steps)
	kuka.MoveBody(45.0, Acceleration, Steps)
	kuka.MoveArm(45.0, Acceleration, Steps)
	kuka.MoveWrist(45.0, Acceleration, Steps)
	kuka.MoveTool(45.0, Acceleration, Steps)
	kuka.MoveDisk(45.0, Acceleration, Steps)

	// Revert
	kuka.MoveBase(-45.0, Acceleration, Steps)
	kuka.MoveBody(-45.0, Acceleration, Steps)
	kuka.MoveArm(-45.0, Acceleration, Steps)
	kuka.MoveWrist(-45.0, Acceleration, Steps)
	kuka.MoveTool(-45.0, Acceleration, Steps)
	kuka.MoveDisk(-45.0, Acceleration, Steps)
}
