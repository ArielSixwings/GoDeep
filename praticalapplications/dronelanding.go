package main

import (
	"../code/imageprocessing"
	"../code/generalizeimage"
	"../code/nonparametric"
	"gocv.io/x/gocv"
	"fmt"
	//"math"
)

func main() {
	
	var dataset generalizeimage.Labelfeatures
	
	/*size of train and know groups*/
	var size int
	var knowsize int
	var trainsize int
	
	/*normalize flags*/
	var normtype gocv.NormType = gocv.NormMinMax

	/*calc sizes*/
	size  = imageprocessing.FolderLength("../code/imageprocessing/Images/danger")
	trainsize = 20 //int(size/2.5)
	knowsize = size - trainsize

	/*set labelsizes*/
	knowls := make([]generalizeimage.Sizelabel,3)
	trainls := make([]generalizeimage.Sizelabel,3)

	for i := 0; i < 3; i++ {
		knowls[i].Size_l  = knowsize
		trainls[i].Size_l = trainsize	
	}

	knowls[0].Label  = "danger"
	trainls[0].Label = "danger"

	knowls[1].Label  = "asphalt"
	trainls[1].Label = "asphalt"

	knowls[2].Label  = "grass"
	trainls[2].Label = "grass"

	/* Know images and features allocation*/
	knowImages 			:= make([]gocv.Mat,3*knowsize)	// 	images
	
	knowGLCMs 			:= make([]gocv.Mat,3*knowsize)	// 	GLCMs
	normalizedknow	 	:= make([]gocv.Mat,3*knowsize)	// 	normalizedGLCMs
	/*Know gclm and normalized glcm internal allocation*/
	for i := 0; i < 3*knowsize; i++ {
		knowGLCMs[i]			= gocv.NewMatWithSize(256, 256, gocv.MatTypeCV8U)	
		normalizedknow[i]		= gocv.NewMat()
	}
	/*Know Features*/
	knowEnergys			:= make([]float64,3*knowsize)	// 	Energy
	knowCorrelations	:= make([]float64,3*knowsize)	// 	Correlation
	knowContrasts		:= make([]float64,3*knowsize)	// 	Contrast

	/* Train images and features allocation*/
	trainImages 		:= make([]gocv.Mat,3*trainsize)	// 	images
	
	trainGLCMs 			:= make([]gocv.Mat,3*trainsize)	// 	GLCMs
	normalizedtrain		:= make([]gocv.Mat,3*trainsize)	// 	normalizedGLCMs
	/*Train gclm and normalized glcm internal allocation*/
	for i := 0; i < 3*trainsize; i++ {
		trainGLCMs[i]			= gocv.NewMatWithSize(256, 256, gocv.MatTypeCV8U)	
		normalizedtrain[i]		= gocv.NewMat()
	}
	/*Train Features*/
	trainEnergys		:= make([]float64,3*trainsize)	// 	Energy
	trainCorrelations	:= make([]float64,3*trainsize)	// 	Correlation
	trainContrasts		:= make([]float64,3*trainsize)	// 	Contrast	
	
	/*temporary set of images that will be used to read each folder*/
	auxImages 			:= make([]gocv.Mat,size)

	/*read and separe each group of images*/
	imageprocessing.ReadFolder(auxImages,"../code/imageprocessing/Images/danger",true,false,false)
	
	for i := 0; i < size; i++ {
		if i < trainsize{
			trainImages[i] = auxImages[i]
		} else{
			knowImages[i-trainsize] = auxImages[i]
		}
	}
	
	imageprocessing.ReadFolder(auxImages,"../code/imageprocessing/Images/asphalt",true,false,false)
	for i := 0; i < size; i++ {
		if i < trainsize{
			trainImages[i+trainsize] = auxImages[i]
		} else{
			knowImages[i+(knowsize-trainsize)] = auxImages[i]
		}
	}
	
	imageprocessing.ReadFolder(auxImages,"../code/imageprocessing/Images/grass",true,false,false)
	for i := 0; i < size; i++ {
		if i < trainsize{
			trainImages[i+(2*trainsize)] = auxImages[i]
		} else{
			knowImages[i+((2*knowsize)-trainsize)] = auxImages[i]
		}
	}	

	/*compute GLCMs and them the normalized GLCM*/
	imageprocessing.GroupGLCM(knowImages, &knowGLCMs, true, true)
	for i := 0; i < 3*knowsize; i++ {
		gocv.Normalize(knowGLCMs[i], &normalizedknow[i], 0.0, 255.0, normtype )		
	}

	
	imageprocessing.GroupGLCM(trainImages, &trainGLCMs, true, true)
	for i := 0; i < 3*trainsize; i++ {
		gocv.Normalize(trainGLCMs[i], &normalizedtrain[i], 0.0, 255.0, normtype )

	}

	/*Extract the features*/
	imageprocessing.GroupFeature(&normalizedknow,knowEnergys,imageprocessing.EnergyFeature, true)
	imageprocessing.GroupFeature(&normalizedknow,knowCorrelations,imageprocessing.CorrelationFeature, true)
	imageprocessing.GroupFeature(&normalizedknow,knowContrasts,imageprocessing.ContrastFeature, true)

	imageprocessing.GroupFeature(&normalizedtrain,trainEnergys,imageprocessing.EnergyFeature, true)
	imageprocessing.GroupFeature(&normalizedtrain,trainCorrelations,imageprocessing.CorrelationFeature, true)
	imageprocessing.GroupFeature(&normalizedtrain,trainContrasts,imageprocessing.ContrastFeature, true)


	fmt.Println("generalizing know data set")
	generalizeimage.Generalize_for_nonparametric(&dataset, knowEnergys, knowCorrelations, knowContrasts,knowls,generalizeimage.Knowflag,90)
	
	fmt.Println("generalizing train data set")
	generalizeimage.Generalize_for_nonparametric(&dataset, trainEnergys, trainCorrelations, trainContrasts,trainls,generalizeimage.Trainflag,60)

	dataset.Printfeatures()

	nonparametric.KNN(&dataset,5)

	dataset.Printresults()
}
