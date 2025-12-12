#!/usr/bin/env bash

YEAR=2025
DAY=$1

curl -s \
  --cookie "session=$AOC" \
  "https://adventofcode.com/$YEAR/day/$DAY/input" \
  > Day$DAY/input.txt


