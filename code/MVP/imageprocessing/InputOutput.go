package imageprocessing

import (
    "gocv.io/x/gocv"
    "path/filepath"
)

/**
 * @brief      Reads an image. colorfull or not
 *
 * @param      Image      The image
 * @param      path       The path
 * @param      show       To show the image
 * @param      save       To save the image
 * @param      colorfull  To read as an rgb image
 *
 * @return     { An image os the type gocv.Mat }
 */
func ReadImage(Image gocv.Mat, path string, show bool, save bool, colorfull bool) gocv.Mat {

    ImagePath := filepath.Join(path) //set path to the base image

    if colorfull {
        Image = gocv.IMRead(ImagePath, gocv.IMReadUnchanged) //read the base image as as RGB
    } else {
        Image = gocv.IMRead(ImagePath, gocv.IMReadGrayScale) //read the base image in grayscale
    }

    if show {
        ShowImage("And this is yout image", Image, 0)
    }

    return Image
}

/**
 * @brief      Saves an image.
 *
 * @param      Name   The name
 * @param      Image  The image
 *
 * @return     { An image os the type gocv.Ma }
 */
func SaveImage(Name string, Image gocv.Mat) {
    gocv.IMWrite(Name, Image) //save the image
}

/**
 * @brief      Shows the image.
 *
 * @param      Menssage  The menssage in the window
 * @param      Image     The image
 * @param      time      The time of the window
 *
 * @return     { An image os the type gocv.Ma }
 */
func ShowImage(Menssage string, Image gocv.Mat, time int) {
    window := gocv.NewWindow(Menssage) //basic window
    window.IMShow(Image)               //show the image
    window.WaitKey(time)
}
