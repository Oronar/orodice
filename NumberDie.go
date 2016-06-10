package orodice

import "math/rand"
import "time"
import "strconv"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type NumberDie struct {
	current_face int
	faces []NumberDieResult
}

func NewNumberDie(size int) NumberDie{
	die := NumberDie{0, make([]NumberDieResult, size)}
	for i:=0;i<size;i++ {
		die.faces[i].Value = i+1
	}
	return die
}

func (d *NumberDie) Roll() {
	d.current_face = rand.Intn(len(d.faces))
}

func (d *NumberDie) GetFace(face int) NumberDieResult {
	return d.faces[face]
}

func (d *NumberDie) GetCurrentFace() NumberDieResult {
	return d.GetFace(d.current_face)
}

type NumberDieResult struct {
	Value int
}

func (r *NumberDieResult) Add(a NumberDieResult) {
	r.Value += a.Value
}

func (r NumberDieResult) String() string{
	return strconv.Itoa(r.Value)
}