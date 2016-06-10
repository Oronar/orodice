package orodice

import "bytes"
import "strconv"
import "sort"
import "strings"
import "fmt"

const DICE_LIMIT = 15

type NumberDicePool struct {
	pool		[]NumberDie
	constant	int
}

func NewNumberDicePool() NumberDicePool{
	pool := NumberDicePool{make([]NumberDie,0), 0}
	return pool
}

func (dp *NumberDicePool) Parse(input string) error{
	dice := strings.Split(input, "+")
	for _, element := range dice {
		if !strings.Contains(element, "d") {
			constant, err := strconv.Atoi(element)
			if err != nil {
				return fmt.Errorf("ERROR: Could not parse constant (%v)", err)
			}
			dp.constant += constant
			continue
		}
		split := strings.SplitN(element, "d", 2)
		number,err := strconv.Atoi(split[0])
		if(err != nil) {
			return fmt.Errorf("ERROR: Could not parse number of dice (%v)", err)
		}
		dietype,err := strconv.Atoi(split[1])
		if(err != nil) {
			return fmt.Errorf("ERROR: Could not parse type of die (%v)", err)
		}
		for i := 0; i<number; i++ {
			dp.Add(NewNumberDie(dietype))
		}
	}
	return nil
}

func (dp *NumberDicePool) Add(die NumberDie) {
	dp.pool = append(dp.pool, die)
}

func (dp *NumberDicePool) Roll() {
	for i := 0; i< len(dp.pool); i++ {
		dp.pool[i].Roll()
	}
}

type Result struct {
	dietype 	int
	dievalue	int
}

type ResultList []Result

func (slice ResultList) Len() int {
	return len(slice)
}

func (slice ResultList) Less(i, j int) bool {
	return slice[i].dietype < slice[j].dietype
}

func (slice ResultList) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (dp NumberDicePool) String() string {
	result := NumberDieResult{}
	resultslist := ResultList{}
	for i := 0;i < len(dp.pool);i++ {
		result.Add(dp.pool[i].GetCurrentFace())
		rs := Result{len(dp.pool[i].faces), dp.pool[i].GetCurrentFace().Value}
		resultslist = append(resultslist,rs)
	}
	sort.Sort(resultslist)

	var buffer bytes.Buffer
	currenttype := 0
	for i := 0; i<len(resultslist); i++ {
		if(resultslist[i].dietype != currenttype && i < DICE_LIMIT) {
			if(currenttype != 0) {
				buffer.WriteString("\n")
			}
			buffer.WriteString("d")
			buffer.WriteString(strconv.Itoa(resultslist[i].dietype))
			buffer.WriteString(": ")
			currenttype = resultslist[i].dietype
		}
		if i < DICE_LIMIT {
			buffer.WriteString(strconv.Itoa(resultslist[i].dievalue))
			buffer.WriteString(", ")
		} else if i == DICE_LIMIT {
			buffer.WriteString("\nToo many dice...truncating.")
		}
	}
		
	buffer.WriteString("\nTotal: ")
	buffer.WriteString(result.String())
	if dp.constant > 0 {
		buffer.WriteString(" + ")
		buffer.WriteString(strconv.Itoa(dp.constant))
		buffer.WriteString(" = ")
		buffer.WriteString(strconv.Itoa(result.Value + dp.constant))
	}
	
	return strings.Replace(buffer.String(), ", \n", "\n", -1)	
}
	