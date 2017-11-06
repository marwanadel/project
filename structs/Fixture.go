package structs

import(

	"strconv"
	"bytes"

)

type Fixture struct {

    HomeTeam Team
    AwayTeam Team
    AwayGoals  int
    HomeGoals  int
    Start float64
    CurrentState int

}

type Team struct {

    Name string

}

func StringfyTeamFixtures(fix []Fixture) string{

	var message bytes.Buffer 

	var fixture bytes.Buffer

	if(len(fix) == 0){

		return "No Recent Matches!"

	}

	for i := 0; i < len(fix); i++ {
		
		homeTeam := fix[i].HomeTeam
		awayTeam := fix[i].AwayTeam

		if(fix[i].CurrentState == 0){

			fixture.WriteString(homeTeam.Name)
			fixture.WriteString(" - ")
			fixture.WriteString(awayTeam.Name)
			fixture.WriteString("\n")
				
		}else{

			if(fix[i].CurrentState == 1){

				fixture.WriteString(homeTeam.Name)
				fixture.WriteString(" ")
				fixture.WriteString(strconv.Itoa(fix[i].HomeGoals))
				fixture.WriteString(" - ")
				fixture.WriteString(strconv.Itoa(fix[i].AwayGoals))
				fixture.WriteString(" ")
				fixture.WriteString(awayTeam.Name)
				fixture.WriteString("	1st. Half")
				fixture.WriteString("\n")

			}else{

				if(fix[i].CurrentState == 2){
				
					fixture.WriteString(homeTeam.Name)
					fixture.WriteString(" ")
					fixture.WriteString(strconv.Itoa(fix[i].HomeGoals))
					fixture.WriteString(" - ")
					fixture.WriteString(strconv.Itoa(fix[i].AwayGoals))
					fixture.WriteString(" ")
					fixture.WriteString(awayTeam.Name)
					fixture.WriteString("	HT")
					fixture.WriteString("\n")

				}else{

					if(fix[i].CurrentState == 3){

						fixture.WriteString(homeTeam.Name)
						fixture.WriteString(" ")
						fixture.WriteString(strconv.Itoa(fix[i].HomeGoals))
						fixture.WriteString(" - ")
						fixture.WriteString(strconv.Itoa(fix[i].AwayGoals))
						fixture.WriteString(" ")
						fixture.WriteString(awayTeam.Name)
						fixture.WriteString("	2nd. Half")
						fixture.WriteString("\n")

					}else{

						fixture.WriteString(homeTeam.Name)
						fixture.WriteString(" ")
						fixture.WriteString(strconv.Itoa(fix[i].HomeGoals))
						fixture.WriteString(" - ")
						fixture.WriteString(strconv.Itoa(fix[i].AwayGoals))
						fixture.WriteString(" ")
						fixture.WriteString(awayTeam.Name)
						fixture.WriteString("	FT")
						fixture.WriteString("\n")

					}

				}

			}

		}	

		message.WriteString(fixture.String()) 

	}

	return message.String()
	
}

func StringfyLiveFixtures(fix []Fixture) string{

	var message bytes.Buffer 

	var fixture bytes.Buffer

	if(len(fix) == 0){

		return "No Live Matches!"

	}

	var live int = 0;

	for i := 0; i < len(fix); i++ {
		
		homeTeam := fix[i].HomeTeam
		awayTeam := fix[i].AwayTeam

		if(fix[i].CurrentState == 0){


				
		}else{

			if(fix[i].CurrentState == 1){

				fixture.WriteString(homeTeam.Name)
				fixture.WriteString(" ")
				fixture.WriteString(strconv.Itoa(fix[i].HomeGoals))
				fixture.WriteString(" - ")
				fixture.WriteString(strconv.Itoa(fix[i].AwayGoals))
				fixture.WriteString(" ")
				fixture.WriteString(awayTeam.Name)
				fixture.WriteString("	1st. Half")
				fixture.WriteString("\n")

				live++

			}else{

				if(fix[i].CurrentState == 2){
				
					fixture.WriteString(homeTeam.Name)
					fixture.WriteString(" ")
					fixture.WriteString(strconv.Itoa(fix[i].HomeGoals))
					fixture.WriteString(" - ")
					fixture.WriteString(strconv.Itoa(fix[i].AwayGoals))
					fixture.WriteString(" ")
					fixture.WriteString(awayTeam.Name)
					fixture.WriteString("	HT")
					fixture.WriteString("\n")

					live++

				}else{

					if(fix[i].CurrentState == 3){

						fixture.WriteString(homeTeam.Name)
						fixture.WriteString(" ")
						fixture.WriteString(strconv.Itoa(fix[i].HomeGoals))
						fixture.WriteString(" - ")
						fixture.WriteString(strconv.Itoa(fix[i].AwayGoals))
						fixture.WriteString(" ")
						fixture.WriteString(awayTeam.Name)
						fixture.WriteString("	2nd. Half")
						fixture.WriteString("\n")

						live++

					}else{



					}

				}

			}

		}

		message.WriteString(fixture.String()) 

	}

	if(live == 0){

		return "No Live Matches!"		

		}

	return message.String()
	
}