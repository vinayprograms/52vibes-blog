// Package export provides markdown generation from session data.
package export

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"crush-export/internal/model"
)

// Markdown generates a markdown document from a session.
func Markdown(session *model.Session) string {
	var sb strings.Builder

	// Header
	sb.WriteString(fmt.Sprintf("# Session: %s\n\n", session.Title))
	sb.WriteString(fmt.Sprintf("**Date**: %s\n", session.CreatedAt.Format("2006-01-02 15:04:05")))

	stats := session.ComputeStatistics()
	sb.WriteString(fmt.Sprintf("**Duration**: %s\n\n", formatDuration(stats.Duration)))

	// Statistics table
	sb.WriteString("## Statistics\n\n")
	sb.WriteString("| Metric | Value |\n")
	sb.WriteString("|--------|-------|\n")
	sb.WriteString(fmt.Sprintf("| Messages | %d |\n", stats.MessageCount))
	sb.WriteString(fmt.Sprintf("| User Messages | %d |\n", stats.UserMessages))
	sb.WriteString(fmt.Sprintf("| Assistant Messages | %d |\n", stats.AssistantMsgs))
	sb.WriteString(fmt.Sprintf("| Files Created | %d |\n", stats.FilesCreated))
	sb.WriteString(fmt.Sprintf("| Files Modified | %d |\n", stats.FilesModified))
	sb.WriteString(fmt.Sprintf("| Errors | %d |\n\n", stats.ErrorCount))

	// Lessons Learned placeholder
	sb.WriteString("## Lessons Learned\n\n")
	sb.WriteString("<!-- Add your learnings here -->\n\n")

	// Key Exchanges (first few significant exchanges)
	keyExchanges := selectKeyExchanges(session.Messages)
	if len(keyExchanges) > 0 {
		sb.WriteString("## Key Exchanges\n\n")
		for _, msg := range keyExchanges {
			writeMessage(&sb, msg)
		}
	}

	// Full Conversation Log
	sb.WriteString("---\n\n## Full Conversation Log\n\n")
	for _, msg := range session.Messages {
		writeMessage(&sb, msg)
	}

	return sb.String()
}

func writeMessage(sb *strings.Builder, msg model.Message) {
	role := titleCase(msg.Role)
	sb.WriteString(fmt.Sprintf("### %s\n\n", role))

	content := msg.Content
	if content != "" {
		sb.WriteString(content)
		sb.WriteString("\n\n")
	}
}

func titleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func formatDuration(d interface{ Hours() float64; Minutes() float64; Seconds() float64 }) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
	}
	if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	}
	return fmt.Sprintf("%ds", seconds)
}

// selectKeyExchanges picks the most significant message exchanges.
// Selection criteria:
// - First user message (context setting)
// - Messages containing questions (user) and their answers (assistant)
// - Final assistant message (conclusion)
func selectKeyExchanges(messages []model.Message) []model.Message {
	if len(messages) < 4 {
		return nil // Too few messages for key exchanges
	}

	var key []model.Message
	seen := make(map[string]bool)

	// First user message
	for _, m := range messages {
		if m.Role == "user" && !seen[m.ID] {
			key = append(key, m)
			seen[m.ID] = true
			break
		}
	}

	// First assistant response
	for _, m := range messages {
		if m.Role == "assistant" && !seen[m.ID] {
			key = append(key, m)
			seen[m.ID] = true
			break
		}
	}

	// Messages with questions
	questionPattern := regexp.MustCompile(`\?`)
	for i, m := range messages {
		if len(key) >= 6 {
			break
		}
		if m.Role == "user" && questionPattern.MatchString(m.Content) && !seen[m.ID] {
			key = append(key, m)
			seen[m.ID] = true
			// Include following assistant response
			if i+1 < len(messages) && messages[i+1].Role == "assistant" && !seen[messages[i+1].ID] {
				key = append(key, messages[i+1])
				seen[messages[i+1].ID] = true
			}
		}
	}

	// Last assistant message
	for i := len(messages) - 1; i >= 0; i-- {
		if messages[i].Role == "assistant" && !seen[messages[i].ID] {
			key = append(key, messages[i])
			break
		}
	}

	return key
}

// Filename generates a safe filename for the session export.
func Filename(session *model.Session) string {
	timestamp := session.CreatedAt.Format("2006-01-02_15-04-05")
	slug := slugify(session.Title)
	if len(slug) > 50 {
		slug = slug[:50]
	}
	return fmt.Sprintf("%s_%s.md", timestamp, slug)
}

// slugify converts a title to a URL-safe slug.
func slugify(s string) string {
	// Lowercase
	s = strings.ToLower(s)

	// Replace spaces with hyphens
	s = strings.ReplaceAll(s, " ", "-")

	// Remove non-alphanumeric except hyphens
	var result strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			result.WriteRune(r)
		}
	}

	// Collapse multiple hyphens
	s = result.String()
	for strings.Contains(s, "--") {
		s = strings.ReplaceAll(s, "--", "-")
	}

	// Trim leading/trailing hyphens
	s = strings.Trim(s, "-")

	if s == "" {
		return "untitled"
	}

	return s
}

// FilenameWithSuffix appends a suffix for collision handling.
func FilenameWithSuffix(session *model.Session, suffix int) string {
	timestamp := session.CreatedAt.Format("2006-01-02_15-04-05")
	slug := slugify(session.Title)
	if len(slug) > 50 {
		slug = slug[:50]
	}
	return fmt.Sprintf("%s_%s-%d.md", timestamp, slug, suffix)
}
