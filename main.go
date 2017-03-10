package main

import (
	"fmt"
	"strings"
	_"reflect"
	"flag"
	"math/rand"
	"bufio"
	"io"
	"os"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

var (
       filename = flag.String("f","test","specify the data file" )
)


func main() {
	data_read("test")
	//ploter()
}

func data_read(filename string) {

        var fp *os.File
        var err error
        fp, err  = os.Open(filename)
	var data []string
	var splited []string
	var i int
	i=0
	//var hoge = []string{}


        if err != nil{
                panic(err)
                }
                defer fp.Close()
        reader := bufio.NewReaderSize(fp, 4096)
        for line := ""; err == nil; line, err = reader.ReadString('\n') {
		data = append(data,line)
		splited = strings.Split(data[i],"\t")
		//fmt.Print(data[i])
		fmt.Println(splited)
		//fmt.Println(splited[-1])
		i++
        }
        if err != io.EOF {
                panic(err)
        }
}

func ploter(){
	rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "graf"
	p.X.Label.Text = "Time(hour)"
	p.Y.Label.Text = "Bandwidth(Mbps)"

	if err := plotutil.AddLinePoints(p, "test", point(5)); err != nil {
		panic(err)
	}

	if err := p.Save(5*vg.Inch, 5*vg.Inch, "graf.png"); err != nil {
		panic(err)
	}
}



func point(n int) plotter.XYs {

	pts := make(plotter.XYs, n)

	for i := 0; i < 5; i++ {
		pts[i].X = float64(i)
		pts[i].Y = float64(i * i)
	}
	return pts
}
