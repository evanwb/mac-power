package main

import (
	"fmt"
	"mac_power/power"

	/* 	"github.com/guptarohit/asciigraph"
	 */

)

func main() {

	fmt.Println("which to measure")
	fmt.Println("1. CPU")
	fmt.Println("2. GPU")

	var input int
	fmt.Scanln(&input)
	fmt.Print("\033[H\033[2J")

	power_channel := make(chan float64)

	if input == 1 {
		go func() {
			for {
				power.GetCPU(power_channel)
			}

		}()
	} else {
		go func() {
			for i := 0; i < 10; i++ {
				power.GetGPU(power_channel)
			}
		}()
	}

	data := []float64{}
	labels := []string{}

	i := 0


	for num := range power_channel {
		data = append(data, num/1000)
		labels = append(labels, fmt.Sprintf("%v", i))
		i++

		var out string;
		if num > 1000 {

			out = fmt.Sprintf("%.2f W", num/1000)
		} else {
			out = fmt.Sprintf("%d mW", int(num))
		}
		if input ==1 {
			fmt.Println("current cpu power:", out )
		} else {
			fmt.Println("current gpu power:", out )
		}

		//asciigraph.Height(20)
		/* var caption string
		if input ==1 {
			caption= "cpu power (W)"
		} else {
			caption= "gpu power (W)"
		} */
		
		/* graph := asciigraph.Plot(data,asciigraph.Caption(caption), asciigraph.Height(30))
		//fmt.Println("\033[0;0H")
		cmd:= exec.Command("clear && printf '\\e[3J'")
		cmd.Run()
		fmt.Print("\033[H\033[2J")

		    fmt.Println(graph) */

		
		

	}

}
