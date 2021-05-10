package process

import (
	"fmt"
	"../ExtractStrategy"

)

func (ip *ImageProcessing)  SetProcessStrategy(ps processStrategy) {
	ip.Strategy = ps
}

func (ip *ImageProcessing) GetImages(ie *extract.ImageExtractor){
	(*ip).FilteredImages = (*ie).Images
	(*ip).Readinfo = (*ie).Readinfo
}

func (ip *ImageProcessing) ProcessGroup(print bool){

	(*ip).Strategy.Allocate(ip)

	for i := 0; i < (*ip).Readinfo.SizeData; i++ {
		if print{
			fmt.Println("processing data: ",i+1)
		}
		(*ip).Strategy.Process(ip,i)
	}
}