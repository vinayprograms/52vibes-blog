package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"crush-export/internal/db"
	"crush-export/internal/export"
)

var (
	Version   = "dev"
	Commit    = "unknown"
	BuildDate = "unknown"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Define flags
	dbPath := flag.String("db", ".crush/crush.db", "Path to Crush database")
	outDir := flag.String("out", "sessions", "Output directory")
	sessionID := flag.String("session", "", "Export specific session ID")
	listOnly := flag.Bool("l", false, "List sessions only")
	showVersion := flag.Bool("version", false, "Show version")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: crush-export [options]\n\n")
		fmt.Fprintf(os.Stderr, "Export Crush AI sessions to markdown files.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Check for unknown flags
	if flag.NArg() > 0 {
		return fmt.Errorf("unknown argument: %s", flag.Arg(0))
	}

	if *showVersion {
		fmt.Printf("crush-export %s\n", Version)
		fmt.Printf("  commit: %s\n", Commit)
		fmt.Printf("  built:  %s\n", BuildDate)
		return nil
	}

	// Check database exists
	if _, err := os.Stat(*dbPath); os.IsNotExist(err) {
		return fmt.Errorf("database not found: %s", *dbPath)
	}

	// Open repository
	repo, err := db.NewSQLiteRepository(*dbPath)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer repo.Close()

	ctx := context.Background()

	if *listOnly {
		return listSessions(ctx, repo)
	}

	return exportSessions(ctx, repo, *outDir, *sessionID)
}

func listSessions(ctx context.Context, repo *db.SQLiteRepository) error {
	sessions, err := repo.ListSessions(ctx)
	if err != nil {
		return err
	}

	if len(sessions) == 0 {
		fmt.Println("No sessions found.")
		return nil
	}

	fmt.Printf("%-40s  %-8s  %-20s  %s\n", "ID", "Messages", "Created", "Title")
	fmt.Println(strings.Repeat("-", 100))

	for _, s := range sessions {
		title := s.Title
		if len(title) > 40 {
			title = title[:37] + "..."
		}
		fmt.Printf("%-40s  %-8d  %-20s  %s\n",
			s.ID,
			s.MessageCount,
			s.CreatedAt.Format("2006-01-02 15:04:05"),
			title,
		)
	}

	fmt.Printf("\nTotal: %d sessions\n", len(sessions))
	return nil
}

func exportSessions(ctx context.Context, repo *db.SQLiteRepository, outDir, sessionID string) error {
	// Create output directory
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	// Get sessions to export
	sessions, err := repo.ListSessions(ctx)
	if err != nil {
		return err
	}

	// Filter to specific session if requested
	if sessionID != "" {
		var found bool
		for _, s := range sessions {
			if s.ID == sessionID {
				sessions = sessions[:0]
				sessions = append(sessions, s)
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("session not found: %s", sessionID)
		}
	}

	if len(sessions) == 0 {
		fmt.Println("No sessions to export.")
		return nil
	}

	// Track filenames to handle collisions
	usedFilenames := make(map[string]bool)
	exported := 0

	for _, summary := range sessions {
		session, err := repo.GetSession(ctx, summary.ID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: skip session %s: %v\n", summary.ID, err)
			continue
		}

		// Skip empty sessions
		if len(session.Messages) == 0 {
			continue
		}

		// Generate filename
		filename := export.Filename(session)

		// Handle collisions
		suffix := 1
		for usedFilenames[filename] {
			suffix++
			filename = export.FilenameWithSuffix(session, suffix)
		}
		usedFilenames[filename] = true

		// Generate markdown
		md := export.Markdown(session)

		// Write file
		path := filepath.Join(outDir, filename)
		if err := os.WriteFile(path, []byte(md), 0644); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}

		fmt.Printf("Exported: %s\n", filename)
		exported++
	}

	fmt.Printf("\nExported %d sessions to %s\n", exported, outDir)
	return nil
}
