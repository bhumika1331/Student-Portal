#!/bin/bash

# Calculate coverage for specific teacher module files only
# Usage: ./teacher_coverage.sh

# List of files to include in coverage calculation
TEACHER_FILES=(
    "teacher.go"
    "teacher_test.go"
    "teacher_services.go"
    "studentresult_teacher.go"
    "document.go"
    "docupload.go"
    "exportResult_test.go"
    "UploadStudentMark_test.go"
    "attendance.go"
    "attendance_test.go"    
)

echo "🧪 Running tests and generating coverage profile..."

# Run tests with coverage profile
go test -coverprofile=coverage.out ./internal

if [ $? -ne 0 ]; then
    echo "❌ Tests failed!"
    exit 1
fi

echo "📊 Filtering coverage for teacher module files only..."

# Create regex pattern for the files we want
PATTERN=""
for file in "${TEACHER_FILES[@]}"; do
    if [ -z "$PATTERN" ]; then
        PATTERN="$file"
    else
        PATTERN="$PATTERN|$file"
    fi
done

echo "🔍 Looking for files matching: $PATTERN"

# Filter coverage.out to only include our teacher files
# Keep the mode line (first line) and any lines containing our files
awk -v pat="$PATTERN" 'NR==1 || $0 ~ pat' coverage.out > teacher_coverage.out

# Check if we found any coverage data
if [ ! -s teacher_coverage.out ] || [ $(wc -l < teacher_coverage.out) -le 1 ]; then
    echo "⚠  No coverage data found for teacher files!"
    echo "📁 Available files in coverage.out:"
    grep -E "\.go:" coverage.out | cut -d: -f1 | sort | uniq
    exit 1
fi

echo "✅ Found coverage data for teacher files:"
grep -E "\.go:" teacher_coverage.out | cut -d: -f1 | sort | uniq

# Calculate coverage percentage for filtered files
echo ""
echo "📈 Coverage report for teacher module:"
go tool cover -func=teacher_coverage.out

echo ""
echo "🎯 *TEACHER MODULE COVERAGE SUMMARY:*"
COVERAGE_SUMMARY=$(go tool cover -func=teacher_coverage.out | grep "total:" | awk '{print $3}')

if [ -n "$COVERAGE_SUMMARY" ]; then
    echo "   Total Coverage: $COVERAGE_SUMMARY"
else
    echo "   Could not calculate total coverage"
fi

# Optional: Generate HTML report for teacher files only
read -p "🌐 Generate HTML coverage report for teacher files? (y/n): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    go tool cover -html=teacher_coverage.out -o teacher_coverage.html
    echo "📄 HTML report saved as: teacher_coverage.html"
fi

# Cleanup
rm -f teacher_coverage.out

echo "✨ Done!"
