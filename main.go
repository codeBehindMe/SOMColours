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
	"github.com/codeBehindMe/gosom/algo"
	"github.com/codeBehindMe/gosom/feed"
	"github.com/codeBehindMe/gosom/mapx"
)

func main(){
	somTenColourHundredIterations := algo.NewSOM(feed.CSVFileFeeder{
		Filename:    "ten_colours.csv",
		FeatureSize: 3,
	},10,10,3,mapx.PseudoZerosOnesInitialiser,100,0.1)
	somTenColourHundredIterations.Train()
	somTenColourHundredIterations.DumpWeightsToFile("somTenColourHundredIterations.json")

	somTenColourTwoHundredIterations := algo.NewSOM(feed.CSVFileFeeder{
		Filename:    "ten_colours.csv",
		FeatureSize: 3,
	},10,10,3,mapx.PseudoZerosOnesInitialiser,200,0.1)
	somTenColourTwoHundredIterations.Train()
	somTenColourTwoHundredIterations.DumpWeightsToFile("somTenColourTwoHundredIterations.json")

	somTenColourFiveHundredIterations := algo.NewSOM(feed.CSVFileFeeder{
		Filename:    "ten_colours.csv",
		FeatureSize: 3,
	},10,10,3,mapx.PseudoZerosOnesInitialiser,500,0.1)
	somTenColourFiveHundredIterations.Train()
	somTenColourFiveHundredIterations.DumpWeightsToFile("somTenColourFiveHundredIterations.json")

	somTenColourThousandIterations := algo.NewSOM(feed.CSVFileFeeder{
		Filename:    "ten_colours.csv",
		FeatureSize: 3,
	},100,100,3,mapx.PseudoZerosOnesInitialiser,1000,0.1)
	somTenColourThousandIterations.Train()
	somTenColourThousandIterations.DumpWeightsToFile("somTenColourThousandIterations.json")
}
