package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"terminal-website/content"
)

type section int

const (
	sectionAbout section = iota
	sectionExperience
	sectionLinks
	sectionCount
)

var sectionNames = []string{"about", "projects", "links"}

// Model is the root Bubbletea model.
type Model struct {
	active section
	width  int
	height int
}

func New(width, height int) Model {
	if width == 0 {
		width = 80
	}
	if height == 0 {
		height = 24
	}
	return Model{width: width, height: height}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab", "right", "l":
			m.active = (m.active + 1) % sectionCount
		case "shift+tab", "left", "h":
			m.active = (m.active - 1 + sectionCount) % sectionCount
		}
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	}
	return m, nil
}

func (m Model) View() string {
	w := m.width
	div := divStyle.Render(strings.Repeat("─", w))

	return strings.Join([]string{
		m.header(w),
		div,
		m.tabBar(),
		div,
		m.body(w),
		div,
		m.footer(),
	}, "\n")
}

func (m Model) header(w int) string {
	name := nameStyle.Render(content.Name)
	quit := hintStyle.Render("q · quit")
	gap := w - lipgloss.Width(name) - lipgloss.Width(quit)
	if gap < 0 {
		gap = 0
	}
	return name + strings.Repeat(" ", gap) + quit
}

func (m Model) tabBar() string {
	var items []string
	for i, label := range sectionNames {
		if section(i) == m.active {
			items = append(items, activeTabStyle.Render(label))
		} else {
			items = append(items, inactiveTabStyle.Render(label))
		}
	}
	return "  " + strings.Join(items, "   ")
}

func (m Model) body(w int) string {
	const portraitW = 74
	const gap = 2
	innerW := w - 4
	showPortrait := innerW >= portraitW+gap+30

	contentW := innerW
	if showPortrait {
		contentW = innerW - portraitW - gap
	}

	var contentStr string
	switch m.active {
	case sectionAbout:
		contentStr = m.aboutBody(contentW)
	case sectionExperience:
		contentStr = m.experienceBody(contentW)
	case sectionLinks:
		contentStr = m.linksBody(contentW)
	}

	if !showPortrait {
		return lipgloss.NewStyle().Padding(1, 2).Render(contentStr)
	}
	portrait := lipgloss.NewStyle().Width(portraitW).Render(content.Portrait)
	raw := lipgloss.JoinHorizontal(lipgloss.Top, portrait, strings.Repeat(" ", gap), contentStr)
	return lipgloss.NewStyle().Padding(1, 2).Render(raw)
}

func (m Model) aboutBody(w int) string {
	var sb strings.Builder
	sb.WriteString(lipgloss.NewStyle().Width(w).Render(content.Title))
	sb.WriteString("\n\n")

	sb.WriteString(content.About)
	sb.WriteString("\n\n")

	sb.WriteString(labelStyle.Render("EXPERIENCE") + "\n")
	for i, job := range content.Experience {
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString("  " + labelStyle.Render(job.Title) + "\n")
		sb.WriteString("  " + hintStyle.Render(
			fmt.Sprintf("%s  ·  %s  ·  %s", job.Company, job.Period, job.Location),
		) + "\n")
		for _, b := range job.Bullets {
			sb.WriteString("    · " + b + "\n")
		}
	}
	sb.WriteString("\n")

	sb.WriteString(labelStyle.Render("SKILLS") + "\n")
	for _, skill := range content.Skills {
		sb.WriteString("  " + labelStyle.Render(skill.Title) + "\n")
		for _, ex := range skill.Examples {
			sb.WriteString("    · " + ex + "\n")
		}
	}

	return lipgloss.NewStyle().Width(w).Render(sb.String())
}

// Experience page todo: rebrand to projects
func (m Model) experienceBody(w int) string {
	var sb strings.Builder
	for i, job := range content.Experience {
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(labelStyle.Render(job.Title) + "\n")
		sb.WriteString(hintStyle.Render(
			fmt.Sprintf("%s  ·  %s  ·  %s", job.Company, job.Period, job.Location),
		) + "\n")
		for _, b := range job.Bullets {
			sb.WriteString("  · " + b + "\n")
		}
	}
	return lipgloss.NewStyle().Width(w).Render(sb.String())
}

func (m Model) linksBody(w int) string {
	var sb strings.Builder
	for _, link := range content.Links {
		sb.WriteString(fmt.Sprintf("%s  %s\n\n",
			labelStyle.Width(12).Render(link.Label),
			hintStyle.Render(link.URL),
		))
	}
	return lipgloss.NewStyle().Width(w).Render(sb.String())
}

func (m Model) footer() string {
	return "  " + hintStyle.Render("← →  navigate   tab  next section")
}
