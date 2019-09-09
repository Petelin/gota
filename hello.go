package main

import (
	"fmt"

	"gota/dataframe"
	"gota/series"

	"rsc.io/quote"
)

func Hello() string {
	return quote.Hello()
}

type DataF struct {
	dataframe.DataFrame
	index []string
}

func (df *DataF) AddElement(objectID, fieldID string, val interface{}) {
	//s := df.Col(fieldID)
	//if s.Err != nil {
	//	df.Set(0, dataframe.New(series.New()))
	//}
	//df.Set()
}

func main() {
	df := dataframe.New(
		series.New([]string{"b", "a"}, series.String, "_id"),
		series.New([]int{1, 2}, series.Int, "age"),
		series.New([]float64{3.0, 4.0}, series.Float, "score"),
	)
	fmt.Println(df.Select("score"))
	df.Capply(func(s series.Series) series.Series {
		min := s.Min()
		max := s.Max()
		for i := 0; i < s.Len(); i++ {
			e := s.Elem(i)
			normaled := (e.Float() - min) / (max - min)
			e.Set(normaled)
		}
		return s
	})
	fmt.Println(df.Select("score"))

	df.Rapply(func(i series.Series) series.Series {
		return i
	})

	df = dataframe.LoadMaps(
		[]map[string]interface{}{
			map[string]interface{}{
				"B": 1,
				"C": true,
				"D": 0,
			},
			map[string]interface{}{
				"A": "",
				"B": 2,
				"C": true,
				"D": 0.5,
			},
		},
	)
	fmt.Println(df)
}
