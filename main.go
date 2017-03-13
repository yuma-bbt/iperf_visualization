package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
	"math/rand"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

func main(){

	var splited []string
	var time []string
	var bandwidth []string
	var i int
	i=0


	fp, err := os.Open("test")
	if(err != nil){
		panic(err)
	}

	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 4096)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		splited = strings.Split(string(line),",")
		time = append (time,splited[0])
		bandwidth = append(bandwidth,splited[1])
		fmt.Print(time[i])
		fmt.Print(":")
		fmt.Println(bandwidth[i])
		i++

	}
	ploter(time,bandwidth)

}

func ploter(time string[], bandwidth string[]){

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "graf"
	p.X.Label.Text = "Time(hour)"
	p.Y.Label.Text = "Bandwidth(Mbps)"

	plotutil.AddLinePoints(p, "",plotter.XYs{time,bandwidth})
	width :=4.0
	height :=4.0
	p.Save(width, height, "line_sample.png")
}


//	if err := plotutil.AddLinePoints(p, "test", point(len(time)); err != nil {
//		panic(err)
//	}
//
//	if err := p.Save(5*vg.Inch, 5*vg.Inch, "graf.png"); err != nil {
//		panic(err)
//	}
//}
//
//
//func point(time , bandwidth ,n int) plotter.XYs {
//
//	pts := make(plotter.XYs, n)
//	data_length = len(time)
//
//	for i := 0; i < data_length; i++ {
//		pts[i].X = int(time[i])
//		pts[i].Y = int(bandwidth[i])
//	}
//	return pts
//}
