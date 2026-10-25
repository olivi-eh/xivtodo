#!/bin/sh
# Regenerates frontend/src/assets/db.json from the CSV exports placed in this directory
set -e
cd "$(dirname "$0")"
uv run csv_to_json.py db.json *.csv
mv db.json ../frontend/src/assets/db.json
