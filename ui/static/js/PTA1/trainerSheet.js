var classData
var talentData
var expertiseData
var itemData
var sheet_
var selectedPoke = ""
var remPoints

window.onload = () => {
  sheet_ = parseInt(document.getElementById("sheet").value)
  remPoints = (parseInt(document.getElementById("stats_total").innerHTML) - parseInt(document.getElementById("stats_allocated").innerHTML))


  if (remPoints > 0){
    document.getElementById("finish_stat_allocation_button").disabled = false
    let buttons = document.getElementsByClassName("stat_allocate_button")

    for(let i=0; i<buttons.length; i++){
      let btn = buttons.item(i)

      btn.style.display = "inline"
      btn.disabled = false
    }
  }

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

window.onbeforeunload = () => {
  let data={id:0, form_name: "update", class1:"", class2:"", class3:"", class4:"", hp:0, atkStage:0, defStage:0, spatkStage:0, spdefStage:0, spdStage:0, notes:""}

  data.id = sheet_

  data.class1 = document.getElementById("class_1").value
  data.class2 = document.getElementById("class_2").value
  data.class3 = document.getElementById("class_3").value
  data.class4 = document.getElementById("class_4").value
  
  data.hp = document.getElementById("current_hp").value
  data.atkStage = document.getElementById("ATK_stage").value
  data.defStage = document.getElementById("DEF_stage").value
  data.spatkStage = document.getElementById("SPATK_stage").value
  data.spdefStage = document.getElementById("SPDEF_stage").value
  data.spdStage = document.getElementById("SPD_stage").value

  data.notes = document.getElementById("notes_textbox").value

  let USP = new URLSearchParams()

  for(let [name, value] of Object.entries(data)){
    USP.append(name, value)
  }

  fetch("/sheet/?id="+data.id, {method: "POST", body: USP})
}

var statsAllocated = {"form_name":"allocate_stats", "HP":0, "ATK":0, "DEF":0, "SPATK":0, "SPDEF":0, "SPD":0}
function allocateStat(stat, qtt, tag){
  let stat_text = tag.parentElement.children.item(1)
  
  if ((remPoints > 0 && qtt == 1) || (qtt == -1 && statsAllocated[stat] > 0)){
    statsAllocated[stat] += qtt
    stat_text.value = parseInt(stat_text.value) + qtt
    remPoints -= qtt

    let aux = document.getElementById("stats_allocated")
    aux.innerHTML = parseInt(aux.innerHTML) + qtt
  }

  if (statsAllocated[stat] > 0){
    stat_text.style.color = "green"
  } else{
    stat_text.style.color = "black"
  }
}

function finishStatAllocation(){
  let USP = new URLSearchParams()

  for(const [key, value] of Object.entries(statsAllocated)){
    USP.append(key, value)
  }

  document.getElementById("finish_stat_allocation_button").disabled = true
  fetch("/sheet/?id="+sheet_, {method: "POST", body:USP}).then(response => {if(response.ok) {window.location.reload()}})
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

function switchClassFormDisplay(tag){
  let form = document.getElementById("class_form")
  let classInfo = document.getElementById("new_class_info")
  let class1 = document.getElementById("class_1").value
  let class2 = document.getElementById("class_2").value
  let class3 = document.getElementById("class_3").value
  let class4 = document.getElementById("class_4").value

  if (class1=="new_class" || class2=="new_class" || class3=="new_class" || class4=="new_class"){
    classInfo.style.display = "none"
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

      let classList = document.getElementById("class_parent")

      for (key of Object.keys(classData)){
        op = document.createElement("option")

        op.value = key
        op.innerHTML=key

        classList.appendChild(op)
      }
    })

    return
  }

  classInfo.style.display = "flex"
  form.style.display="none"

  fetch("/data/specific/?module=PTA2&type=trainerClass&id=" + tag.value).then(response => response.json()).then(data => {
    console.log(data)
    document.getElementById("new_class_info_name").innerHTML = data.Name
    document.getElementById("new_class_info_description").innerHTML = data.Description
    document.getElementById("new_class_info_parent_class").innerHTML = data.ParentClass
    document.getElementById("new_class_info_requisites").innerHTML = data.Requirements

    let expertiseList = document.getElementById("new_class_info_expertises")
    expertiseList.innerHTML = ""
    for (exp of data.Expertises){
      let span = document.createElement("span")

      let checkbox = document.createElement("input")
      checkbox.type = "checkbox"
      checkbox.id = "expertise_"+exp.Name
      checkbox.name = "expertise_"+exp.Name

      let label = document.createElement("label")
      label.for = checkbox.id
      label.innerHTML = exp.Name

      span.appendChild(checkbox)
      span.appendChild(label)
      expertiseList.appendChild(span)
    }

    let talentsList = document.getElementById("new_class_info_talents")
    talentsList.innerHTML = ""
    for (tal of data.BasicTalents){
      let talentBox = document.createElement("div")
      let str = `\
<div class="talent_box_header">\
  <div class="talent_box_info">\
    <h2 class="talent_name" readonly>${tal.Name}</h2>\
    <div class="talent_box_middle">\
      <h2 class="talent_target" readonly>${tal.Target}</h2>\
      <h2 class="talent_frequency" readonly>${tal.Frequency}</h2>\
    </div>\
    <textarea class="talent_description" name="description" readonly>${tal.Description}</textarea>\
  </div>\
  <div class="talent_icons">\
    <img src="/static/img/PTA1/Pokeball_icon.png">\
  </div>\
</div>`

      talentBox.classList.add("talent_box")
      talentBox.innerHTML = str

      talentsList.appendChild(talentBox)
    }

    for (tal of data.PossibleTalents){
      let talentBox = document.createElement("div")
      let str = `\
<div class="talent_box_header">\
  <div class="talent_box_info">\
    <h2 class="talent_name" readonly>${tal.Name}</h2>\
    <div class="talent_box_middle">\
      <h2 class="talent_target" readonly>${tal.Target}</h2>\
      <h2 class="talent_frequency" readonly>${tal.Frequency}</h2>\
    </div>\
    <textarea class="talent_description" name="description" readonly>${tal.Description}</textarea>\
  </div>\
  <div class="talent_icons">\
    <img src="/static/img/PTA1/Pokeball_icon.png">\
  </div>\
</div>`

      talentBox.classList.add("talent_box")
      talentBox.innerHTML = str

      talentsList.appendChild(talentBox)
    }
      talentsList.children[0].style.backgroundColor="#A0F33F"
      talentsList.children[1].style.backgroundColor="#A0F33F"

  })

}

function selectClass(tag){
  tag.disabled = true
  let USP = new URLSearchParams()
  let data = {id:sheet_, form_name:"add_class", class: document.getElementById("new_class_info_name").innerHTML}
  let expertises = document.getElementById("new_class_info_expertises").getElementsByTagName("input")
  for (ex of expertises){
    if (ex.checked){
      data[ex.name]=ex.checked
    }
  }

  for (const [key, val] of Object.entries(data)){
    USP.append(key, val)
  }

  fetch('/sheet/?id='+sheet_, {method: "POST", body: USP}).then(response => {
    if (response.ok) {window.location.reload()}
    else{alert("Erro ao adicionar classe. Por favor tente novamente mais tarde.")}
  })
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
    
    aux = selectedPoke.ondblclick
    selectedPoke.ondblclick = tag.ondblclick
    tag.ondblclick = aux

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

  let name = document.getElementById("dex_poke_name")
  let sprite = document.getElementById("dex_sprite")
  let caughtIcon = document.getElementById("dex_caught_icon")
  let desc = document.getElementById("dex_poke_description")
  let search = document.getElementById("dex_search").value

  fetch("/data/specific?module=PTA2&type=species&id="+search).then(response => response.json()).then(data => {
    console.log(data)
    if (data.Number == 0){
      sprite.style.filter = "brightness:(0%)"
      desc.innerHTML = "Espécie desconhecida"
      sprite.src = "/static/img/PTA1/Pokeball_icon.png"
      caughtIcon.style.opacity = "0"
      return
    }

    if (Array.from(dexSeen)[data.Number-1] == 0){
      sprite.style.filter = "brightness(0)"
      desc.innerHTML = "???"
    }else{
      sprite.style.filter = ""
      desc.innerHTML = data.Name
    }

    name.innerHTML = data.Name
    sprite.src = data.Sprite

    if (Array.from(dexCaught)[data.Number-1] == 0){
      caughtIcon.style.opacity = "0"
    }else{
      caughtIcon.style.opacity = "100%"
    }
      
  })
}

function addClass(n, tag){
  tag.style.display = "none"
  document.getElementById("class_" + n).hidden = false
}
