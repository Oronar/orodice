package orodice

import "math/rand"
import "time"
import "strconv"
import "bytes"
import "strings"

const (
	ABILITY = 1
	PROFICIENCY = 2
	DIFFICULTY = 3
	CHALLENGE = 4
	BONUS = 5
	SETBACK = 6
	FORCE = 7
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type SWRPGDie struct {
	Name			string
	current_face 	int
	faces 			[]SWRPGDieResult
}

func NewSWRPGDie(dietype int) SWRPGDie{
	var die SWRPGDie
	switch dietype {
		case ABILITY:
			die = SWRPGDie{"Ability", 0, make([]SWRPGDieResult, 8)}
			die.faces[0] = SWRPGDieResult{Success: 1}
			die.faces[1] = SWRPGDieResult{Advantage: 1}
			die.faces[2] = SWRPGDieResult{Success: 1, Advantage: 1}
			die.faces[3] = SWRPGDieResult{Success: 2}
			die.faces[4] = SWRPGDieResult{Advantage: 1}
			die.faces[5] = SWRPGDieResult{Success: 1}
			die.faces[6] = SWRPGDieResult{Advantage: 2}
			die.faces[7] = SWRPGDieResult{}
		case PROFICIENCY:
			die = SWRPGDie{"Proficiency", 0, make([]SWRPGDieResult, 10)}
			die.faces[0] = SWRPGDieResult{Advantage: 2}
			die.faces[1] = SWRPGDieResult{Advantage: 1}
			die.faces[2] = SWRPGDieResult{Advantage: 2}
			die.faces[3] = SWRPGDieResult{Triumph: 1}
			die.faces[4] = SWRPGDieResult{Success: 1}
			die.faces[5] = SWRPGDieResult{Success: 1, Advantage: 1}
			die.faces[6] = SWRPGDieResult{Success: 1}
			die.faces[7] = SWRPGDieResult{Success: 1, Advantage: 1}
			die.faces[8] = SWRPGDieResult{Success: 2}
			die.faces[9] = SWRPGDieResult{}
		case DIFFICULTY:
			die = SWRPGDie{"Difficulty", 0, make([]SWRPGDieResult, 8)}
			die.faces[0] = SWRPGDieResult{Threat: 1}
			die.faces[1] = SWRPGDieResult{Failure: 1}
			die.faces[2] = SWRPGDieResult{Failure:1, Threat: 1}
			die.faces[3] = SWRPGDieResult{Threat: 1}
			die.faces[4] = SWRPGDieResult{}
			die.faces[5] = SWRPGDieResult{Threat: 2}
			die.faces[6] = SWRPGDieResult{Failure: 2}
			die.faces[7] = SWRPGDieResult{Threat: 1}
		case CHALLENGE:
			die = SWRPGDie{"Challenge", 0, make([]SWRPGDieResult, 12)}
			die.faces[0] = SWRPGDieResult{Threat: 2}
			die.faces[1] = SWRPGDieResult{Threat: 1}
			die.faces[2] = SWRPGDieResult{Threat: 2}
			die.faces[3] = SWRPGDieResult{Threat: 1}
			die.faces[4] = SWRPGDieResult{Failure: 1, Threat: 1}
			die.faces[5] = SWRPGDieResult{Failure: 1}
			die.faces[6] = SWRPGDieResult{Failure: 1, Threat: 1}
			die.faces[7] = SWRPGDieResult{Failure: 1}
			die.faces[8] = SWRPGDieResult{Failure: 2}
			die.faces[9] = SWRPGDieResult{Despair: 1}
			die.faces[10] = SWRPGDieResult{Failure: 2}
			die.faces[11] = SWRPGDieResult{}
		case BONUS:
			die = SWRPGDie{"Bonus", 0, make([]SWRPGDieResult, 6)}
			die.faces[0] = SWRPGDieResult{Success: 1, Advantage: 1}
			die.faces[1] = SWRPGDieResult{Advantage: 1}
			die.faces[2] = SWRPGDieResult{Advantage: 2}
			die.faces[3] = SWRPGDieResult{}
			die.faces[4] = SWRPGDieResult{Success: 1}
			die.faces[5] = SWRPGDieResult{}
		case SETBACK:
			die = SWRPGDie{"Setback", 0, make([]SWRPGDieResult, 6)}
			die.faces[0] = SWRPGDieResult{}
			die.faces[1] = SWRPGDieResult{}
			die.faces[2] = SWRPGDieResult{Threat: 1}
			die.faces[3] = SWRPGDieResult{Threat: 1}
			die.faces[4] = SWRPGDieResult{Failure: 1}
			die.faces[5] = SWRPGDieResult{Failure: 1}
		case FORCE:
			die = SWRPGDie{"Force", 0, make([]SWRPGDieResult, 12)}
			die.faces[0] = SWRPGDieResult{Dark: 1}
			die.faces[1] = SWRPGDieResult{Dark: 1}
			die.faces[2] = SWRPGDieResult{Dark: 1}
			die.faces[3] = SWRPGDieResult{Dark: 1}
			die.faces[4] = SWRPGDieResult{Dark: 1}
			die.faces[5] = SWRPGDieResult{Dark: 1}
			die.faces[6] = SWRPGDieResult{Light: 2}
			die.faces[7] = SWRPGDieResult{Light: 2}
			die.faces[8] = SWRPGDieResult{Light: 2}
			die.faces[9] = SWRPGDieResult{Light: 1}
			die.faces[10] = SWRPGDieResult{Light: 1}
			die.faces[11] = SWRPGDieResult{Dark: 2}
	}
	return die
}

func (d *SWRPGDie) Roll() {
	d.current_face = rand.Intn(len(d.faces))
}

func (d *SWRPGDie) GetFace(face int) SWRPGDieResult {
	return d.faces[face]
}

func (d *SWRPGDie) GetCurrentFace() SWRPGDieResult {
	return d.GetFace(d.current_face)
}

type SWRPGDieResult struct {
	Success 	int
	Failure		int
	Advantage 	int
	Threat		int
	Triumph		int
	Despair		int
	Light		int
	Dark		int
}

func (r *SWRPGDieResult) Add(a SWRPGDieResult) {
	r.Success += a.Success
	r.Failure += a.Failure
	r.Advantage += a.Advantage
	r.Threat += a.Threat
	r.Triumph += a.Triumph
	r.Despair += a.Despair
	r.Light += a.Light
	r.Dark += a.Dark
}

func (r *SWRPGDieResult) Calculate() {
	r.Success += r.Triumph
	r.Failure += r.Despair
	if r.Success >= r.Failure {
		r.Success -= r.Failure
		r.Failure = 0
	} else if r.Success < r.Failure {
		r.Failure -= r.Success
		r.Success = 0
	}
	
	if r.Advantage >= r.Threat {
		r.Advantage -= r.Threat
		r.Threat = 0
	} else if r.Advantage < r.Threat {
		r.Threat -= r.Advantage
		r.Advantage = 0
	}
}

func (r SWRPGDieResult) String() string{
	var buffer bytes.Buffer
	if(r.Success > 0) {
		buffer.WriteString("SUCCESS:")
		buffer.WriteString(strconv.Itoa(r.Success))
		buffer.WriteString(", ")
	}
	if(r.Failure > 0) {
		buffer.WriteString("FAILURE:")
		buffer.WriteString(strconv.Itoa(r.Failure))
		buffer.WriteString(", ")
	}
	if(r.Advantage > 0) {
		buffer.WriteString("ADVANTAGE:")
		buffer.WriteString(strconv.Itoa(r.Advantage))
		buffer.WriteString(", ")
	}
	if(r.Threat > 0) {
		buffer.WriteString("THREAT:")
		buffer.WriteString(strconv.Itoa(r.Threat))
		buffer.WriteString(", ")
	}
	if(r.Triumph > 0) {
		buffer.WriteString("TRIUMPH:")
		buffer.WriteString(strconv.Itoa(r.Triumph))
		buffer.WriteString(", ")
	}
	if(r.Despair > 0) {
		buffer.WriteString("DESPAIR:")
		buffer.WriteString(strconv.Itoa(r.Despair))
		buffer.WriteString(", ")
	}
	if(r.Light > 0) {
		buffer.WriteString("LIGHT:")
		buffer.WriteString(strconv.Itoa(r.Light))
		buffer.WriteString(", ")
	}
	if(r.Dark > 0) {
		buffer.WriteString("DARK:")
		buffer.WriteString(strconv.Itoa(r.Dark))
		buffer.WriteString(", ")
	}
	return strings.TrimSuffix(buffer.String(), ", ")
}