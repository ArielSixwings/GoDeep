package process

import(
	"fmt"
	"../basicdata"
	"strconv"
	"strings"
)

func (sp *StatisticProcessing) ConvertData(){
	
	split := make([]string,len((*sp).Texts[0]))
	(*sp).Information = make([]cartesian.Features,890) //temporary solution
	
	firts := true
	second := true

	for i := 0; i < len((*sp).Texts[0]); i++ {
		
		if second{
			second = false
			continue
		} else{
			if firts{
				firts = false
				continue				
			} else{
				if len((*sp).Texts[0][i]) == 0{
					break
				}
			}
		}

		split = strings.Split((*sp).Texts[0][i], ",")

		(*sp).Information[i-2].Features[0] , _= strconv.ParseFloat(split[2], 64)
		(*sp).Information[i-2].Features[2] , _= strconv.ParseFloat(split[6], 64)

		if split[5] == "male" {
			(*sp).Information[i-2].Features[1] = 1
		} else{
			(*sp).Information[i-2].Features[1] = 0
		}

		if split[1] == "1"{
			(*sp).Information[i-2].Label = "Survived"
		} else{
			(*sp).Information[i-2].Label = "Do not Survived"		
		}
	}

	(*sp).SortData()	
}

func (sp *StatisticProcessing) SortData() {

	sortedData := make([]cartesian.Features,len((*sp).Information))

	LabelA := (*sp).Information[1].Label

	lowerVectorIndex := 0
	upperVectorIndex := 0

	for i := 0; i < len((*sp).Information); i++ {
		if (*sp).Information[i].Label == LabelA{
			sortedData[len((*sp).Information) - 1 - upperVectorIndex] = (*sp).Information[i]
			upperVectorIndex++
		} else {
			sortedData[lowerVectorIndex] = (*sp).Information[i]
			lowerVectorIndex++
		}
	}

	for i := 0; i < len(sp.Information); i++ {

		if len(sp.Information[i].Label) == 0{
			break
		}
	}

	sp.Information = sortedData
}

func (sp StatisticProcessing) PrintFeatures() {
	for i := 0; i < len(sp.Information); i++ {
		if len(sp.Information[i].Label) == 0{
			break
		}
		fmt.Println(i,"     ",sp.Information[i])
	}
}