function openTab(event, tab_name){
    var tabs = document.getElementsByClassName("tab_body")
    
    for (let i=0; i<tabs.length; i++){
        tabs[i].style.display = "none"
        tabs[i].className.replace(" active", "")
    }

    event.currentTarget.className += " active"
    document.getElementById(tab_name + "_tab").style.display = "flex"
}

function openSheet(id){
  let w = window.screen.width.toString()
  let h = window.screen.height.toString()
  window.open("/sheet", "", "width="+w+", height="+h)
}
