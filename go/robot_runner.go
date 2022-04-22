package main

import "fmt"

type (
        // Position stores x and y coordinates on a plane
        Position struct {
                x int32
                y int32
        }

        // Direction is possible directions for robots position & movement,
        // which are N, E, S, W
        Direction rune

        // Command is possible commands which can be given to the robot,
        // which are L, R, M
        Command rune

        // Robot has its position on plane and direction it is facing to.
        // Changes position/direction by "Act"ing based on a command
        Robot struct {
                pos Position
                dir Direction

                plane   Plane
                visited []Position
        }

        // Plane is a rectangluar plane, can have particles on it.
        //
        // TODO: particles are not implemented and problem has no
        //      inputs for particles
        Plane struct {
                rows int32 // number of cols
                cols int32 // number of rows

                particles []Position
        }
)

const (
        // closckwise starting from North
        N Direction = 'N' // North
        E           = 'E' // East
        S           = 'S' // South
        W           = 'W' // West
)

const (
        // commands
        L Command = 'L' // Turn left (90 deg)
        R         = 'R' // Turn right (90 deg)
        M         = 'M' // move forward
)

// NewRobot creates and returns a poiter to new Robot.
// x, y & dir are starting coordinates and direction of the robot
func NewRobot(x, y int32, dir Direction, plane Plane) *Robot {
        return &Robot{
                pos: Position{
                        x: x,
                        y: y,
                },
                dir:   dir,
                plane: plane,
        }
}

// isVisitied returns true if position p is already visited by robot,
// false otherwise
func (r *Robot) isVisited(p Position) bool {
        for _, v := range r.visited {
                if v.x == p.x && v.y == p.y {
                        return true
                }
        }
        return false
}

// Act given the cmd Command, will move or turn the robot
func (r *Robot) Act(cmd Command) bool {
        switch cmd {
        case M:
                var next Position
                switch r.dir {
                case N:
                        next = Position{r.pos.x, r.pos.y + 1}
                case E:
                        next = Position{r.pos.x + 1, r.pos.y}
                case S:
                        next = Position{r.pos.x, r.pos.y - 1}
                case W:
                        next = Position{r.pos.x - 1, r.pos.y}
                }

                if !r.plane.IsValid(next) {
                        return false
                }
                if r.isVisited(next) {
                        return false
                }
                // otherwise move
                r.visited = append(r.visited, r.pos)
                r.pos = next
        default:
                r.dir = turnDirection(r.dir, cmd)
        }
        return true
}

// NewPlace creates and returns a new rectangular plane.
// m & n are coordinates of top-right corner of plane, plane
// starts at bottom-left corner at (0, 0)
func NewPlane(m, n int32, particles []Position) Plane {
        return Plane{
                rows:      n + 1,
                cols:      m + 1,
                particles: particles,
        }
}

// IsValid checks if position is valid for the plane
func (p *Plane) IsValid(pos Position) bool {
        if pos.x < 0 || pos.x >= p.cols || pos.y < 0 || pos.y >= p.rows {
                return false
        }
        for _, par := range p.particles {
                if par.x == pos.x && par.y == pos.y {
                        return false
                }
        }
        return true
}

// turnDirection returns next direction based on command,
// gives same direction if c is not a valid command
func turnDirection(d Direction, c Command) Direction {
        switch c {
        case R:
                switch d {
                case N:
                        return E
                case E:
                        return S
                case S:
                        return W
                case W:
                        return N
                }
        case L:
                switch d {
                case N:
                        return W
                case E:
                        return N
                case S:
                        return E
                case W:
                        return S
                }
        }
        return d
}

func main() {
        // input plane coordinates
        var (
                pm int32
                pn int32
        )
        fmt.Println("Enter top-right coordinates of plane")
        fmt.Scan(&pm, &pn)

        // input particle coordinates for plane
        var (
                a, b      int32
                particles []Position
        )
        fmt.Println("Enter particle coordinates (enter -1 -1 to end particles input)")
        for {
                fmt.Scan(&a, &b)
                if a == -1 && b == -1 {
                        break
                }
                particles = append(particles, Position{a, b})
        }

        // input robot position & direction
        var (
                rx   int32
                ry   int32
                rdir string
        )
        fmt.Println("Enter robot position and direction")
        fmt.Scan(&rx, &ry, &rdir)

        // input commands for robot
        var commands string
        fmt.Println("Enter commands for robot")
        fmt.Scan(&commands)

        plane := NewPlane(pm, pn, particles)

        r := NewRobot(rx, ry, Direction(rdir[0]), plane)

        for _, cmd := range commands {
                if ok := r.Act(Command(cmd)); !ok {
                        break
                }
        }

        // fmt.Println(r.visited)
        fmt.Println(r.pos.x, r.pos.y, string(r.dir))
}
