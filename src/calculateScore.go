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

type Interest struct {
	Lv1 []string
	Lv2 []string
	Lv3 []string
}

type Person struct {
	Name string
}

func main() {
	var select1 = []string{"food", "thai food", "fried rice"}
	var select2 = []string{"food", "thai food", "kraprao"}
	var selectList = [][]string{select1, select2}
	
	var interest = groupInterest(selectList)
	var resultList = execute(interest)
	
	for i := range resultList {
		fmt.Println( "Name is " + resultList[i].Name + ", score is " + strconv.Itoa(resultList[i].Value) );
	}
}

func groupInterest(selectList [][]string) Interest {
	var interestLv1Map = make(map[string]int)
	var interestLv2Map = make(map[string]int)
	var interestLv3Map = make(map[string]int)

	for i := range selectList {
		var selectTmp = selectList[i]
		interestLv1Map[selectTmp[0]] = 1;
		interestLv2Map[selectTmp[1]] = 1;
		interestLv3Map[selectTmp[2]] = 1;
	}
	
	var interestLv1 = createListFromMap(interestLv1Map)
	var interestLv2 = createListFromMap(interestLv2Map)
	var interestLv3 = createListFromMap(interestLv3Map)

	//var interest = Interest{[]string{"food"}, []string{"thai food"}, []string{"fried rice", "kraprao"}}
	var interest = Interest{interestLv1, interestLv2, interestLv3}
	fmt.Printf("Global interest is ")
	fmt.Println(interest)
	
	return interest;
}

func createListFromMap(myMap map[string]int) []string {
	var myArr = make([]string, len(myMap))
	var i = 0
	for k, _ := range myMap {
		myArr[i] = k
		i += 1
	}
	return myArr
}

func calculateScoreTotal(interest Interest) map[string]int {
	var scoreMap = make(map[string]int)

	calculateScore(scoreMap, interest.Lv1, 1, 1);
	calculateScore(scoreMap, interest.Lv2, 2, 2);
	calculateScore(scoreMap, interest.Lv3, 3, 5);
	
	return scoreMap
}

func calculateScore(scoreMap map[string]int, categoryList []string, level int, weight int) {
	fmt.Printf("interest is ")
	fmt.Println(categoryList)
	for i := range categoryList {
		var category = categoryList[i];
		var personList = getPersonByInterest(level, category)
		for j := range personList {
			fmt.Println("  friend name is " + personList[j])
			scoreMap[personList[j]] += weight
		}
	}
}

func getPersonByInterest(level int, category string) []string {
	// use level as column name to retrieve person list from table
	if (level == 1) {
		if (category == "food") {
			return []string{"a", "b"}
		}
	} else if (level == 2) {
		if (category == "thai food") {
			return []string{"a", "c"}
		}
	} else if (level == 3) {
		if (category == "kraprao") {
			return []string{"d"}
		}
	}
	return []string{}
}

func execute(interest Interest) PairList {
	var scoreMap = calculateScoreTotal(interest)
	return getSortListFromMap(scoreMap)
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
