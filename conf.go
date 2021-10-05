 package conf

import (

	"fmt"
	"math/rand"
	"os"
	"encoding/json"
)

type Config struct {
	Pdist      int
	Height     int
	Wide       int
	MaxTrials  int
	Practice   int
	BreakDelay int
	EndDelay   int
	Iti        int
	BlockNames []string
}

type Trial struct {
	Condition   int
	Top         int
	Mid         int
	Bottom      int
	Orientation int
	NAdjust     int
	Pse         float32
	Standard    int
	Gap         int
}

type Block struct {
	BlockName string
	Trials    []Trial
}


func shuffle(slTrials []Trial) {
	//  randomize trials in this block
	var imax int
	imax = len(slTrials)
	for i := 0; i < imax; i++ {
		k := rand.Int() % (imax - i)
		k = k + i
		slTrials[i], slTrials[k] = slTrials[k], slTrials[i]
	}
}

/* MakeBlocks needs to be called from the main with consecutive
elements of an array of Block.  tNum, above, is like a
C static and is a index into TrialPool.  It is used to create
slices, Block.Trials, from TrialPool.
*/

func (Blk *Block) MakeBlocks(fname string) {
	fname += ".json"
	f, err := os.Open(fname)
	check(err)
	defer f.Close()
	b1 := make([]byte, 5000)
	n1, err := f.Read(b1)
	check(err)
	b3 := b1[0:n1]
	error := json.Unmarshal(b3, &Blk)
	check(error)
	shuffle(Blk.Trials)
}

func (Blk *Block) GiveBlocks() {

	for i := 0; i < len(Blk.Trials); i++ {
		Blk.Trials[i].Pse = rand.Float32() * 3.0
		Blk.Trials[i].NAdjust = rand.Int() % 20
		Blk.Trials[i].Standard = 200
	}

}

func (trial Trial) GetLine() (s string) {

	s = fmt.Sprintf("%d %d %d %d %d %f %d %d\n",
		trial.Condition,
		trial.Top,
		trial.Mid,
		trial.Bottom,
		trial.Orientation,
		trial.Pse,
		trial.NAdjust,
		trial.Standard)
	return s
}


func (conf *Config) Configure(fname string) {
	
	f, err := os.Open(fname)
	check(err)

	b1 := make([]byte, 800)
	n1, err := f.Read(b1)
//	fmt.Printf("number of char %v: %v\n", n1, string(b1))
	check(err)
	defer f.Close()
	b3 := b1[0:n1]
	error := json.Unmarshal(b3, &conf)
	check(error)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
