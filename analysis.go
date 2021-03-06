package main

import (
	"github.com/EndlessCheng/mahjong-helper/util"
	"fmt"
	"strings"
	"github.com/fatih/color"
)

func _printIncShantenResults14(shanten int, incShantenResults14 util.WaitsWithImproves14List) {
	if len(incShantenResults14) == 0 {
		return
	}

	if len(incShantenResults14[0].OpenTiles) > 0 {
		fmt.Print("鸣牌后")
	}
	fmt.Println(util.NumberToChineseShanten(shanten+1) + "：")
	for _, result := range incShantenResults14 {
		printWaitsWithImproves13_oneRow(result.Result13, result.DiscardTile, result.OpenTiles)
	}
}

func analysisTiles34(roundWindTile34 int, selfWindTile34 int, tiles34 []int, leftTiles34 []int, isOpen bool) error {
	humanTiles := util.Tiles34ToStr(tiles34)
	fmt.Println(humanTiles)
	fmt.Println(strings.Repeat("=", len(humanTiles)))

	countOfTiles := util.CountOfTiles34(tiles34)
	switch countOfTiles % 3 {
	case 1:
		result := util.CalculateShantenWithImproves13(roundWindTile34, selfWindTile34, tiles34, leftTiles34, isOpen)
		fmt.Println(util.NumberToChineseShanten(result.Shanten) + "：")
		printWaitsWithImproves13_oneRow(result, -1, nil)
	case 2:
		shanten, results14, incShantenResults14 := util.CalculateShantenWithImproves14(roundWindTile34, selfWindTile34, tiles34, leftTiles34, isOpen)

		if shanten == -1 {
			color.HiRed("【已胡牌】")
			break
		}

		if shanten == 0 {
			color.HiRed("【已听牌】")
		}

		fmt.Println(util.NumberToChineseShanten(shanten) + "：")
		for _, result := range results14 {
			printWaitsWithImproves13_oneRow(result.Result13, result.DiscardTile, result.OpenTiles)
		}
		_printIncShantenResults14(shanten, incShantenResults14)
	default:
		return fmt.Errorf("参数错误: %d 张牌", countOfTiles)
	}

	fmt.Println()

	return nil
}

func analysisMeld(roundWindTile34 int, selfWindTile34 int, tiles34 []int, leftTiles34 []int, targetTile34 int, allowChi bool) {
	// 原始手牌分析
	isOpen := util.CountOfTiles34(tiles34) < 13
	result := util.CalculateShantenWithImproves13(roundWindTile34, selfWindTile34, tiles34, leftTiles34, isOpen)

	// 副露分析
	shanten, results14, incShantenResults14 := util.CalculateMeld(roundWindTile34, selfWindTile34, tiles34, targetTile34, allowChi, leftTiles34)

	if len(results14) == 0 && len(incShantenResults14) == 0 {
		return
	}

	raw := util.Tiles34ToStr(tiles34) + " + " + util.Tile34ToStr(targetTile34) + "?"
	fmt.Println(raw)
	fmt.Println(strings.Repeat("=", len(raw)))

	fmt.Println("当前" + util.NumberToChineseShanten(result.Shanten) + "：")
	printWaitsWithImproves13_oneRow(result, -1, nil)

	if shanten == -1 {
		color.HiRed("【已胡牌】")
		return
	}

	// 打印结果
	// FIXME: 选择很多时如何精简何切选项？
	const maxShown = 10

	if len(results14) > 0 {
		fmt.Println("鸣牌后" + util.NumberToChineseShanten(shanten) + "：")
		shownResults14 := results14
		if len(shownResults14) > maxShown {
			shownResults14 = shownResults14[:maxShown]
		}
		for _, result := range shownResults14 {
			printWaitsWithImproves13_oneRow(result.Result13, result.DiscardTile, result.OpenTiles)
		}
	}

	shownIncResults14 := incShantenResults14
	if len(shownIncResults14) > maxShown {
		shownIncResults14 = shownIncResults14[:maxShown]
	}
	_printIncShantenResults14(shanten, shownIncResults14)
}

func analysisHumanTiles(humanTiles string) (tiles34 []int, err error) {
	splits := strings.Split(humanTiles, "+")
	if len(splits) == 2 {
		tiles34, err = util.StrToTiles34(splits[0])
		if err != nil {
			return
		}

		var targetTile34 int
		targetTile34, err = util.StrToTile34(splits[1])
		if err != nil {
			return
		}

		analysisMeld(27, 27, tiles34, nil, targetTile34, true)
		return
	}

	tiles34, err = util.StrToTiles34(humanTiles)
	if err != nil {
		return
	}

	err = analysisTiles34(27, 27, tiles34, nil, false)
	return
}
