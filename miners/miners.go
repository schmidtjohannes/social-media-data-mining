package miners


type MinerInterface interface{
        QueryGroups (MinerData, error)
}

type MinerData struct {
	Data []*interface{}
}
