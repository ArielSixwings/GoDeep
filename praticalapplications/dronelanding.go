package main

import (
	"../src/imageprocessing"
	"../src/learnstrategy/nonparametric"
	"../src/learnstrategy"
	"../src/basicdata"
	"gocv.io/x/gocv"
	"fmt"
	//"math"
)

func main() {
	
	var dataset learnstrategy.DataSet
	
	/*size of test and train groups*/
	var size int
	var trainsize int
	var testsize int
	
	/*normalize flags*/
	var normtype gocv.NormType = gocv.NormMinMax

	/*calc sizes*/
	size  = imageprocessing.FolderLength("../src/imageprocessing/Images/danger")
	testsize = 25//int(size/2.5)
	trainsize = size - testsize

	/*set labelsizes*/
	trainls := make([]cartesian.Sizelabel,3)
	testls := make([]cartesian.Sizelabel,3)

	for i := 0; i < 3; i++ {
		trainls[i].Size_l  = trainsize
		testls[i].Size_l = testsize	
	}

	trainls[0].Label  = "danger"
	testls[0].Label = "danger"

	trainls[1].Label  = "asphalt"
	testls[1].Label = "asphalt"

	trainls[2].Label  = "grass"
	testls[2].Label = "grass"

	/* train images and features allocation*/
	trainImages 			:= make([]gocv.Mat,3*trainsize)	// 	images
	
	trainGLCMs 			:= make([]gocv.Mat,3*trainsize)	// 	GLCMs
	normalizedtrain	 	:= make([]gocv.Mat,3*trainsize)	// 	normalizedGLCMs
	/*train gclm and normalized glcm internal allocation*/
	for i := 0; i < 3*trainsize; i++ {
		trainGLCMs[i]			= gocv.NewMatWithSize(256, 256, gocv.MatTypeCV8U)	
		normalizedtrain[i]		= gocv.NewMat()
	}
	/*train Features*/
	trainEnergys			:= make([]float64,3*trainsize)	// 	Energy
	trainCorrelations	:= make([]float64,3*trainsize)	// 	Correlation
	trainContrasts		:= make([]float64,3*trainsize)	// 	Contrast

	/* test images and features allocation*/
	testImages 		:= make([]gocv.Mat,3*testsize)	// 	images
	
	testGLCMs 			:= make([]gocv.Mat,3*testsize)	// 	GLCMs
	normalizedtest		:= make([]gocv.Mat,3*testsize)	// 	normalizedGLCMs
	/*test gclm and normalized glcm internal allocation*/
	for i := 0; i < 3*testsize; i++ {
		testGLCMs[i]			= gocv.NewMatWithSize(256, 256, gocv.MatTypeCV8U)	
		normalizedtest[i]		= gocv.NewMat()
	}
	/*test Features*/
	testEnergys		:= make([]float64,3*testsize)	// 	Energy
	testCorrelations	:= make([]float64,3*testsize)	// 	Correlation
	testContrasts		:= make([]float64,3*testsize)	// 	Contrast	
	
	/*temporary set of images that will be used to read each folder*/
	auxImages 			:= make([]gocv.Mat,size)

	/*read and separe each group of images*/
	fmt.Println("Reading danger folder")
	imageprocessing.ReadFolder(auxImages,"../src/imageprocessing/Images/danger",false,true,false)
	
	for i := 0; i < size; i++ {
		if i < testsize{
			testImages[i] = auxImages[i]
		} else{
			trainImages[i-testsize] = auxImages[i]
		}
	}
	
	fmt.Println("Reading asphalt folder")
	imageprocessing.ReadFolder(auxImages,"../src/imageprocessing/Images/asphalt",false,true,false)
	for i := 0; i < size; i++ {
		if i < testsize{
			testImages[i+testsize] = auxImages[i]
		} else{
			trainImages[i+(trainsize-testsize)] = auxImages[i]
		}
	}
	
	fmt.Println("Reading grass folder")
	imageprocessing.ReadFolder(auxImages,"../src/imageprocessing/Images/grass",false,true,false)
	for i := 0; i < size; i++ {
		if i < testsize{
			testImages[i+(2*testsize)] = auxImages[i]
		} else{
			trainImages[i+((2*trainsize)-testsize)] = auxImages[i]
		}
	}

	/*compute GLCMs and them the normalized GLCM*/
	fmt.Println("Computing train GLCMs")
	imageprocessing.GroupGLCM(trainImages, &trainGLCMs, false, true)
	for i := 0; i < 3*trainsize; i++ {
		gocv.Normalize(trainGLCMs[i], &normalizedtrain[i], 0.0, 255.0, normtype )		
	}

	fmt.Println("Computing test GLCMs")
	imageprocessing.GroupGLCM(testImages, &testGLCMs, false, true)
	for i := 0; i < 3*testsize; i++ {
		gocv.Normalize(testGLCMs[i], &normalizedtest[i], 0.0, 255.0, normtype )

	}

	/*Extract the features*/
	fmt.Println("Computing train features")
	imageprocessing.GroupFeature(&normalizedtrain,trainEnergys,imageprocessing.EnergyFeature, false)
	imageprocessing.GroupFeature(&normalizedtrain,trainCorrelations,imageprocessing.CorrelationFeature, false)
	imageprocessing.GroupFeature(&normalizedtrain,trainContrasts,imageprocessing.ContrastFeature, false)

	fmt.Println("Computing test features")
	imageprocessing.GroupFeature(&normalizedtest,testEnergys,imageprocessing.EnergyFeature, false)
	imageprocessing.GroupFeature(&normalizedtest,testCorrelations,imageprocessing.CorrelationFeature, false)
	imageprocessing.GroupFeature(&normalizedtest,testContrasts,imageprocessing.ContrastFeature, false)


	fmt.Println("Generalizing train data set")
	dataset.Build(trainEnergys, trainCorrelations, trainContrasts,trainls,learnstrategy.Trainflag,3*trainsize)
	
	fmt.Println("Generalizing test data set")
	dataset.Build(testEnergys, testCorrelations, testContrasts,testls,learnstrategy.Testflag,3*testsize)

	// fmt.Println("Computing centroid")
	// dataset.Centroid()

	// fmt.Println("Conputing radius")
	// dataset.Calcradius()

	// fmt.Println("Conputing CalcCenterdistance")
	// dataset.CalcCenterdistance()

	// fmt.Println("Filtering data set")
	// dataset.Filterdataset(dataset.MinCaoszoneRule)	

	fmt.Println("Calling Kmeans")
	Kmeans := &nonparametric.Kmeans{} 	// lfu := &lfu{}
	dataset.SetLearnStrategy(Kmeans) 	// cache.setEvictionAlgo(lru)
	dataset.ProcessLearn() 	//nonparametric.KNN(&dataset,3)
	dataset.Printresults()
    
	fmt.Println("Calling KNN")
	knn := &nonparametric.Knn{} 	// lfu := &lfu{}
	dataset.SetLearnStrategy(knn) 	// cache.setEvictionAlgo(lru)
	dataset.ProcessLearn() 	//nonparametric.KNN(&dataset,3)
	dataset.Printresults()

	//dataset.GroupCenterdists()
}