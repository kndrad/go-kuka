# go-kuka

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)

## About <a name = "about"></a>

This project was created as part of a university project, and it simulates the control of a Kuka robot. The robot is represented in the program by different parts, and each part can be moved individually with a specified degree, acceleration, and number of steps. The project is written in Go and allows users to generate a `.dat` file which contains the movement data for each part of the robot. This data can be used for further analysis or control of a real-world Kuka robot.

## Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before you begin, ensure you have met the following requirements:

- You have installed Go. If not, you can download it from [here](https://golang.org/dl/).
- You have a basic understanding of Go programming language.

### Installing

Follow these steps to get a development environment running:

1. Clone the repository to your local machine.
2. Navigate to the project directory.

## Usage <a name = "usage"></a>

The main program, `main.go`, creates an instance of the Kuka robot and performs a series of movements on each of its parts. It moves each part to a specified degree, then reverts the movement. The movement data for each part is written to a `.dat` file.

You can run the program by using the following command:

```
go run main.go
```

The created `.dat` file will be in the same directory. The format of the data in the `.dat` file is as follows:

``` 
KukaTheta-1 KukaTheta-2 KukaTheta-3 KukaTheta-4 KukaTheta-5 KukaTheta-6
0.00 0.00 0.00 0.00 0.00 0.00
...
```

Each line represents a single step in time, and the six values represent the current position (theta value) of each part of the robot.

## Viewing the Kuka Robot Animation

To see the Kuka robot in action, follow these steps:

1. Navigate to the `RoboWorksDemo3.0` directory in this repository.
2. From there, go into the `Examples` directory.
3. Launch the `Kuka.scn` file.
4. From the navigation bar, select `Animation` -> `From File`.
5. Choose the `Kuka.dat` file from the root directory of this repository.

Please note that these instructions are applicable to Windows only.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
