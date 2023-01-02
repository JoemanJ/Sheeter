var trainerLvlTable

window.onload = () => {
  fetch("/data/PTA1/trainerLvlTable").then(response => response.json()).then(data => trainerLvlTable = data)
}

function calcMod(tag) {
  //Mudar os pontos restantes
  let remPts = document.getElementById("remaining_points")
  let lvl = parseInt(document.getElementById("lvl").value, 10)
  let hp = parseInt(document.getElementById("hp").value, 10)
  let atk = parseInt(document.getElementById("atk").value, 10)
  let def = parseInt(document.getElementById("def").value, 10)
  let spatk = parseInt(document.getElementById("spatk").value, 10)
  let spdef = parseInt(document.getElementById("spdef").value, 10)
  let spd = parseInt(document.getElementById("spd").value, 10)

  console.log(trainerLvlTable.total_status)
  console.log(lvl)

  remPts.innerHTML = 60 + trainerLvlTable.total_status[lvl] - (hp + atk + def + spatk + spdef + spd)

  let mod = document.getElementById(tag.id + "_mod")

  if (tag.value < 10) {
    mod.value = tag.value - 10
  } else {
    mod.value = Math.floor((tag.value - 10) / 2)
  }

  //Não deixar enviar se tiver pontos demais
  if (parseInt(remPts.innerHTML, 10) < 0) {
    document.getElementById("submit").disabled = true
  } else {
    document.getElementById("submit").disabled = false
  }
}
