#!/bin/bash

# 🎓 Calculate coverage for academic module files only
# Usage: ./academic_coverage.sh

# List of target Go files
ACADEMIC_FILES=(
    "registrar.go"
    "student.go"
    "course.go"
    "teacher.go"
    "enroll.go"
    "attendence.go"
)

echo "🧪 Running tests and generating coverage profile..."

# Run tests and create full coverage profile
go test -coverprofile=coverage.out ./internal

# Exit if tests failed
if [ $? -ne 0 ]; then
    echo "❌ Tests failed!"
    exit 1
fi

echo "📊 Filtering coverage for academic module files only..."

# Join academic filenames into regex pattern
PATTERN=""
for file in "${ACADEMIC_FILES[@]}"; do
    [[ -z "$PATTERN" ]] && PATTERN="$file" || PATTERN="$PATTERN|$file"
done

echo "🔍 Matching files using pattern: $PATTERN"

# Filter relevant lines for academic files
awk -v pat="$PATTERN" 'NR==1 || $0 ~ pat' coverage.out > academic_coverage.out

# Check that file isn't empty
if [ ! -s academic_coverage.out ] || [ "$(wc -l < academic_coverage.out)" -le 1 ]; then
    echo "⚠️  No coverage data found for academic files!"
    echo "📁 Files in coverage.out:"
    grep -E "\.go:" coverage.out | cut -d: -f1 | sort | uniq
    exit 1
fi

echo "✅ Academic coverage files detected:"
grep -E "\.go:" academic_coverage.out | cut -d: -f1 | sort | uniq

echo ""
echo "📈 Coverage Report for Academic Module:"
go tool cover -func=academic_coverage.out

echo ""
echo "🎯 **ACADEMIC MODULE COVERAGE SUMMARY:**"
COVERAGE_SUMMARY=$(go tool cover -func=academic_coverage.out | grep "total:" | awk '{print $3}')
if [ -n "$COVERAGE_SUMMARY" ]; then
    echo "   Total Coverage: $COVERAGE_SUMMARY"
else
    echo "   Could not calculate total coverage"
fi

# Ask to generate HTML
read -p "🌐 Generate HTML report for academic module? (y/n): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    go tool cover -html=academic_coverage.out -o academic_coverage.html
    echo "📄 HTML report saved as: academic_coverage.html"
fi

# Clean up
rm -f academic_coverage.out

echo "✨ Done!"
