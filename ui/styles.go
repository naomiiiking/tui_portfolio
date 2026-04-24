package ui

import "github.com/charmbracelet/lipgloss"

var (
	colorText   = lipgloss.AdaptiveColor{Light: "#111827", Dark: "#c72c41"}
	colorAccent = lipgloss.AdaptiveColor{Light: "#6D28D9", Dark: "#dedede"}
	colorMuted  = lipgloss.AdaptiveColor{Light: "#6B7280", Dark: "#6B7280"}
	colorDim    = lipgloss.AdaptiveColor{Light: "#D1D5DB", Dark: "#374151"}

	nameStyle = lipgloss.NewStyle().
			Bold(true)

	activeTabStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorAccent).
			Underline(true)

	inactiveTabStyle = lipgloss.NewStyle().
				Foreground(colorMuted)

	divStyle = lipgloss.NewStyle().
			Foreground(colorDim)

	hintStyle = lipgloss.NewStyle().
			Foreground(colorMuted)

	labelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorAccent)

)
