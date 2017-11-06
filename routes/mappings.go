package routes

import(
	
	"strings"

)

func mappings(s string) (string, int){

	s = strings.ToLower(s)

	switch s{

		case "premier league" , "pl": return "competition_id", 2
		
		//team matches
		case "liverpool", "reds": return "team_id", 1
		case "burnley": return "team_id", 480
		case "leicester city": return "team_id", 481
		case "brighton & hove albion", "hove albion": return "team_id", 482
		case "watford": return "team_id", 483
		case "southampton": return "team_id", 69
		case "chelsea", "blues": return "team_id", 7
		case "stoke city", "stoke": return "team_id", 8
		case "west ham united", "west ham", "westham": return "team_id", 202 
		case "manchester united", "manu", "man u", "manutd", "man utd": return "team_id", 2
		case "everton", "toffees": return "team_id", 12
		case "huddersfield town", "huddersfield": return "team_id", 557
		case "bournemouth": return "team_id", 558 
		case "swansea city", "swansea": return "team_id", 15
		case "crystal palace": return "team_id", 337
		case "arsenal", "afc", "gunners", "4th": return "team_id", 18 
		case "newcastle united", "newcastle": return "team_id", 19
		case "manchester city", "man city", "citizens", "city": return "team_id", 20
		case "west bromwich albion", "west brom": return "team_id", 14
		case "tottenham hotspur", "spurs", "tottenham", "they suck": return "team_id", 13

		//team profiles
		case "liverpool profile", "reds profile": return "team_profile", 1
		case "burnley profile": return "team_profile", 480
		case "leicester city profile": return "team_profile", 481
		case "brighton & hove albion profile", "hove albion profile": return "team_profile", 482
		case "watford profile": return "team_profile", 483
		case "southampton profile": return "team_profile", 69
		case "chelsea profile", "blues profile": return "team_profile", 7
		case "stoke city profile", "stoke profile": return "team_profile", 8
		case "west ham united profile", "west ham profile", "westham profile": return "team_profile", 202 
		case "manchester united profile", "manu profile", "man u profile", "manutd profile", "man utd profile": return "team_profile", 2
		case "everton profile", "toffees profile": return "team_profile", 12
		case "huddersfield town profile", "huddersfield profile": return "team_profile", 557
		case "bournemouth profile": return "team_profile", 558 
		case "swansea city profile", "swansea profile": return "team_profile", 15
		case "crystal palace profile": return "team_profile", 337
		case "arsenal profile", "afc profile", "gunners profile", "4th profile": return "team_profile", 18 
		case "newcastle united profile", "newcastle profile": return "team_profile", 19
		case "manchester city profile", "man city profile", "citizens profile", "city profile": return "team_profile", 20
		case "west bromwich albion profile", "west brom profile": return "team_profile", 14
		case "tottenham hotspur profile", "spurs profile", "tottenham profile", "they suck profile": return "team_profile", 13


		case "live scores", "scores": return "live", 0

		default: return "",0

	}
	
}