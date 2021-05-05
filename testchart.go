package main

import (
	"fmt"
	"os"
	"github.com/wcharczuk/go-chart" //expose package "chart"
)

func main() {
	xvalues := []float64{0.0,1.0,2.0,3.0,4.0,5.0}
	yvalues := []float64{1.0,1.0,2.0,3.0,5.0,8.0}

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xvalues,
				YValues: yvalues,
			},
		},
	}

	err := graph.Render(chart.PNG,os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}