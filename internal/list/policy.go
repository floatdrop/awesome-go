package list

import (
	"fmt"
	"time"
)

// Policy is the machine-enforced part of the entry rules. It lives in
// awesome.json so contributors can read the thresholds next to the data.
type Policy struct {
	MinAgeDays           int  `json:"min_age_days"`
	MinStars             int  `json:"min_stars"`
	MaxInactiveDays      int  `json:"max_inactive_days"`
	MaxDescriptionLength int  `json:"max_description_length"`
	RequireLicense       bool `json:"require_license"`
	RequireGo            bool `json:"require_go"`
	AllowForks           bool `json:"allow_forks"`
}

// DefaultMaxDescriptionLength applies when the policy leaves it unset.
const DefaultMaxDescriptionLength = 100

func (p Policy) maxDescription() int {
	if p.MaxDescriptionLength <= 0 {
		return DefaultMaxDescriptionLength
	}
	return p.MaxDescriptionLength
}

func daysBetween(from, to time.Time) int {
	return int(to.Sub(from).Hours() / 24)
}

// Dead reports whether a repository must be removed from the list regardless
// of anything else, and why.
func (p Policy) Dead(m RepoMeta) (reason string, dead bool) {
	switch {
	case m.NotFound:
		return "repository no longer exists on GitHub", true
	case m.Archived:
		return "repository is archived", true
	case m.Disabled:
		return "repository is disabled", true
	}
	return "", false
}

// Stale reports whether the repository has been inactive for longer than the
// policy tolerates. Stale entries are reported, not removed automatically:
// a finished library is not the same thing as a dead one.
func (p Policy) Stale(m RepoMeta, now time.Time) (days int, stale bool) {
	if p.MaxInactiveDays <= 0 || m.PushedAt.IsZero() {
		return 0, false
	}
	days = daysBetween(m.PushedAt, now)
	return days, days > p.MaxInactiveDays
}

// Check returns every way a repository violates the policy for a *new* entry,
// honouring the entry's exemptions. An empty result means it may be listed.
func (p Policy) Check(m RepoMeta, now time.Time, e Entry) []string {
	var problems []string
	if reason, dead := p.Dead(m); dead {
		problems = append(problems, reason)
		if m.NotFound {
			return problems
		}
	}
	if m.Fork && !p.AllowForks && !e.IsExempt(ExemptFork) {
		problems = append(problems, "repository is a fork")
	}
	if p.MinAgeDays > 0 && !m.CreatedAt.IsZero() {
		if age := daysBetween(m.CreatedAt, now); age < p.MinAgeDays {
			problems = append(problems, fmt.Sprintf("repository is %d days old, must be at least %d", age, p.MinAgeDays))
		}
	}
	if p.MinStars > 0 && m.Stars < p.MinStars {
		problems = append(problems, fmt.Sprintf("repository has %d stars, needs at least %d", m.Stars, p.MinStars))
	}
	if days, stale := p.Stale(m, now); stale && !e.IsExempt(ExemptInactive) {
		problems = append(problems, fmt.Sprintf("last push was %d days ago, must be within %d", days, p.MaxInactiveDays))
	}
	if p.RequireLicense && m.License == "" {
		problems = append(problems, "no license detected by GitHub")
	}
	if p.RequireGo && !m.IsGo {
		problems = append(problems, "primary language is not Go and there is no go.mod at the repository root")
	}
	return problems
}
