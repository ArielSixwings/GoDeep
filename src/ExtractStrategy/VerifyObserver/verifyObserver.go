package main

import (
	"fmt" 
	"errors"
	"../../basicdata"
)
// func (v *Verifier) getID() string {
// 	return c.id
// }
 
type ReaderParametersVerifier struct{
	readerp readerParameters
}

type SplitVerifier struct{
	this [][]string
}

type ReadOriginsVerifier struct{
	origins []string
	candidate []string
}

type ReadinfoVerifier struct{
	Readinfo cartesian.ReadInformation
}

type ImagesVerifier struct{
	ReadinfoVerifier
	images []gocv.Mat
}

type TextExtractorVerifier struct{
	ReadinfoVerifier
	Texts [][]string
}


func (this *ReaderParametersVerifier) verify(){}
func (this *SplitVerifier) verify(){}

func (this *ReadOriginsVerifier) verify()errors{
	if len((*this).origins) == 0{
		if len((*this).candidate) == 0 {
			return errors.New("no candidate do origins were provided")
		} else{
			return nil
		}
	} else{
		if len((*this).candidate) == 0 {
			return errors.New("no candidate do origins were provided")
		} else{
			if (*this).candidate[0] == ".." || (*this).candidate[0] == "." {
				return nil
			} else{
				return errors.New("origins candidate dont have a valid format")
			}
		}		
	}
}

func (this *ImagesVerifier) verify(){
	if len((*this).images) == (*this).Readinfo.SizeData  {
		return nil 
	} else{
		return errors.New("the data weren't properly read or the SizeData weren't properly defined")
	}
}

func (this *TextExtractorVerifier) verify(){
	if len((*this).Texts) == (*this).Readinfo.SizeData  {
		return nil 
	} else{
		return errors.New("the data weren't properly read or the SizeData weren't properly defined")
	}	
}

