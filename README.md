# Hypercube

A Go simulation of an n-dimensional hypercube object.

## Overview

This code simulates a hypercube (starting with a 4D tesseract) and extends to n-dimensional space. It defines the structure, simulates rotation, and projects the object into 3D space for visualization purposes.

### Initial 4D Hypercube (Tesseract)
- **Vertices**: 16 (each coordinate is either -1 or 1)
- **Edges**: 32 (connecting vertices that differ in exactly one coordinate)

### Core Functionality
- **Creation**: Generates the hypercube structure.
- **Rotation**: Supports rotation in XY and ZW planes (ZW is 4D-specific).
- **Projection**: Implements a simple perspective projection from 4D to 3D space.
- **Simulation**: Rotates the hypercube over time, projects 4D coordinates to 3D, and prints sample vertex positions.

### Extending to Higher Dimensions (n > 4)
To extend the simulation beyond 4D, you can:
- Make the dimension configurable in `PointND`.
- Programmatically generate 2^n vertices.
- Add rotation planes (e.g., XY, XZ, XW, YZ, YW, ZW, etc.).
- Adjust the projection function for higher dimensions.

### Output
The simulation outputs the hypercube’s rotation through 4D space, with coordinates projected to 3D. In a real application, these coordinates would typically be fed into a graphics library (e.g., OpenGL) for visualization rather than just printed.

#### Sample Output
Frame at t=0.00
Sample projected vertices (first 4):
Vertex 0: [-0.50, -0.50, -0.50]
Vertex 1: [-0.50, -0.50, 0.50]
Vertex 2: [-0.50, 0.50, -0.50]
Vertex 3: [-0.50, 0.50, 0.50]
Number of edges: 32

Frame at t=0.10
Sample projected vertices (first 4):
Vertex 0: [-0.52, -0.47, -0.50]
Vertex 1: [-0.52, -0.47, 0.50]
Vertex 2: [-0.47, 0.52, -0.50]
Vertex 3: [-0.47, 0.52, 0.50]
Number of edges: 32


## Key Features

- **Configurable Dimensions**: The `NewHypercube` function accepts a dimension parameter, supporting any number of dimensions (memory-limited).
- **Efficient Vertex Generation**: Uses bit operations to generate 2^n vertices, with coordinates set to -1 or 1.
- **Edge Generation**: Creates edges between vertices differing in exactly one coordinate, leveraging bitwise XOR for efficiency.
- **Generalized Rotation**: Implements rotation across all possible pairs of dimensions (n*(n-1)/2 planes) with varying speeds.
- **Scalable Projection**: Projects to 3D using the first three dimensions when available, applying a simple perspective projection based on the last dimension.
- **Memory Consideration**: Theoretically supports thousands of dimensions, though practical limits depend on available memory due to exponential vertex growth (2^n).

### Usage Examples
```go
// Configurable parameters
dimensions := 10     // Number of dimensions
verticesToShow := 20 // Number of vertices to display
```

## Understanding Frames in the Hypercube Simulation

In this hypercube simulation, "frames" refer to discrete snapshots or time steps that illustrate the hypercube's rotation through n-dimensional space. They allow us to visualize the object’s dynamic behavior by showing its state—specifically, the positions of its projected vertices—at regular intervals.

### Purpose of Frames

- **Animation Simulation**:
  - The hypercube starts as a static object. By applying incremental rotations at each step, we simulate motion.
  - Frames represent these updates, enabling observation of the hypercube’s evolution over time, mimicking an animation (though output here is text-based).

- **Time Discretization**:
  - The variable `t` in the simulation loop (`for t := 0.0; t < 2*math.Pi; t += 0.2`) represents time or an angle parameter. Each loop iteration is a new frame.
  - The step size (e.g., `0.2`) controls frame frequency—smaller steps yield smoother transitions, while larger steps make motion coarser.

- **Visualization of Change**:
  - At each frame, the hypercube’s vertices are rotated and projected from n-dimensional space to 3D. Printing a subset of these projected coordinates shows positional changes due to rotation.
  - For example, vertex coordinates shift between `Frame at t=0.00` and `Frame at t=0.20`, demonstrating the rotation effect.

- **Debugging and Analysis**:
  - Frames provide checkpoints to inspect the simulation, verifying that rotations and projections work correctly or analyzing the hypercube’s behavior across dimensions.

### Implementation

- **Loop Structure**: The `main` function uses a loop incrementing `t` from 0 to 2π (a full rotation cycle). Each iteration is a frame:
  ```go
  for t := 0.0; t < 2*math.Pi; t += 0.2 {
      hypercube.Rotate(t)
      PrintFrame(hypercube, t, verticesToShow)
  }
  ```