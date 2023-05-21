package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	Base int = iota
	Body
	Arm
	Wrist
	Tool
	BlackDisk
)

// Part represents a part of the robot. Each part has an ID, a name, and a theta value
// representing its current position.
type Part struct {
	ID    int
	Name  string
	Theta float64
}

func NewPart(id int, name string) *Part {
	return &Part{
		ID:    id,
		Name:  name,
		Theta: 0,
	}
}

const (
	Steps        = 300
	Acceleration = 15.0
)

// Move calculates a list of new theta positions for the part, based on the input theta,
// acceleration, and number of steps. It uses a sine-based easing function to smoothly
// transition from the part's current position to the target position. The part's final
// position is updated to the target position.
func (p *Part) Move(theta float64, acceleration float64, steps int) []float64 {
	if steps < 0 {
		steps = 0
	}
	var (
		moves = make([]float64, 0)
		start = p.Theta
		end   = start + theta
		acc   = math.Abs(acceleration)
	)
	for i := 1; i <= Steps; i++ {
		theta := start + (end-start)*math.Pow(math.Sin(float64(i)/float64(Steps)*(math.Pi/2)), acc)
		moves = append(moves, theta)
	}
	p.Theta = end
	return moves
}

// Reset calculates a list of new theta positions for the part, smoothly transitioning
// from the current position to 0. The part's final position is updated to 0.
func (p *Part) Reset(steps int) []float64 {
	if steps < 0 {
		steps = 0
	}
	var (
		moves = make([]float64, 0)
		start = p.Theta
		end   = 0.0
	)
	for i := 1; i <= steps; i++ {
		coef := math.Cos(float64(i) / float64(steps) * (math.Pi / 2))
		theta := start * coef
		moves = append(moves, theta)
	}
	p.Theta = end
	return moves
}

// Parts is a slice of Part pointers. The Label method generates a string that describes
// all parts in the slice.
type Parts []*Part

type Label string

func (p Parts) Label() Label {
	var b strings.Builder
	for _, part := range p {
		b.WriteString(fmt.Sprintf("%s ", part.Name))
	}
	return Label(b.String())
}

func (l Label) Write(file *os.File) error {
	_, err := file.WriteString(string(l))
	if err != nil {
		return err
	}
	file.Write([]byte("\n"))
	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}

// Parts is a slice of Part pointers. The Label method generates a string that describes
// all parts in the slice.
type Kuka struct {
	parts Parts
	file  *os.File
}

func NewKuka() (*Kuka, error) {
	f, err := os.Create("Kuka.dat")
	if err != nil {
		return nil, err
	}
	parts := Parts{
		NewPart(Base, "KukaTheta-1"),
		NewPart(Body, "KukaTheta-2"),
		NewPart(Arm, "KukaTheta-3"),
		NewPart(Wrist, "KukaTheta-4"),
		NewPart(Tool, "KukaTheta-5"),
		NewPart(BlackDisk, "KukaTheta-6"),
	}
	if err := parts.Label().Write(f); err != nil {
		return nil, err
	}
	return &Kuka{
		parts: parts,
		file:  f,
	}, nil
}

// MovePart moves a specified part of the robot, records the new theta positions of all
// parts to the data file, and updates the part's final position in the robot's parts slice.
func (k *Kuka) MovePart(id int, theta float64, acceleration float64, steps int) error {
	moves := k.parts[id].Move(theta, acceleration, steps)

	var thetas []float64
	for _, part := range k.parts {
		thetas = append(thetas, part.Theta)
	}

	var b strings.Builder
	for _, move := range moves {
		thetas[id] = move

		for _, theta := range thetas {
			b.WriteString(fmt.Sprintf("%.2f ", theta))
		}
		b.WriteByte('\n')

		if _, err := k.file.WriteString(b.String()); err != nil {
			return err
		}
		if err := k.file.Sync(); err != nil {
			return err
		}
		b.Reset() // Reset the builder for the next iteration
	}
	if err := k.file.Sync(); err != nil {
		return err
	}
	return nil
}

// MoveBase, MoveBody, MoveArm, MoveWrist, MoveTool, and MoveDisk are convenience methods
// that move specific parts of the robot. They call the MovePart method with the appropriate
// part ID.
func (k *Kuka) MoveBase(theta, acceleration float64, steps int) error {
	if err := k.MovePart(Base, theta, acceleration, steps); err != nil {
		return err
	}
	return nil
}

// ... similar methods for Body, Arm, Wrist, Tool, and Disk ...
func (k *Kuka) MoveBody(theta, acceleration float64, steps int) error {
	if err := k.MovePart(Body, theta, acceleration, steps); err != nil {
		return err
	}
	return nil
}

func (k *Kuka) MoveArm(theta, acceleration float64, steps int) error {
	if err := k.MovePart(Arm, theta, acceleration, steps); err != nil {
		return err
	}
	return nil
}

func (k *Kuka) MoveWrist(theta, acceleration float64, steps int) error {
	if err := k.MovePart(Wrist, theta, acceleration, steps); err != nil {
		return err
	}
	return nil
}

func (k *Kuka) MoveTool(theta, acceleration float64, steps int) error {
	if err := k.MovePart(Tool, theta, acceleration, steps); err != nil {
		return err
	}
	return nil
}

func (k *Kuka) MoveDisk(theta, acceleration float64, steps int) error {
	if err := k.MovePart(BlackDisk, theta, acceleration, steps); err != nil {
		return err
	}
	return nil
}
