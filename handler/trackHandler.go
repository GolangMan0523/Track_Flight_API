package handler

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func TrackHandler(c *fiber.Ctx) error {
	fmt.Println("Sort flights!", c.Query("path"))
	// unsortedPath := []byte(`
	// 	[
	// 		["GSO", "IND"],
	// 		["ATL", "GSO"],
	// 		["SFO", "ATL"],
	// 		["IND", "ASA"],
	// 		["KUA", "MAN"],
	// 		["ASA", "KUA"]
	// 	]
	// `)
	unsortedPath := []byte(c.Query("path"))

	var unsortedPathSlice [][]string
	json.Unmarshal(unsortedPath, &unsortedPathSlice)

	fmt.Println("=> Before order=>", unsortedPathSlice)

	if len(unsortedPathSlice) == 0 {
		return c.Status(500).SendString("Error occurs!")
	}

	finalPath, err := getFinalPath(unsortedPathSlice)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	jsonFinalPath, err := json.Marshal(finalPath)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	fmt.Println("=> Result =>", string(jsonFinalPath))
	return c.SendString(string(jsonFinalPath))
}

func getFinalPath(paths [][]string) ([]string, error) {
	var src, des []string

	for _, path := range paths {
		src = append(src, path[0])
		des = append(des, path[1])
	}

	firstSrc := findFirstSrc(src, des)
	fmt.Println("=> firstSRC=>", firstSrc)

	var sorted [][]string
	sortedPath, err := getOrderPath(paths, sorted, firstSrc)

	if err != nil {
		return nil, err
	}

	var finalPath []string
	finalPath = append(finalPath, sortedPath[0][0], sortedPath[len(sortedPath)-1][1])
	return finalPath, nil
}

func findFirstSrc(src []string, des []string) string {
	var firstSrc string
	for _, srcItem := range src {
		var count int
		for _, desItem := range des {
			if srcItem == desItem {
				count++
			}
		}
		if count == 0 {
			firstSrc = srcItem
		}
	}
	return firstSrc
}

func getOrderPath(paths [][]string, sorted [][]string, src string) ([][]string, error) {
	var count int
	var err error
	if len(sorted) == len(paths) {
		fmt.Println("=> Finish sort =>", sorted)
		return sorted, nil
	} else {
		for _, path := range paths {
			if path[0] == src {
				count++
				sorted = append(sorted, path)
			}
		}
		if count == 0 {
			err = errors.New("Error occurs")
			return nil, err
		} else {
			return getOrderPath(paths, sorted, sorted[len(sorted)-1][1])
		}
	}
}
