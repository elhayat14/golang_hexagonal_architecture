package query_builder

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"time"
)

func BuildAndDirectFilter(filter map[string]interface{}, excludeZero map[string]bool) bson.D {
	var result bson.D
	var filterData bson.D
	for key, value := range filter {
		val := reflect.ValueOf(value)
		if !val.IsZero() {
			filterData = append(filterData, bson.E{Key: key, Value: value})
			continue
		}
		if excludeZero != nil {
			if _, ok := excludeZero[key]; ok && val.IsZero() {
				filterData = append(filterData, bson.E{Key: key, Value: value})
			}
		}
	}
	if filterData != nil {
		result = bson.D{{Key: "$match", Value: filterData}}
	}
	return result
}
func BuildDateRangeFilter(filter map[string]map[string]time.Time) bson.D {
	var result bson.D
	var filterData bson.D
	for key, value := range filter {
		startVal := value["start"]
		endVal := value["end"]
		if !startVal.IsZero() && !endVal.IsZero() {
			element := bson.E{
				Key: key,
				Value: bson.D{
					{Key: "$gte", Value: startVal},
					{Key: "$lt", Value: endVal},
				},
			}
			filterData = append(filterData, element)
		}
	}
	if filterData != nil {
		result = bson.D{{Key: "$match", Value: filterData}}
	}
	return result
}
func Skip(skip int) bson.D {
	return bson.D{
		{Key: "$skip", Value: skip},
	}
}
func Limit(limit int) bson.D {
	return bson.D{
		{Key: "$limit", Value: limit},
	}
}
func Count() bson.D {
	return bson.D{
		{Key: "$count", Value: "totalData"},
	}
}
func Sort(data map[string]int) bson.D {
	var sortedBy bson.D
	for index, value := range data {
		sortedBy = append(sortedBy, bson.E{Key: index, Value: value})
	}
	result := bson.D{
		{Key: "$sort", Value: sortedBy},
	}
	return result
}
func Unwind(field string, preserveNullAndEmptyArrays bool) bson.D {
	return bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: field},
			{Key: "preserveNullAndEmptyArrays", Value: preserveNullAndEmptyArrays},
		}},
	}
}
func BuildSortDataQuery(data map[string]string) map[string]int {
	var result map[string]int
	result = make(map[string]int)
	for index, value := range data {
		intValue := 1
		switch value {
		case "desc":
			intValue = -1
		case "asc":
			intValue = 1
		}
		result[index] = intValue
	}
	return result
}
