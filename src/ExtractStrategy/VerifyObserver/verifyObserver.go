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
	this cartesian.ReadInformation
}

type ImagesVerifier struct{
	this []gocv.Mat
}

type TextExtractorVerifier struct{
	this [][]string
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
func (this *ReadinfoVerifier) verify(){}
func (this *ImagesVerifier) verify(){}
func (this *TextExtractorVerifier) verify(){}

