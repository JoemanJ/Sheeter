var classData
var talentData
var expertiseData
var itemData
var sheet_
var selectedPoke = ""

window.onload = () => {
  sheet_ = parseInt(document.getElementById("sheet").value)

  fetch("/data/PTA1/itemData").then(response => response.json()).then(data => {
    itemData = data

    let list = document.getElementById("item_list")

    for ( const key of Object.keys(data)){
      op = document.createElement("option")
      op.value = key
      op.innerHTML = key

      list.appendChild(op)
    }
  })

  fetch("/data/PTA1/classData").then(response => response.json()).then(data => {
    classData = data

    let classes = document.getElementsByClassName("class_select")
    console.log(classes)

    for (let i=0; i<4; i++){
      el = classes.item(i)

      for (key of Object.keys(classData)){
      op = document.createElement("option")
      op.value = key
      op.innerHTML = key

      el.appendChild(op)
      }
    }
  })
}

window.BeforeUnloadEvent() = () => {
  let data={id:0, form_name: "update", class1:"", class2:"", class3:"", class4:"", atkMod:0, defMod:0, spatkMod:0, spdefMod:0, spdMod:0, notes:""}

  data.id = sheet_

  data.class1 = document.getElementById("class_1").value
  data.class2 = document.getElementById("class_2").value
  data.class3 = document.getElementById("class_3").value
  data.class4 = document.getElementById("class_4").value
  
  data.atkMod = document.getElementById("ATK_mod").value
  data.defMod = document.getElementById("DEF_mod").value
  data.spatkMod = document.getElementById("SPATK_mod").value
  data.spdefMod = document.getElementById("SPDEF_mod").value
  data.spdMod = document.getElementById("SPD_mod").value

  data.notes = document.getElementById("notes_textbox").value

  let USP = new URLSearchParams()

  for(let [name, value] of Object.entries(data)){
    USP.append(name, value)
  }

  fetch("/sheet/?id="+data.id, {method: "POST", body: USP})
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

function switchBallIcons(tag) {
  let icon = tag.firstElementChild

  icon.src="/static/img/PTA1/Pokeball_open_icon.png"
  
  tag.onmouseout = () => icon.src="/static/img/PTA1/Pokeball_icon.png"
}

function openSheet(id){
  let w = window.screen.width.toString()
  let h = window.screen.height.toString()
  window.open("/sheet/?id="+id, "", "width="+w+", height="+h+", menubar=no, toolbar=no, status=no")
}

function switchClassFormDisplay(){
  let form = document.getElementById("class_form")
  let class1 = document.getElementById("class_1").value
  let class2 = document.getElementById("class_2").value
  let class3 = document.getElementById("class_3").value
  let class4 = document.getElementById("class_4").value

  if (class1=="new_class" || class2=="new_class" || class3=="new_class" || class4=="new_class"){
    form.style.display="flex"

    fetch("/data/PTA1/talentData").then(response => response.json()).then(data =>{
      talentData = data

      let talentList = document.getElementById("talents")

      for (key of Object.keys(talentData)){
        li = document.createElement("li")

        cb = document.createElement("input")
        cb.type="checkbox"
        cb.value="t_"+key
        cb.name="t_"+key

        li.appendChild(cb)
        li.append(" " + key)
        talentList.appendChild(li)

        op = document.createElement("option")
        op2 = document.createElement("option")
        op.value=key
        op2.value=key
        op.innerHTML=key
        op2.innerHTML=key

        document.getElementById("class_basic_talent1").appendChild(op)
        document.getElementById("class_basic_talent2").appendChild(op2)
      }
    })

    fetch("/data/PTA1/expertiseData").then(response => response.json()).then(data =>{
      expertiseData = data

      let expertiseList = document.getElementById("expertises")

      for (key of Object.keys(expertiseData)){
        li = document.createElement("li")

        cb = document.createElement("input")
        cb.type="checkbox"
        cb.value="e_"+key
        cb.name="e_"+key

        li.appendChild(cb)
        li.append(" " + key)
        expertiseList.appendChild(li)
      }
    })

    fetch("/data/PTA1/classData").then(response => response.json()).then(data =>{
      classData = data

      let classList = document.getElementById("parentClass")

      for (key of Object.keys(classData)){
        op = document.createElement("option")

        op.value = key
        op.innerHTML=key

        classList.appendChild(op)
      }
    })

    return
  }

  form.style.display="none"
}

function switchTalentFormDisplay(){
  let form = document.getElementById("talent_form");
  let bt1 = document.getElementById("class_basic_talent1").value
  let bt2 = document.getElementById("class_basic_talent2").value
  let nt = document.getElementById("new_talent").checked

  if (bt1=="new_talent" || bt2=="new_talent" || nt){
    form.style.display = "flex"


    return
  }

  form.style.display = "none"
}

function switchExpertiseFormDisplay(){
  let form = document.getElementById("expertise_form");
  let ne = document.getElementById("new_expertise").checked

  if (ne){
    form.style.display = "flex"
    return
  }

  form.style.display = "none"
}

function selectPoke(tag){
  if (!selectedPoke){
    selectedPoke = tag
    tag.classList.add("selected_poke")
    tag.classList.toggle("poke_box")
  }
  else{
    let USP = new URLSearchParams()

    let data = {id: sheet_, form_name: "switch_poke", poke1: selectedPoke.id, poke2: tag.id}
    selectedPoke.classList.remove("selected_poke")
    selectedPoke.classList.toggle("poke_box")

    for(let [name, value] of Object.entries(data)){
      USP.append(name, value)
    }

    fetch("/sheet/?id="+data.id, {method: "POST", body: USP})

    let aux = selectedPoke.innerHTML
    selectedPoke.innerHTML = tag.innerHTML
    tag.innerHTML = aux

    aux = selectedPoke.id
    selectedPoke.id = tag.id
    tag.id = aux

    selectedPoke = ""
  }
}

function changeNewItem(tag){
  let name = document.getElementById("item_name")
  let descr = document.getElementById("item_description")

  if (tag.value == "new"){
    name.style.display="block"
    descr.disabled = false
    return
  }

  name.style.display="none"
  name.value = ""
  descr.disabled = true

  descr.innerHTML = itemData[tag.value].Description  
}

function addItem(sheet, itemName, qtt, i, tag){
  let USP = new URLSearchParams()

  let data = {form_name: "item_form", id: sheet, i_name: itemName, factor: i, i_qtt: qtt, i_description: ""}
  data.i_description = document.getElementById("item_description").value

  let name
  let i_qtt

  if (!itemName){
    name = document.getElementById("item_name").value
    if (!name){
      name = document.getElementById("item_list").value
    }
    i_qtt = document.getElementById("item_qtt").value
  }else{
    name = itemName
    i_qtt = qtt.toString()
  }

  if (sheet){
    data.id=sheet
  }else{
    data.id=sheet_
  }
  data.i_name = name
  data.factor = i
  data.i_qtt = i_qtt

  for(let [name, value] of Object.entries(data)){
    USP.append(name, value)
  }

  fetch("/sheet/?id="+data.id, {method: "POST", body: USP})

  if (tag.parentNode.id != "X"){
    console.log(tag.parentNode.parentNode.children[0].children[0].value)
    tag.parentNode.parentNode.children[0].children[0].value = parseInt(tag.parentNode.parentNode.children[0].children[0].value) + data.i_qtt * i
    if (parseInt(tag.parentNode.parentNode.children[0].children[0].value) <= 0 ){
      // tag.parentNode.parentNode.children[0].children[0].value = "0"
    }
  }
}

function fetchDexData(){
  let dexSeen = document.getElementById("seen_pokemon").value
  let dexCaught = document.getElementById("caught_pokemon").value

  let search = document.getElementById("dex_search").value
  let sprite = document.getElementById("dex_sprite")
  let caughtIcon = document.getElementById("dex_caught_icon")
  let desc = document.getElementById("dex_poke_description")

  fetch("/data/specific?module=PTA2&type=species&id="+search).then(response => response.json()).then(data => {
    console.log(data)
    if (data.Number == 0){
      desc.innerHTML = "Espécie desconhecida"
      sprite.src = "/static/img/PTA1/Pokeball_icon.png"
      return
    }

    if (Array.from(dexSeen)[data.Number-1] == 0){
      sprite.style.filter = "brightness:(0%)"
      desc.innerHTML = "???"
    }else{
      sprite.style.filter = ""
      desc.innerHTML = data.Name
    }

    sprite.src = data.Sprite

    if (Array.from(dexCaught)[data.Number-1] == 0){
      caughtIcon.style.opacity = "0"
    }else{
      caughtIcon.style.opacity = "100%"
    }
      
  })
}
