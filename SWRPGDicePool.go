package orodice

import "bytes"
import "strings"
import "strconv"
import "fmt"

type SWRPGDicePool struct {
	pool	[]SWRPGDie
}

func NewSWRPGDicePool() SWRPGDicePool{
	pool := SWRPGDicePool{make([]SWRPGDie,0)}
	return pool
}

func (dp *SWRPGDicePool) Parse(input string) error{
	ParseMap := map[string]int{
		"a": 1,
		"p": 2,
		"d": 3,
		"c": 4,
		"b": 5,
		"s": 6,
		"f": 7,
	}
	dice := strings.Split(input, "+")
	for _, element := range dice {
		if !strings.Contains(element, "d") {
			constant, err := strconv.Atoi(element)
			if err != nil {
				return fmt.Errorf("ERROR: Could not parse constant success (%v)", err)
			}
			die := SWRPGDie{"Const Success", 0, make([]SWRPGDieResult, 1)}
			die.faces[0] = SWRPGDieResult{Success: constant}
			dp.Add(die)	
			continue
		}
		split := strings.SplitN(element, "d", 2)
		number,err := strconv.Atoi(split[0])
		if(err != nil) {
			return fmt.Errorf("ERROR: Could not parse number of dice (%v)", err)
		}
		dietype := split[1]
		if(ParseMap[dietype] == 0) {
			return fmt.Errorf("ERROR: Could not parse type of die (%v)", err)
		}
		for i := 0; i<number; i++ {
			dp.Add(NewSWRPGDie(ParseMap[dietype]))
		}	
	}
	
	return nil
}

func (dp *SWRPGDicePool) Add(die SWRPGDie) {
	dp.pool = append(dp.pool, die)
}

func (dp *SWRPGDicePool) Roll() {
	for i := 0; i< len(dp.pool); i++ {
		dp.pool[i].Roll()
	}
}

func (dp SWRPGDicePool) String() string {
	var buffer bytes.Buffer
	result := SWRPGDieResult{}
	for i := 0;i < len(dp.pool);i++ {
		result.Add(dp.pool[i].GetCurrentFace())
		if i < DICE_LIMIT {
			buffer.WriteString(dp.pool[i].Name)
			buffer.WriteString(": ")
			buffer.WriteString(dp.pool[i].GetCurrentFace().String())
			buffer.WriteString("\n")
		} else if i == DICE_LIMIT {
			buffer.WriteString("Too many dice...truncating.\n")
		}
	}
	result.Calculate()
	buffer.WriteString("Total: ")
	buffer.WriteString(result.String())
	
	return buffer.String()
}
	