package loader

import (
	"encoding/csv"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"

	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (l *Loader) Search(fil string) []getlin.Vector {
	var err error

	var dst string
	{
		dst = filepath.Join(l.bas, l.rep, "/data/", fil)
	}

	var des *os.File
	{
		des, err = os.Open(dst)
		if err != nil {
			panic(err)
		}
	}

	var rdr *csv.Reader
	{
		rdr = csv.NewReader(des)
	}

	var all [][]string
	{
		all, err = rdr.ReadAll()
		if err != nil {
			panic(err)
		}
	}

	var sta int
	{
		sta = rand.Intn(len(all))
	}

	if sta+l.bat > len(all)-1 {
		sta = (len(all) - 1) - l.bat
	}

	var end int
	{
		end = sta + l.bat
	}

	var vec []getlin.Vector
	for _, x := range all[sta:end] {
		v := vector.New(vector.Config{
			// Ind: musind(x[0]),
			Inp: musinp(x[1:]),
			Out: musout(x[0]),
		})
		vec = append(vec, v)
	}

	return vec
}

func musind(str string) [][]uint8 {
	var tru int
	{
		tru = musint(str)
	}

	var ran int
	for {
		{
			ran = rand.Intn(10)
		}

		if ran != tru {
			break
		}
	}

	var lis []uint8
	{
		lis = make([]uint8, 10)
	}

	for i := range lis {
		if i == tru || i == ran {
			lis[i] = 1
		}
	}

	return [][]uint8{
		lis,
	}
}

func musinp(str []string) []uint8 {
	var bit []uint8

	for _, x := range str {
		var num int
		{
			num = musint(x)
		}

		if num <= 127 {
			bit = append(bit, 0)
		}

		if num >= 128 {
			bit = append(bit, 1)
		}
	}

	return bit
}

func musint(str string) int {
	num, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		panic(err)
	}

	return int(num)
}

func musout(str string) []uint8 {
	var bit []uint8
	{
		bit = make([]uint8, 10)
	}

	{
		bit[musint(str)] = 1
	}

	return bit
}
