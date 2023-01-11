var POKEMONEXPTABLE = [0, 25, 50, 100, 150, 200, 400, 500, 600, 1000, 1500, 2000,  3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 11500, 13000, 14500, 16000, 17500, 19000, 20500, 22000, 23500, 25000, 27500, 30000, 32500, 35000, 37500, 40000, 42500, 45000, 47500, 50000, 55000, 60000, 65000, 70000, 75000, 80000, 85000, 90000, 95000, 100000, 110000, 120000, 130000, 140000, 150000, 160000, 170000, 180000, 190000, 200000, 210000, 220000, 230000, 240000, 250000, 260000, 270000, 280000, 290000, 300000, 310000, 320000, 330000, 340000, 350000, 360000, 370000, 380000, 390000, 400000, 410000, 420000, 430000, 440000, 450000, 460000, 470000, 480000, 490000, 500000, 510000, 520000, 530000, 540000, 550000, 560000, 570000, 580000, 590000, 600000]
var sheet_
var movesData = ""
var remPoints

window.onload = () => {
  updateXpMeter()
  sheet_ = document.getElementById("sheet").value

  remPoints = (parseInt(document.getElementById("stats_total").innerHTML) - parseInt(document.getElementById("stats_allocated").innerHTML))

  fetch("/data/PTA1/moveData").then(response => response.json()).then(data => {
    movesData = data
    for (el of document.getElementsByClassName("move_name")){
      for(const key of Object.keys(movesData)){
        let op = document.createElement("option")
        op.value = key
        op.innerHTML = key
        el.appendChild(op)
      }
    }
  })
}

function openTab(event, tab_name){
    var tabs = document.getElementsByClassName("tab_body")
    
    for (let i=0; i<tabs.length; i++){
        tabs[i].style.display = "none"
        tabs[i].className.replace(" active", "")
    }

    event.currentTarget.className += " active"
    document.getElementById(tab_name + "_tab").style.display = "flex"
}

function updateXpMeter(){
  let meter = document.getElementById("xp_meter")

  meter.min = POKEMONEXPTABLE[parseInt(document.getElementById("lvl").innerHTML) -1 ]
  meter.max = POKEMONEXPTABLE[parseInt(document.getElementById("lvl").innerHTML)]
}

function switchMoveInfo(tag){

  let select = tag.getElementsByClassName("move_name")[0]
  let descriptors = tag.getElementsByClassName("move_descriptors")[0]
  let type = tag.getElementsByClassName("move_type")[0]
  let damage = tag.getElementsByClassName("move_damage")[0]
  let acc = tag.getElementsByClassName("move_accuracy")[0]
  let frequency = tag.getElementsByClassName("move_frequency")[0]
  let reach = tag.getElementsByClassName("move_reach")[0]
  let effect = tag.getElementsByClassName("move_effect")[0]
  
  let move = movesData[select.value]
  console.log(move)
  descriptors.innerHTML = move.Descriptors
  type.innerHTML = move.Type
  damage.innerHTML = move.Damage.X.toString() + " d" + move.Damage.N.toString() + " + " + move.Damage.Mod.toString()
  acc.innerHTML = move.AccDiff.toString()
  frequency.innerHTML = move.Frequency
  reach.innerHTML = move.Reach
  effect.innerHTML = move.Effect
}

function registerNewMove(){
  document.getElementById("new_move_send_button").disabled = true

  let name = document.getElementById("new_move_name").value
  let descriptors = document.getElementById("new_move_descriptors").value
  let type = document.getElementById("new_move_type").value
  let damage1 = document.getElementById("new_move_damage_1").value
  let damage2 = document.getElementById("new_move_damage_2").value
  let damage3 = document.getElementById("new_move_damage_3").value
  let acc = document.getElementById("new_move_accuracy").value
  let frequency = document.getElementById("new_move_frequency").value
  let reach = document.getElementById("new_move_reach").value
  let effect = document.getElementById("new_move_effect").value

  let data = {form_name:"new_move", name: name, descriptors: descriptors, type: type, damage1: damage1, damage2: damage2, damage3: damage3, acc: acc, frequency: frequency, reach: reach, effect: effect}

  let USP = new URLSearchParams()

  for (const [key, val] of Object.entries(data)){
    USP.append(key, val)
  }

  fetch("/sheet/?id="+sheet_, {method: "POST", body: USP}).then(response => {if (response.ok){window.location.reload()}})
}
