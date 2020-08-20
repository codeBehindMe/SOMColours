/*
   SOMColours
   Copyright (C) 2020  aarontillekeratne

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

/*
  Author: aarontillekeratne
  Contact: github.com/codeBehindMe
*/

package main

import (
	"github.com/codeBehindMe/gosom"
	"github.com/codeBehindMe/gosom/feed"
	"github.com/codeBehindMe/gosom/mapx"
)

func main() {

	somTenColourHundredIterations := gosom.NewSOM64(feed.CSVFileFeeder{
		Filename:    "ten_colours.csv",
		FeatureSize: 3,
	}, 10, 10, 100, 0.1, mapx.RandomNormalInitialiser)
	somTenColourHundredIterations.Train()
	err := somTenColourHundredIterations.DumpMapToFile("somTenColourHundredIterations.json")
	if err != nil {
		panic(err)
	}

	somTenColourTwoHundredIterations := gosom.NewSOM64(feed.CSVFileFeeder{
		Filename:    "ten_colours.csv",
		FeatureSize: 3,
	}, 10, 10, 200, 0.1, mapx.RandomNormalInitialiser)
	somTenColourTwoHundredIterations.Train()
	err = somTenColourTwoHundredIterations.DumpMapToFile("somTenColourTwoHundredIterations.json")
	if err != nil {
		panic(err)
	}

	somTenColourFiveHundredIterations := gosom.NewSOM64(feed.CSVFileFeeder{
		Filename:    "ten_colours.csv",
		FeatureSize: 3,
	}, 10, 10, 500, 0.1, mapx.RandomNormalInitialiser)
	somTenColourFiveHundredIterations.Train()
	err = somTenColourFiveHundredIterations.DumpMapToFile("somTenColourFiveHundredIterations.json")
	if err != nil {
		panic(err)
	}

	somTenColourThousandIterations := gosom.NewSOM64(feed.CSVFileFeeder{
		Filename:    "ten_colours.csv",
		FeatureSize: 3,
	}, 100, 100, 1000, 0.1, mapx.RandomNormalInitialiser)
	somTenColourThousandIterations.Train()
	err = somTenColourThousandIterations.DumpMapToFile("somTenColourThousandIterations.json")
	if err != nil {
		panic(err)
	}

}
