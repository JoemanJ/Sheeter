package PTA1

import (
	"Joe/sheeter/pkg/general"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func MovesEndpoint(w http.ResponseWriter, r *http.Request){
  switch r.Method {
  case "GET":
    move, err := GetMove(r.URL.Query().Get("name"))
    if err != nil{
      w.WriteHeader(http.StatusNotFound)
      return
    }

    r, err := json.Marshal(move)
    if err != nil{
      w.WriteHeader(http.StatusInternalServerError)
      return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(r))
    
  case "POST":
    r.ParseForm()

    // fmt.Println(r.Form)

    errLog := ""

    name := r.FormValue("name")
    if name == ""{
      errLog = fmt.Sprintf("%s%s", errLog, "move name cannot be empty\n")
    }

    Type := r.FormValue("type")
    if Type == ""{
      errLog = fmt.Sprintf("%s%s", errLog, "type cannot be empty\n")
    }

    aptitude := r.FormValue("aptitude")
    descriptors := strings.Split( r.FormValue("descriptors"), "," )

    accDiff, err := strconv.Atoi( r.FormValue("accDif") )
    if err != nil{
      errLog = fmt.Sprintf("%s%s", errLog, "accDiff cannot be empty (if none, use 0)\n")
    }

    dmgType, err := strconv.Atoi( r.FormValue("dmgType") )
    if err != nil{
      errLog = fmt.Sprintf("%s%s", errLog, "accDiff cannot be empty (if none, use 0)\n")
    }

    diceQtt, err := strconv.Atoi( r.FormValue("diceQtt") )
    if err != nil{
      errLog = fmt.Sprintf("%s%s", errLog, "dice quantity cannot be empty (if none, use 0)\n")
    }

    diceSides, err := strconv.Atoi( r.FormValue("diceSides") )
    if err != nil{
      errLog = fmt.Sprintf("%s%s", errLog, "dice sides cannot be empty (if none, use 0)\n")
    }

    diceMod, err := strconv.Atoi( r.FormValue("mod") )
    if err != nil{
      errLog = fmt.Sprintf("%s%s", errLog, "move power modifier cannot be empty (if none, use 0)\n")
    }

    effect := r.FormValue("effect")
    // if effect == ""{
    //   errLog = fmt.Sprintf("%s%s", errLog, "move effect cannot be empty\n")
    // }

    if errLog != ""{
      w.WriteHeader(http.StatusBadRequest)
      // r, _ := json.Marshal(errLog)
      w.Write([]byte(errLog))
      return
    }

    reach := r.FormValue("reach")
    frequency := r.FormValue("frequency")
    contests := r.FormValue("contests")

    dice := sheeters.CreateDiceSet(diceQtt, diceSides, diceMod)

    move, err := RegisterMove(name, Type, aptitude, descriptors, accDiff, dmgType, dice, reach, frequency, contests, effect)
    if err != nil{
      w.WriteHeader(http.StatusInternalServerError)
      return
    }

    r, _ := json.Marshal(move)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(r))

  case "PUT":

  default:
    w.WriteHeader(http.StatusMethodNotAllowed)
    
  }
}
