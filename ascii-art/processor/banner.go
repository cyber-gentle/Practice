package processor

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func LoadBanner(bannerName string) (map[rune][]string, error) {
	var (
		validBanner = map[string]bool{
			"standard.txt":   true,
			"shadow.txt":     true,
			"thinkertoy.txt": true,
		}

		bannerMap  = make(map[rune][]string)
		charHeight []string
	)

	if !strings.HasSuffix(bannerName, ".txt") {
		bannerName += ".txt"
	}

	_, valid := validBanner[bannerName]
	if !valid {
		return nil, fmt.Errorf(
			"invalid banner: %s\nValid banners are : standard, shadow and thinkertoy",
			strings.TrimSuffix(strings.TrimPrefix(bannerName, "banners/"), ".txt"))
	} else {
		bannerName = path.Join("banners/", strings.TrimPrefix(bannerName, "banners/"))
		banners, err := os.ReadFile(bannerName)
		if err != nil {
			return nil, err
		}
		banners = []byte(strings.ReplaceAll(string(banners), "\r\n", "\n"))
		bannerLines := strings.Split(string(banners), "\n")

		for i := rune(32); i <= rune(126); i++ {
			index := (i - 32) * 9
			charHeight = bannerLines[index+1 : index+9]
			bannerMap[i] = charHeight

			for row := range 8 {
				fmt.Println(charHeight[row])
			}
		}

		for key, charHeight := range bannerMap {
			if len(charHeight) != 8 {
				return nil, fmt.Errorf("invalid character at index %d", key)
			}
		}
	}
	return bannerMap, nil
}
