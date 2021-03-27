package nonparametric

import (
	"../genericdata"
	"errors"
	"fmt"
	"math"
	"sort"
	"../basicdata"

)

type Knn struct {
	interestgroup 		[]cartesian.Interest
}