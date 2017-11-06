package structs

import(

	"strconv"
	"bytes"

)

type TeamProfile struct {
    
    Name string
    ShortCode  string
    Players []Player
    DefaultHomeVenue Stadium

}


type Player struct {

    Name string
    Number  int
    Position string

}

type Stadium struct {

    Name string
    Capacity  int

}

func stringfyPlayer(player Player) string{
	
	var plr bytes.Buffer

	plr.WriteString("	")
	plr.WriteString(strconv.Itoa(player.Number))
	plr.WriteString(". ")
	plr.WriteString(player.Name)

	plr.WriteString(" (")
	plr.WriteString(player.Position)
	plr.WriteString(")")

	return plr.String()

}

func stringfyPlayers(players []Player) string{
	
	var plrs bytes.Buffer

	for i := 0; i < len(players); i++ {
		
		plrs.WriteString(stringfyPlayer(players[i]))
		
		if(i != len(players) - 1){

			plrs.WriteString("\n")
		
		}

	}

	return plrs.String()

}


func stringfyStadium(stadium Stadium) string{
	
	var std bytes.Buffer

	std.WriteString("	Name: ")
	std.WriteString(stadium.Name)
	std.WriteString("\n")
	std.WriteString("	Capacity: ")
	std.WriteString(strconv.Itoa(stadium.Capacity))

	return std.String()

}

func StringfyTeamProfile(team TeamProfile) string{
	
	var profile bytes.Buffer

	profile.WriteString("Name: ")
	profile.WriteString(team.Name)
	profile.WriteString(" (")
	profile.WriteString(team.ShortCode)
	profile.WriteString(")")
	profile.WriteString("\n")
	profile.WriteString("Home Turf:\n")
	profile.WriteString(stringfyStadium(team.DefaultHomeVenue))
	profile.WriteString("\n")
	profile.WriteString("Players:\n")
	profile.WriteString(stringfyPlayers(team.Players))
	

	return profile.String()

}

