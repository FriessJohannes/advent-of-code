package common

import "math"

type IntPoint struct {
    X int
    Y int
}

type FloatPoint struct {
    X float64
    Y float64 
}

type Vector struct {
    Start FloatPoint
    End FloatPoint
}

func (pos *IntPoint) IsEqual(other IntPoint) bool {
    return pos.X == other.X && pos.Y == other.Y
}

func (pos *FloatPoint) IsInRange(start FloatPoint, end FloatPoint) bool {
    return pos.X >= start.X && pos.X < end.X && pos.Y >= start.Y && pos.Y < end.Y
}

func (pos *IntPoint) ToFloatPoint() FloatPoint {
    return FloatPoint{
        X: float64(pos.X),
        Y: float64(pos.Y),
    }
}

func (pos *FloatPoint) ToIntPoint() IntPoint {
    return IntPoint{
        X: int(math.Round(pos.X)),
        Y: int(math.Round(pos.Y)),
    }
}

func (vec *Vector) Inverse() Vector {
    return Vector{
        Start: vec.End,
        End: vec.Start,
    }
}

func (vec Vector) Components() (float64, float64) {
    return vec.End.X - vec.Start.X, vec.End.Y - vec.Start.Y
}

func (vec *Vector) Rotate(angle float64) {
    dx, dy := vec.Components()

    rotatedDX := dx*math.Cos(angle) - dy*math.Sin(angle)
    rotatedDY := dx*math.Sin(angle) + dy*math.Cos(angle)

    rotatedEnd := FloatPoint{
        X: vec.Start.X + rotatedDX,
        Y: vec.Start.Y + rotatedDY,
    }

    vec.End = rotatedEnd
}


func (vec *Vector) Scale(scalar float64) {
    dx, dy := vec.Components()

    scaledDX := dx * scalar
    scaledDY := dy * scalar

    scaledEnd := FloatPoint{
        X: vec.Start.X + scaledDX,
        Y: vec.Start.Y + scaledDY,
    }

    vec.End = scaledEnd
}

func (vec *Vector) MoveTo(newStart FloatPoint) {
    dx, dy := vec.Components()

    newEnd := FloatPoint{
        X: newStart.X + dx,
        Y: newStart.Y + dy,
    }

    vec.Start = newStart
    vec.End = newEnd
}
