package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"example.com/conf"

	//	"fmt"
	"strings"
	"time"
)

func main() {

	var blocks [10]conf.Block
	var pBlock *conf.Block
	var tstamp string

	GVar := new(conf.Config)
	GVar.Configure("globals.json")

	for i := 0; i < len(GVar.BlockNames); i++ {
		pBlock = &blocks[i]
		pBlock.MakeBlocks(GVar.BlockNames[i])
	}

	for i := 0; i < len(GVar.BlockNames); i++ {
		pBlock = &blocks[i]
		pBlock.BlockName = GVar.BlockNames[i]
		pBlock.GiveBlocks()
	}
	//	Save Data
	//	file = "./" + file|os.O_APPEND,
	fo, er := os.OpenFile("jack.dat", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if er != nil {
		log.Fatal(er)
	}
	defer fo.Close()

	w := bufio.NewWriter(fo)

	defer w.Flush()

	fi, err := os.Open("adjust.ini")
	defer fi.Close()

	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(fi)
	for {
		str, err := r.ReadString('\n')

		if err == io.EOF || strings.HasPrefix(str, "[headerLines]") {
			break
		}

	}
	t := time.Now()
	tstamp = t.Format(time.ANSIC) + "\n"
	w.WriteString(tstamp)

	for {
		str, err := r.ReadString('\n')
		if err == io.EOF || strings.HasPrefix(str, "[") {
			break
		}

		w.WriteString(str)
	}

	for j := 0; j < len(GVar.BlockNames); j++ {

		for i := 0; i < len(blocks[j].Trials); i++ {
			s := blocks[j].Trials[i].GetLine()

			w.WriteString(s)
		}
	}

}
