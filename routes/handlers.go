package routes

import (

	"net/http"
	"encoding/json"
	"structs"
	"api"
	"io/ioutil"
	"strings"
	"crypto/md5"
	"encoding/hex"
	"time"
	"strconv"
	"fmt"

)

var sessions []string


func home(w http.ResponseWriter, r *http.Request) {

	//redicrect to /welcome
  
	http.Redirect(w, r, "/welcome", 200)

}

func welcome(w http.ResponseWriter, r *http.Request) {

	//reply with uuid to be used in subsequent requests and responses

	// Generate a UUID.
	hasher := md5.New()
	hasher.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10)))
	uuid := hex.EncodeToString(hasher.Sum(nil))

	// Create a session for this UUID
	sessions = append(sessions, uuid)

    m := structs.WelcomeMessage{Message: "Welcome", Uuid: uuid}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(m)

}

func findSession(id string) bool{
	
	for i := 0; i < len(sessions); i++ {
			
		if(strings.Compare(id, sessions[i]) == 0){

			return true

		}

	}

	return false

}

func chat(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
	
		http.Error(w, "Only POST requests are allowed.", http.StatusMethodNotAllowed)
		return
	
	}

	uuid := r.Header.Get("Authorization")
	
	if uuid == "" {
	
		http.Error(w, "Missing or empty Authorization header.", http.StatusUnauthorized)
		return
	
	}

	// Make sure a session exists for the extracted UUID
	sessionFound := findSession(uuid)
	
	if !sessionFound {
	
		http.Error(w, fmt.Sprintf("No session found for: %v.", uuid), http.StatusUnauthorized)
		return
	
	}

   b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
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
	msg := structs.Message{}
	err = json.Unmarshal(b, &msg)
	
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

	filter, code := mappings(msg.Message)

	

	if(strings.Compare(filter, "") == 0){

		msg := structs.Message{}

		msg.Message = "Oops! We could not get that. Can you try again!"

		output, err := json.Marshal(msg)
		
		if err != nil {
		
			http.Error(w, err.Error(), 500)
			return
		
		}

		w.Header().Set("content-type", "application/json")
		w.Write(output)	
		
		return

	}else{

		if(strings.Compare(filter, "team_id") == 0){

			if(code != 0){

				api.GetTeamMatches(w, r, b, code)			
			
			}else{

				msg := structs.Message{}

				msg.Message = "We either cannot seem to recognize this team or you are trying something outside of the Premier League! Please, try again!"

				output, err := json.Marshal(msg)
				
				if err != nil {
				
					http.Error(w, err.Error(), 500)
					return
				
				}

				w.Header().Set("content-type", "application/json")
				w.Write(output)	
				
				return

			}

		}else{

			if(strings.Compare(filter, "competition_id") == 0){

				if(code != 0){

					api.GetLeagueTable(w, r, b, code)
					
				}else{

					msg := structs.Message{}

					msg.Message = "We either cannot seem to recognize this team or you are trying something outside of the Premier League! Please, try again!"

					output, err := json.Marshal(msg)
					
					if err != nil {
					
						http.Error(w, err.Error(), 500)
						return
					
					}

					w.Header().Set("content-type", "application/json")
					w.Write(output)	
					
					return

				}
										

			}else{

				if(strings.Compare(filter, "team_profile") == 0){

					if(code != 0){

						api.GetTeamProfile(w, r, b, code)
						
					}else{

						msg := structs.Message{}

						msg.Message = "We either cannot seem to recognize this team or you are trying something outside of the Premier League! Please, try again!"

						output, err := json.Marshal(msg)
						
						if err != nil {
						
							http.Error(w, err.Error(), 500)
							return
						
						}

						w.Header().Set("content-type", "application/json")
						w.Write(output)	
						
						return

					}
						
					
				}else{

					if(strings.Compare(filter, "live") == 0){

						if(code == 0){

							api.GetLiveScores(w, r, b, 2)
						
						}else{

							msg := structs.Message{}

							msg.Message = "We either cannot seem to recognize this team or you are trying something outside of the Premier League! Please, try again!"

							output, err := json.Marshal(msg)
							
							if err != nil {
							
								http.Error(w, err.Error(), 500)
								return
							
							}

							w.Header().Set("content-type", "application/json")
							w.Write(output)	
							
							return

						}
						
					
					}else{

						msg := structs.Message{}

						msg.Message = "We either cannot seem to recognize this team or you are trying something outside of the Premier League! Please, try again!"

						output, err := json.Marshal(msg)
						
						if err != nil {
						
							http.Error(w, err.Error(), 500)
							return
						
						}

						w.Header().Set("content-type", "application/json")
						w.Write(output)	
						
						return

					}

				}

			}

		}

	}

}
