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

var sectionNames = []string{"about", "experience", "links"}

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
	innerW := w - 4
	var raw string
	switch m.active {
	case sectionAbout:
		raw = lipgloss.NewStyle().Width(innerW).Render(content.About)
	case sectionExperience:
		raw = m.experienceBody(innerW)
	case sectionLinks:
		raw = m.linksBody(innerW)
	}
	return lipgloss.NewStyle().Padding(1, 2).Render(raw)
}

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
