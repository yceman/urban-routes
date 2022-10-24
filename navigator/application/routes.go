package routes

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Routes struct {
	ID        string
	ClientID  string
	Positions []Position
}

type Position struct {
	Lat  float64
	Long float64
}

type PartialRoutePosition struct {
	ID       string
	ClientID string
	Position []float64
	Finished bool
}

func (r *Routes) LoadPositions() error {
	if r.ID == "" {
		return errors.New("route id not informed")
	}
	f, err := os.Open("destination/" + r.ID + ".txt")
	if err != nil {
		return nil
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return nil
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}
		r.Positions = append(r.Positions, Position{
			Lat:  lat,
			Long: long,
		})
	}
	return nil
}

func (r *Routes) ExportJsonPositions() ([]string, error) {
	var routes PartialRoutePosition
	var result []string
	total := len(r.Positions)
	for k, v := range r.Positions {
		routes.ID = r.ID
		routes.ClientID = r.ClientID
		routes.Position = []float64{v.Lat, v.Long}
		routes.Finished = false

		if total-1 == k {
			routes.Finished = true
		}
		jsonRoute, err := json.Marshal(routes)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))
	}
	return result, nil
}
