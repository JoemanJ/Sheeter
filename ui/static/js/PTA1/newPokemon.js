var speciesObj

window.onload = () => {
  for (form of document.getElementsByTagName("form")) {
    for (child of form.children) {
      child.required = true
    }
  }
  getSpecies()
  getAbilities()
  getCapacities()
  switchAFormDisplay()
  switchCFormDisplay()
}

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

function getSpeciesAbilities() {
  console.log(speciesObj)

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

  for (a of speciesObj[spc].Abilities) {
    console.log(a)
    let opt = document.createElement("option")
    opt.text = a.Name;
    opt.value = a.Name;
    ablt.appendChild(opt)
  }

  for (a of speciesObj[spc].HighAbilities) {
    console.log(a)
    let opt = document.createElement("option")
    opt.text = a.Name;
    opt.value = a.Name;
    ablt.appendChild(opt)
  }
}

function getCapacities() {
  var capacitiesObj

  fetch("/data/PTA1/capacityData").then(response => response.json()).then(function(data) {

    capacitiesObj = data
    c_list = document.getElementById("c_list")

    for (key of Object.keys(capacitiesObj)) {
      let ip = document.createElement("input")
      ip.type = "checkbox";
      ip.name = "c_" + key;

      let li = document.createElement("li")
      li.appendChild(ip)
      li.append(" " + key)

      c_list.appendChild(li)
    }
  });
}

function getAbilities() {
  var abilitiesObj

  fetch("/data/PTA1/abilityData").then(response => response.json()).then(function(data) {

    abilitiesObj = data
    a_list = document.getElementById("a_list")
    ha_list = document.getElementById("ha_list")


    for (key of Object.keys(abilitiesObj)) {
      let ip = document.createElement("input")
      ip.type = "checkbox";
      ip.name = "a_" + key;

      let li = document.createElement("li")
      li.appendChild(ip)
      li.append(" " + key)

      let h_ip = document.createElement("input");
      h_ip.type = "checkbox"
      h_ip.name = "ha_" + key;

      let h_li = document.createElement("li")
      h_li.appendChild(h_ip)
      h_li.append(" " + key)


      a_list.appendChild(li)
      ha_list.appendChild(h_li)
    }
  });
}

function switchCFormDisplay() {
  CForm = document.getElementById("capacity_form")

  if (document.getElementById("new_c").checked) {
    CForm.style.display = "flex";
    return;
  }
  CForm.style.display = "none"
}

function switchAFormDisplay() {
  Aform = document.getElementById("ability_form")

  if (document.getElementById("new_a").checked || document.getElementById("new_ha").checked) {
    Aform.style.display = "flex";
    return;
  }
  Aform.style.display = "none"
}
