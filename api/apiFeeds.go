package api

import (

	"net/http"
	"encoding/json"
	"golazo/structs"
	"io/ioutil"
	"strconv"
	"time"

)

func GetTeamProfile(w http.ResponseWriter, r *http.Request, b []byte, code int) {

	req, er := http.NewRequest("GET", "https://api.crowdscores.com/v1/teams/" + strconv.Itoa(code), nil)


	if er != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	q := req.URL.Query()
	
	//change to request header instead of filter
	q.Add("api_key", "3edf983a707b45fc9db98cac6dc36d9d")

	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())

	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}	

	b, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	
	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	// Unmarshal
	var profile structs.TeamProfile
	err = json.Unmarshal(b, &profile)
	
	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	msg := structs.Message{}

	msg.Message = structs.StringfyTeamProfile(profile)

	output, err := json.Marshal(msg)
	
	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
	
}

func GetTeamMatches(w http.ResponseWriter, r *http.Request, b []byte, code int) {
	
	req, er := http.NewRequest("GET", "https://api.crowdscores.com/v1/matches", nil)

	if er != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	q := req.URL.Query()
	
	//change to request header instead of filter
	q.Add("api_key", "3edf983a707b45fc9db98cac6dc36d9d")

	q.Add("team_id", strconv.Itoa(code))
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())

	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	b, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	
	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	// Unmarshal
	var fix []structs.Fixture
	err = json.Unmarshal(b, &fix)
	
	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	msg := structs.Message{}

	msg.Message = structs.StringfyTeamFixtures(fix)

	output, err := json.Marshal(msg)
	
	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)

}

func GetLeagueTable(w http.ResponseWriter, r *http.Request, b []byte, code int) {
	
	req, er := http.NewRequest("GET", "https://api.crowdscores.com/v1/league-tables", nil)

	if er != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	q := req.URL.Query()

	//change to request header instead of filter
	q.Add("api_key", "3edf983a707b45fc9db98cac6dc36d9d")

	q.Add("competition_id", strconv.Itoa(code))
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())

	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	b, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	
	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	// Unmarshal
	
	var table []structs.Table 
	err = json.Unmarshal(b, &table)
	
	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	msg := structs.Message{}

	if(len(table) == 0){

		msg.Message = "No table currently available!"

	}else{

		msg.Message = structs.StringfyTable(table[0])

	}

	output, err := json.Marshal(msg)
	
	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}	

	w.Header().Set("content-type", "application/json")
	w.Write(output)

}

func GetLiveScores(w http.ResponseWriter, r *http.Request, b []byte, code int) {

    req, er := http.NewRequest("GET", "https://api.crowdscores.com/v1/matches", nil)

    if er != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	q := req.URL.Query()

	//change to request header instead of filter
	q.Add("api_key", "3edf983a707b45fc9db98cac6dc36d9d")

	q.Add("competition_id", strconv.Itoa(code))

	timein := time.Now().Local().Add(time.Hour * time.Duration(-3) +
                                 time.Minute * time.Duration(0) +
                                 time.Second * time.Duration(0))



	q.Add("from", timein.Format("2006-01-02T15:04:05-07:00"))
								 
	q.Add("to", time.Now().Format("2006-01-02T15:04:05-07:00"))
	
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())

	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	b, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	
	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	// Unmarshal
	var fix []structs.Fixture
	err = json.Unmarshal(b, &fix)

	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	msg := structs.Message{}

	msg.Message = structs.StringfyLiveFixtures(fix)

	output, err := json.Marshal(msg)
	
	if err != nil {
	
		msg := structs.Message{}

		msg.Message = "Oops! An error has occured. Please, try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return
	
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)

}