package readerstrategy

import (
	"../basicdata"
	"path/filepath"
	"fmt"
	"strings"
	"errors"
)

func (dr *DataReader) SetReadStrategy(rs readStrategy) {
	dr.Strategy = rs
}
func (dr *DataReader) ProcessRead(){
	dr.Strategy.ReadData(dr)
}
func (dr *DataReader) ProcessAllocation(size int){
	(*dr).Readinfo.SizeData = size //temporary solution
	dr.Strategy.Allocate(dr)
}


/**
 * [ReadFolder description: read all images at some folder]
 * @param {[type]} Images    *[]gocv.Mat [An Array of gocv.Mat that will be used to contain the images of the folder]
 * @param {[type]} folder    string      [folder name]
 * @param {[type]} print     bool        [if its true, print the names]
 * @param {[type]} show      bool        [if its true, show the images]
 * @param {[type]} colorfull bool        [if its is true take a 3 chanel rbg image]
 */
//func (ie *ImageExtractor) ReadFolder(folder string, print bool, show bool, colorfull bool,index ...int) int{
func (dr *DataReader) ReadFolder(index ...int) int{
	
	var files []string
	var first bool = true
	var i int

	err := filepath.Walk((*dr).path, visit(&files))
	if err != nil {
		panic(err)
	} else {
		if (*dr).Readinfo.SizeData == 0{
			(*dr).ProcessAllocation(3*(len(files)-1)) 	//temporary solution
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
		if (*dr).print {

			fmt.Println("geting image:     ", file)

		}
		(*dr).ProcessRead()
		i++
	}
	return len(files)-1
}
//func (dr *DataReader) getFolderName(path string,index int){
func (dr *DataReader) getFolderName(index int){
	if len((*dr).split) == 0{
		(*dr).split = make([][]string,len((*dr).readOrigins))
	}
	(*dr).split[index] = append(strings.Split((*dr).path, "/"))
	for i := 0; i < len((*dr).split[index]); i++ {
		fmt.Println((*dr).split[index][i])
	}
}

func (dr *DataReader) SetOrigins(origins []string) ([]bool,error){
	
	var originsIntegrity bool = true
	path := make([][]string,len(origins))
	statusorigins := make([]bool,len(origins))

	for i := 0; i < len(origins); i++ {
		
		path[i] = append(strings.Split(origins[i], "/"))
		statusorigins[i] = (*dr).verifycandidate(path[i])
		
		if originsIntegrity {
			originsIntegrity = statusorigins[i] 
		}
	}

	if originsIntegrity{
		(*dr).readOrigins = origins
		return statusorigins,nil
	}else{
		return statusorigins,errors.New("There was an error to set the origins, path provided is not valid")
	}
}

func (dr *DataReader) Read() error{
	if len((*dr).readOrigins) == 0{
		return errors.New("Origins were not provided, use ReadFloder or define the Origins")
	} else{
		for i := 0; i < len((*dr).readOrigins); i++ {
			
			(*dr).getFolderName(i)
			(*dr).setLabelbyPath(i)
			
			if i == 0{
				(*dr).Readinfo.Labelsize[i].Size_l = (*dr).ReadFolder()
			} else{
				(*dr).Readinfo.Labelsize[i].Size_l = (*dr).ReadFolder(i*(*dr).Readinfo.Labelsize[i-1].Size_l) //temporary solution

			}
		}
		(*dr).Readinfo.SizeData = 150
		return nil
	}
}

func (dr *DataReader) setLabelbyPath(index int,meaningfulname ...int){
	if len((*dr).Readinfo.Labelsize) == 0{
		(*dr).Readinfo.Labelsize = make([]cartesian.Sizelabel,len((*dr).readOrigins))
	}
	if len(meaningfulname) == 0{
		(*dr).Readinfo.Labelsize[index].Label = (*dr).split[index][len((*dr).split[index])-1]
		fmt.Println("The label that was defined: ", (*dr).Readinfo.Labelsize[index].Label)
	} else{
		fmt.Println("REMEMBER TO IMPLEMENT THAT OPTION")
		(*dr).Readinfo.Labelsize[index].Label = (*dr).split[index][len((*dr).split[index])-1]
		fmt.Println("The label that was defined: ", (*dr).Readinfo.Labelsize[index].Label)
	}
}

func (dr DataReader) verifycandidate(candidate []string) bool{
	if candidate[0] == ".." || candidate[0] == "." { //"../src/imagehandler/Images/danger" or "./Images/grass_1.png"
		return true
	} else{
		return false
	}
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
