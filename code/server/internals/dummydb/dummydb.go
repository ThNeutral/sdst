package dummydb

type WriteDBRequest struct {
	Table string
	Row   map[string]string
}

type ReadDBRequest struct {
	Table  string
	Fields []string
	Data   chan []map[string]string
}

type DummyDB struct {
	Data         map[string][]map[string]string
	WriteRequest chan WriteDBRequest
	ReadRequest  chan ReadDBRequest
}

func GetDummyDB() *DummyDB {
	return &DummyDB{
		Data:         make(map[string][]map[string]string),
		WriteRequest: make(chan WriteDBRequest),
		ReadRequest:  make(chan ReadDBRequest),
	}
}

func (dummydb *DummyDB) Run() {
	for {
		select {
		case request := <-dummydb.WriteRequest:
			if _, exists := dummydb.Data[request.Table]; !exists {
				dummydb.Data[request.Table] = []map[string]string{}
			}
			dummydb.Data[request.Table] = append(dummydb.Data[request.Table], request.Row)

		case request := <-dummydb.ReadRequest:
			result := []map[string]string{}
			if table, tableExists := dummydb.Data[request.Table]; tableExists {
				for _, row := range table {
					rowResult := make(map[string]string)
					for _, field := range request.Fields {
						if value, fieldExists := row[field]; fieldExists {
							rowResult[field] = value
						} else {
							rowResult[field] = ""
						}
					}
					result = append(result, rowResult)
				}
			}
			request.Data <- result
		}
	}
}
