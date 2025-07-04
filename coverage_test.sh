#!/bin/bash
#This is a shell script for bundling all my tests , as go does not automatically targets the files lol
# Run tests and generate coverage profile
go test -coverprofile=cover.out ./internal > /dev/null

# Filter for your target files and sum coverage
FILES="student_service.go|student_placement_service.go"

# Extract relevant lines, sum covered/total, and compute percentage
awk -v files="$FILES" '
BEGIN { covered=0; total=0; }
$1 ~ files {
    n=split($3, arr, "/");
    covered += arr[1];
    total += arr[2];
    printf "%s\n", $0;
}
END {
    if (total > 0) {
        printf "\nTotal for target files: %.1f%% (%d/%d statements covered)\n", (covered/total)*100, covered, total;
    } else {
        print "No coverage data for target files.";
    }
}
' < <(go tool cover -func=cover.out)
