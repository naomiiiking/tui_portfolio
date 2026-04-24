package content

// Name is shown in the header on every screen.
const Name = "naomi king"

// About is your bio. Use \n for paragraph breaks.
const About = `Replace this with a few sentences about who you are.

Where you're based, what you work on, what you care about.`

// Job is one entry in your work history.
type Job struct {
	Title    string
	Company  string
	Period   string
	Location string
	Bullets  []string
}

// Experience lists your roles, newest first.
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

type Link struct {
	Label string
	URL   string
}
var Links = []Link{
	{Label: "github",  URL: "https://github.com/naomiiiking"},
	{Label: "linkedin",  URL: "https://www.linkedin.com/in/naomi-king-0374891ba/"},
	{Label: "website", URL: "https://"},
}
