# BOIDS

## Etymology

Coined by Craig Reynolds, mimicking bird movements using a computer program. (Bird + Android = BOID)

## Program

- Each bird will be assigned a thread and then all of the thread will communicate using some rules to mimick bird flocking simulation.
- Simulation will be 2d.
- Graphical simulation will be using Ebiten.
- BOID structure will need Position in space, velocity (speed and direction) and unique ID.
- To represent velocity we can have a point structure like (x, y) where x and y will added to current position of the the boid so that it reaches the next position. Using this we can show speed and direction both.
