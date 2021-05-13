package cartesian

import(
	"image/color"
)

func red(thecolor *color.RGBA){
	(*thecolor).R = 255
	(*thecolor).G = 0
	(*thecolor).B = 0
	(*thecolor).A = 1
}

func blue(thecolor *color.RGBA){
	(*thecolor).R = 0
	(*thecolor).G = 0
	(*thecolor).B = 255
	(*thecolor).A = 1
}

func purple(thecolor *color.RGBA){
	(*thecolor).R = 107
	(*thecolor).G = 66
	(*thecolor).B = 143
	(*thecolor).A = 1
}

func copper(thecolor *color.RGBA){
	(*thecolor).R = 184
	(*thecolor).G = 115
	(*thecolor).B = 51
	(*thecolor).A = 1
}

func silver(thecolor *color.RGBA){
	(*thecolor).R = 170
	(*thecolor).G = 169
	(*thecolor).B = 173
	(*thecolor).A = 1
}

func gold(thecolor *color.RGBA){
	(*thecolor).R = 255
	(*thecolor).G = 215
	(*thecolor).B = 0
	(*thecolor).A = 1
}

func pink(thecolor *color.RGBA){
	(*thecolor).R = 255
	(*thecolor).G = 20
	(*thecolor).B = 147
	(*thecolor).A = 1
}