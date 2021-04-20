package imageextractor

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"errors"
	"gocv.io/x/gocv"
	"../../basicdata"
)
func (ie *ImageExtractor) Allocate(size int){
	
	(*ie).Images = make([]gocv.Mat,size)
	
	for i := 0; i < size; i++ {
		(*ie).Images[i] = gocv.NewMat()
	}
}
/**
 * [ReadImage description: read an image following the parameters]
 * @param {[type]} Image     gocv.Mat [the image]
 * @param {[type]} path      string   [the path to the image]
 * @param {[type]} show      bool     [if it is true show the image]
 * @param {[type]} colorfull bool     [if it is true read the image as a 3 chanel rbg]
 * @return {[type]}   gocv.Mat        [the readed image]
 */
func (ie *ImageExtractor)  ReadImage(path string, show bool, colorfull bool, i int){

	ImagePath := filepath.Join(path) //set path to the base image
	
	if colorfull {
		(*ie).Images[i] = gocv.IMRead(ImagePath, gocv.IMReadUnchanged) //read the base image as as RGB

	} else {
		(*ie).Images[i] = gocv.IMRead(ImagePath, gocv.IMReadGrayScale) //read the base image in grayscale

	}

	if show {
		ShowImage("And this is yout image", (*ie).Images[i], 100)

	}
}

/**
 * [ReadFolder description: read all images at some folder]
 * @param {[type]} Images    *[]gocv.Mat [An Array of gocv.Mat that will be used to contain the images of the folder]
 * @param {[type]} folder    string      [folder name]
 * @param {[type]} print     bool        [if its true, print the names]
 * @param {[type]} show      bool        [if its true, show the images]
 * @param {[type]} colorfull bool        [if its is true take a 3 chanel rbg image]
 */
func (ie *ImageExtractor) ReadFolder(folder string, print bool, show bool, colorfull bool,index ...int) int{
	
	var files []string
	var first bool = true
	var i int

	err := filepath.Walk(folder, visit(&files))
	if err != nil {
		panic(err)
	} else {
		if len((*ie).Images) == 0{
			(*ie).Allocate(3*(len(files)-1)) 	//temporary solution
		}
	}

	for _, file := range files {

		if first {
			if len(index) > 0{
				i = index[0]
			}else{
				i = 0
			}
			
			fmt.Println(file)
			first = false
			continue
		}
		if print {

			fmt.Println("geting image:     ", file)

		}
		(*ie).ReadImage(file, show, colorfull,i)
		i++
	}
	return len(files)-1
}

/**
 * [ShowImage description: Shows the image]
 * @param {[type]} Menssage string   [menssage in the window]
 * @param {[type]} Image    gocv.Mat [image to be showed]
 * @param {[type]} time     int      [time of the window]
 */
func ShowImage(Menssage string, Image gocv.Mat, time int) {
	window := gocv.NewWindow(Menssage) //basic window
	window.IMShow(Image)               //show the image
	window.WaitKey(time)
}
/**
 * [SaveImage description: Saves an image]
 * @param {[type]} Name  string   [name to be saved]
 * @param {[type]} Image gocv.Mat [the image to be saved]
 */
func (ie ImageExtractor) SaveImage(i int,Name string) {
	gocv.IMWrite(Name, ie.Images[i]) //save the image
}

/**
 * [visit description:]
 * @param  {[type]} files *[]string        [array of files names]
 * @return {[type]} filepath.WalkFunc      [parameter used at filepath.Walk()]
 */
func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

/**
 * [FolderLength description: get the number of files in the folder]
 * @param {[type]} folder string [name of folder]
 * @return {[type]} int          [lenght of the folder(number of files)]
 */
func FolderLength(folder string) int {
	var files []string

	err := filepath.Walk(folder, visit(&files))

	if err != nil {
		panic(err)
	}
	return (len(files) - 1)
}

func (ie *ImageExtractor) getFolderName(path string,index int){
	if len((*ie).split) == 0{
		(*ie).split = make([][]string,len((*ie).readOrigins))
	}
	(*ie).split[index] = append(strings.Split(path, "/"))
	for i := 0; i < len((*ie).split[index]); i++ {
		fmt.Println((*ie).split[index][i])
	}
}

func (ie ImageExtractor) verifycandidate(candidate []string) bool{
	if candidate[0] == ".." || candidate[0] == "." { //"../src/imagehandler/Images/danger" or "./Images/grass_1.png"
		return true
	} else{
		return false
	}
}

func (ie *ImageExtractor) SetOrigins(origins []string) ([]bool,error){
	
	var originsIntegrity bool = true
	path := make([][]string,len(origins))
	statusorigins := make([]bool,len(origins))

	for i := 0; i < len(origins); i++ {
		
		path[i] = append(strings.Split(origins[i], "/"))
		statusorigins[i] = (*ie).verifycandidate(path[i])
		
		if originsIntegrity {
			originsIntegrity = statusorigins[i] 
		}
	}

	if originsIntegrity{
		(*ie).readOrigins = origins
		return statusorigins,nil
	}else{
		return statusorigins,errors.New("There was an error to set the origins, path provided is not valid")
	}
}

func (ie *ImageExtractor) Read() error{
	if len((*ie).readOrigins) == 0{
		return errors.New("Origins were not provided, use ReadFloder or define the Origins")
	} else{
		for i := 0; i < len((*ie).readOrigins); i++ {
			
			(*ie).getFolderName((*ie).readOrigins[i],i)
			(*ie).setLabelbyPath(i)
			
			if i == 0{
				(*ie).Readinfo.Labelsize[i].Size_l = (*ie).ReadFolder((*ie).readOrigins[i],true,true,false)
			} else{
				(*ie).Readinfo.Labelsize[i].Size_l = (*ie).ReadFolder((*ie).readOrigins[i],true,true,false,i*(*ie).Readinfo.Labelsize[i-1].Size_l) //temporary solution

			}
		}
		(*ie).Readinfo.SizeData = 150
		return nil
	}
}

func (ie *ImageExtractor) setLabelbyPath(index int,meaningfulname ...int){
	if len((*ie).Readinfo.Labelsize) == 0{
		(*ie).Readinfo.Labelsize = make([]cartesian.Sizelabel,len((*ie).readOrigins))
	}
	if len(meaningfulname) == 0{
		(*ie).Readinfo.Labelsize[index].Label = (*ie).split[index][len((*ie).split[index])-1]
		fmt.Println("The label that was defined: ", (*ie).Readinfo.Labelsize[index].Label)
	} else{
		fmt.Println("REMEMBER TO IMPLEMENT THAT OPTION")
		(*ie).Readinfo.Labelsize[index].Label = (*ie).split[index][len((*ie).split[index])-1]
		fmt.Println("The label that was defined: ", (*ie).Readinfo.Labelsize[index].Label)
	}
}