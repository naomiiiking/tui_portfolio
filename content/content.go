package content

import "strings"

// Name is shown in the header on every screen.
const Name = "naomi king"

// About is your bio. Use \n for paragraph breaks.
const About = `Founding engineer for the AI Solutions team at Goodnotes. Modernising workflows and building internal, AI native applications for engineers, designers, insights and more.`

// Job is one entry in your work history.
type Job struct {
	Title    string
	Company  string
	Period   string
	Location string
	Bullets  []string
}
var Experience = []Job{
	{
		Title:    "AI Engineer",
		Company:  "Goodnotes",
		Period:   "2025 – present",
		Location: "London, UK",
		Bullets: []string{
			"Something you built or led",
			"Another thing worth mentioning",
		},
	},
}

type Skill struct {
	Title    string
	Examples []string
}

var Skills = []Skill{
	{
		Title: "AI Development",
		Examples: []string{
			"LLMs: Fine tuning, orchestraction, deployment locally and on cloud, Ollama, Hugging Face",
			"AI Frameworks: Langraph, AI SDK, Claude/OpenAI SDK",
			"RAG: Vector search pipelines with Chroma, FAISS, PGVector, AWS Bedrock",
			"MCP: RBAC tool usage through OAuth 2 and OIDC",
			"ML: Sckitlearn, PyTorch",
		},
	},
	{
		Title: "Software engineering",
		Examples: []string{
			"Programming: Python, Typescript, Go, SQL",
			"Servers: Express.js, Bun",
			"Chat: ChatSDK, Bolt",
			"Integrations: Fast API, webhooks",
		},
	},
	{
		Title: "Infrastructure",
		Examples: []string{
			"Cloud Compute: AWS, VPS hosting, Neoclouds, Google Cloud Services",
			"Containerisation: Docker, Kubernetes, ArgoCD",
			"GitOps: Terraform, CI/CD",
		},
	},
}

type Link struct {
	Label string
	URL   string
}

var Links = []Link{
	{Label: "github", URL: "https://github.com/naomiiiking"},
	{Label: "linkedin", URL: "https://www.linkedin.com/in/naomi-king-0374891ba/"},
	{Label: "website", URL: "https://"},
}

const rawPortrait = 
`......................................................................
.......................................................................
.......................................................................
.........................--.    .......................................
......................:=**=-=-::+=:-...................................
.....................=####*#@@@@@%*. %.................................
....................=##%%%%@@@@@%#*:  .................................
...................=#*+***#@#+*%#+=:   ................................
..................:#*=-=*%%#=.#@-:-:   ................................
..................=++::+@@@@@@@@+:-    ................................
.................:=--::+%@%=+@%=..+#..  ...............................
.................=:.:.-+%@+-*#+:. .:.  :...............................
................:-.:..:=%@@@@#-.  ..   :...............................
...................:..::*@##@%+..-.  . :...............................
..................:-.....   .++:..     ................................
.................:=+.:... .  .::.   .. ................................
................:=++--...     ..    ..  .+*#*-:........................
...........-=*##@##=-=:.::    ...   ..:=*@@@@@@*:......................
........=@@@@@@@@@@::-:.-##=**#%##.:::%@@@@@@@%*=:.....................
.......*@@@@@@@@@@.-=:::+@@@@%#@@@=--#+%@@@@@%#+-:.....................
......-@@@@@@@@@@:-*@@@@@@@@%**=*@=--=@@@@@@@%*=-:.....................
......*@@@@@@@@@*:-#@@@@@@@@@@-.:@=::-#@@@@@%#*+-:.....................
.....:%@@@@@@@@%===#@@@@@@@@@#*%%*.::::%@@#*##*=-:.....................
.....-@@@@@@@@@+==.=@@@@@@@@@#:     :-:.*+-:=#*=-:.....................
.....-@@@@@@@%*+=-:=@@@@@@@%-+.  ..-:---:-..:*+=-:.....................
.....=@@@@@@@#-+===:%@@@@@%-+@=   --::-::.. .==--::....................
.....+@@@@@@@*-===::#@@@#+:#@@+:*:    .::. ..-==--:....................
....:+@@@@@@%=++--@@@@%=.:%@#--..-    .:.   .:==--:....................
....:*@@@@@@#==#@@@@@*:..-+=-:.              .==--:....................
....:*@@@@@@@@@@@@@#=..                       :==--:...................
....:*@@@@@@@@@@@%+:.                       . .-=--:...................
....:*@@@@@@@@@%*:.                         .. .-----:.................
....:+@@@@@@%#=:..               ....       .. .-=++=-.................
.....-@@@%#+-...               .::::::     ...  :++**=:................
.....:*%*-....        .::. ..  :--:::.:.   ...  :*##*+:................
.............       .:::::. ...---::.....  .... :*##*=:................`

var Portrait = func() string {
	lines := strings.Split(rawPortrait, "\n")
	return strings.Join(lines, "\n")
}()
