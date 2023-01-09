var navLinks = document.querySelectorAll("nav a");
for (var i = 0; i < navLinks.length; i++) {
	var link = navLinks[i]
	if (link.getAttribute('href') == window.location.pathname) {
		link.classList.add("live");
		break;
	}
}

function Roll(){
  let n = parseInt(document.getElementById("header_dice_qtt").value)
  let d = parseInt(document.getElementById("header_dice_sides").value)
  let x = parseInt(document.getElementById("header_dice_mod").value)
  let result = document.getElementById("header_roll_result")

  let sum = 0

  for (let i=0; i<n; i++){
    let die = Math.floor(Math.random()*d + 1)
    sum += die
  }

  result.innerHTML = sum + x
}
