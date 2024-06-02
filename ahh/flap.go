package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", game)
	http.ListenAndServe(":8080", nil)
}
func game(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte(`<head>
		<title>goofy bird</title>
	</head>
 	<body>
 		<style>
 			body {
 				font-family: "Comic Sans MS", "Comic Sans", cursive;
 				color:white;
 			}
 			.gamebg {
		background-color: lightblue;
		border: solid ;
		border-color: black;
		border-width: 1px;
		height: 100vh;
		left: 50%;
		margin: 0 auto;
		position: fixed;
		text-align: center;
		transform: translate(-50%, -1%);
		width: 330px;
	}
	.gametitle{
		color: black;
		background-color: yellow;
		border: solid ;
		text-align: center;
		width: 330px;
		transform: translate(-0.5%, 0%);
	}
	.bird { 
		position: fixed;
		top: 400px;
		right: 270px;
		width: 50px;
		height: 50px;
	}
	.pipe1 {
		z-index: -1;
		background-color: green;
		position: fixed;
		height: 25vh;

		right: 0px;
		width: 20px;
		top: 82.88px;
	}
	.pipe2 {
		background-color: blue;
		position: fixed;
		height: 25vh;

		right: 0px;
		width: 20px;
		bottom:0px;
	}
	.button {
		position:fixed;
		top: 5px;
		left: 5px;
	}
	h1 {
		color: white;
		text-shadow: -1px 0 black, 0 1px black, 1px 0 black, 0 -1px black;

		z-index: 1;
	}
	h2 {
		text-shadow: -1px 0 black, 0 1px black, 1px 0 black, 0 -1px black;

		z-index: 1;
	}
 		</style>
 		<div id="gamebg" class="gamebg" onclick="movement()">
 			<div class="gametitle"><h1 >goofy bird</h1></div>
			<form action="/" method="get">
			<input type="submit" value="Restart">
			</form>
 			<h2><span>Time Survived: <span class="score" id="score">0</span> Seconds</span></h2>
 			<h1 id="alert"></h1>
 			<img id="bird" class="bird" src="/static/bird.jpeg">
 			<div id="pipe2" class="pipe2"></div>
 			<div id="pipe1" class="pipe1"></div>
 		</div>

 	 <script>
 	 	const gamebg = document.getElementById("gamebg")
 	 	const alert = document.getElementById("alert")
 	 	const score = document.getElementById("score")
 	 	const scoretimer = setInterval(scoreEngine, 1000)
 	 	const gametimer = setInterval(gameEngine, 17)
 	 	function scoreEngine() {
 	 		score.innerHTML = eval(parseInt(score.innerHTML)+1)
 	 	}
 	 	function gameEngine(){
 	 		//html elements
 	 		const bird = document.getElementById("bird")
 	 		const birdcss = window.getComputedStyle(bird);
 	 		const birdheight = birdcss.getPropertyValue('top').replace('px', '')
 	 		const pipe1 = document.getElementById("pipe1")
 	 		const pipe1css = window.getComputedStyle(pipe1);
 	 		const pipe2 = document.getElementById("pipe2")
 	 		const pipe2css = window.getComputedStyle(pipe2);
 	 		//move pipe
 	 		const pipe1pos = pipe1css.getPropertyValue('right').replace('px', '')
 	 		pipe1.style.right = eval(parseInt(pipe1pos) + 2) + 'px' 
 	 		const pipe2pos = pipe2css.getPropertyValue('right').replace('px', '')
 	 		pipe2.style.right = eval(parseInt(pipe2pos) + 2) + 'px'
 	 		if (pipe1pos > 300){
 	 			pipe1.remove() ;
 	 			pipe2.remove() ;
 	 		}
 	 		//rng pipe height
 	 		const pipe1height = pipe1css.getPropertyValue('height').replace('px', '')
 	 		const pipe2height = pipe2css.getPropertyValue('height').replace('px', '')
 	 		if (pipe1height == '') {
 	 			const rng = Math.ceil(Math.random() * 25)
 	 			const makepipe1 = document.createElement("div")
 	 		makepipe1.id = "pipe1"
 	 		makepipe1.className = "pipe1"
 	 		gamebg.appendChild(makepipe1)
 	 		makepipe1.style.height = rng + 'vh'

 	 			const makepipe2 = document.createElement("div")
 	 		makepipe2.id = "pipe2"
 	 		makepipe2.className = "pipe2"
 	 		gamebg.appendChild(makepipe2)
 	 		makepipe2.style.height = eval(50 - rng) + 'vh'
 	 		}
 	 		//gravity
 	 		const birdpos = birdcss.getPropertyValue('top').replace('px', '')
 	 		bird.style.top = eval(parseInt(birdpos) + 2) + 'px' 
 	 		if (birdpos < 83) {
 	 			bird.style.top = 84 + 'px'
 	 		}
 	 		//lose
 	 		if (birdpos  > document.documentElement.clientHeight - 50) {
 	 		alert.innerHTML = "you lose (cringe)"
 	 		clearInterval(scoretimer)
 	 		clearInterval(gametimer)
 	 		}
 	 		//console.log(eval(parseInt(birdheight) + 50), birdheight - 82, window.innerHeight - pipe2height, pipe1height
 	 		if (pipe1pos > 250){
 	 			if (eval(parseInt(birdheight) - 82) < pipe1height|| eval(parseInt(birdheight) + 50) > window.innerHeight - pipe2height) {
 	 			alert.innerHTML = "you lose (cringe)"
 	 			clearInterval(gametimer)
 	 			clearInterval(scoretimer)
 	 			}
 	 			else {

 	 			}
 	 		}
 	 		
 	 	}

 	 	function movement() {
 	 		const bird = document.getElementById("bird")
 	 		const birdcss = window.getComputedStyle(bird);
 	 		const birdpos = birdcss.getPropertyValue('top').replace('px', '')
 	 		bird.style.top = eval(parseInt(birdpos) - 100) + 'px' 
 	 		
 	 		
 	 	}
 	 
	</script>
 	</body>`))
	}
}
