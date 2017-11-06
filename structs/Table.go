package structs

import(

	"strconv"
	"bytes"

)

type Table struct {

    LeagueTable []TeamStanding
    Subround string

}

type TeamStanding struct {

    Name string
	Points int
	Wins int 
	Losses int 
	Draws int
	GamesPlayed int
	GoalsFor int
	GoalsAgainst int
	GoalsDiff int

}

func stringfyTeamStanding(tableRow TeamStanding, position int) string{

	var teamStanding bytes.Buffer

	teamStanding.WriteString(strconv.Itoa(position))
	teamStanding.WriteString(". ")
	teamStanding.WriteString(tableRow.Name)
	teamStanding.WriteString(" Points: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.Points))
	teamStanding.WriteString(", W: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.Wins)) 
	teamStanding.WriteString(", L: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.Losses))
	teamStanding.WriteString(", D: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.Draws))
	teamStanding.WriteString(", GP: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.GamesPlayed))
	teamStanding.WriteString(", GF: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.GoalsFor))
	teamStanding.WriteString(", GA: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.GoalsAgainst))
	teamStanding.WriteString(", GD: ")
	teamStanding.WriteString(strconv.Itoa(tableRow.GoalsDiff))

	return teamStanding.String()




	
}

func StringfyTable(table Table) string{

	var leagueTable bytes.Buffer

	for i := 0; i < len(table.LeagueTable); i++ {
		
		row := table.LeagueTable[i]
		leagueTable.WriteString(stringfyTeamStanding(row, i+1))
		leagueTable.WriteString("\n")

	}

	return leagueTable.String()
	
}