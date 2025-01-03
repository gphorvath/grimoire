---
model: Claude-3.5-Sonnet-200K (or similar capable model)
input: GitHub-style release notes/changelog
output: Formatted, categorized release notes with emojis
version: 0.1.0
updated: 2024-11-02
author: Gregory Horvath
email: gregory@horvath.ai
tags:
   - code
   - review
---

# Release Notes Formatter Prompt

Purpose: Transform plain GitHub release notes into a well-organized, easily scannable changelog with consistent categorization and formatting.

Please reorganize these release notes for better readability. Follow these guidelines:

1. Group changes into these categories, using the specified emojis:
   - "🚀 New Features" (feat:)
   - "🐛 Fixes" (fix:)
   - "📖 Documentation" (docs:)
   - "🔧 Maintenance" (chore:)

2. Prioritize sections in this order:
   - Features first (most important user-facing changes)
   - Bug fixes
   - Documentation updates
   - Maintenance/chores

3. Within each category, sort changes by:
   - Impact on users (more impactful first)
   - Scope of change (broader changes before narrow ones)

4. Use markdown formatting with ### for section headers

5. Preserve all original metadata including:
   - Pull request numbers
   - Contributors
   - Links
   - Any "New Contributors" sections
   - Changelog links

6. The date of this release is {{.Date}}.

Please maintain the original technical details while making the notes more scannable and organized.

---

{{.UserPrompt}}
