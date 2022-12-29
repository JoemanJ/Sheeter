window.onload = () => {
  for (child of document.getElementById("pokemon_form").children) {
    child.required = true
  }
  getSpecies()
}

var speciesObj

function getSpecies() {
  fetch("/data/PTA1/speciesData").then(response => response.json()).then(function(data) {

    speciesObj = data

    for (key of Object.keys(speciesObj)) {
      let opt = document.createElement("option")
      opt.text = key;
      opt.value = key;
      document.getElementById("species").appendChild(opt)
    }
  });
}

function getAbilities() {
  spc = document.getElementById("species").value
  ablt = document.getElementById("ability")
  spcForm = document.getElementById("species_form")

  if (spc == "new") {
    ablt.disabled = true
    spcForm.style.display = "flex"
    return;
  }

  ablt.disabled = false
  spcForm.style.display = "none"
  for (let i = 0; i < ablt.length; i++) {
    ablt.remove(i);
  }

  for (key of Object.keys(speciesObj[spc].Abilities)) {
    let opt = document.createElement("option")
    opt.text = key;
    opt.value = key;
    ablt.appendChild(opt)
  }
}
