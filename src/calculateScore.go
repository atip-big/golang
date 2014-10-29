package main
 
import (
	"fmt"
	"strconv"
	"sort"
)

// A data structure to hold a key/value pair.
type Pair struct {
  Name string
  Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func main() {
	var interestList = []string{"football", "pingpong"}
	var pairList = calculateScore(interestList)
	
	for i := range pairList {
		fmt.Println( "Name is " + pairList[i].Name + ", score is " + strconv.Itoa(pairList[i].Value) );
	}
}

func calculateScore(interestList []string) PairList {
	var scoreMap = make(map[string]int)
	
	for i := range interestList {
		fmt.Println("interest is " + interestList[i])
		
		var personList = getPersonByInterest(interestList[i])
		for j := range personList {
			fmt.Println("  friend name is " + personList[j])
			
			scoreMap[personList[j]] += 1
		}
	}
	
	return getSortListFromMap(scoreMap)
}

func getPersonByInterest(interest string) []string {
	if (interest == "football") {
		return []string{"c", "b"}
	} else if (interest == "pingpong") {
		return []string{"a", "c"}
	}
	return []string{}
}

func getSortListFromMap(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i += 1
	}
	sort.Sort(sort.Reverse(p))
	return p
}

