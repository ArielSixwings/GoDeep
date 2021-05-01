package process

func (dp *DataProcessing) SetProcessStrategy(ps processStrategy) {
	dp.Strategy = ps
}

func (dp *DataProcessing) GroupProcess(){
	for i := 0; i < (*dp).Readinfo.SizeData; i++ {
		(*dp).Strategy.Process()
	}
}