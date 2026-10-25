import csv
import requests
import time
from bs4 import BeautifulSoup

# This script takes the duties CSV file and generate text files
# containing the duty names in each language.
# --> Usage: uv run duty_names.py (from the directory containing duties.csv)

names = {
    "fr": [],
    "de": [],
    "jp": []
}

with open("duties.csv", encoding='utf-8') as csvf: 
    csvReader = csv.DictReader(csvf)

    for row in csvReader:
        lodestoneID = row["LodestoneID"]
        for locale in names:
            if (lodestoneID == ""):
                names[locale].append("")
                continue

            url = "https://" + locale + ".finalfantasyxiv.com/lodestone/playguide/db/duty/" + row["LodestoneID"]
            page = requests.get(url, timeout=30)
            page.raise_for_status()

            soup = BeautifulSoup(page.content, "html.parser")
            results = soup.find("h2", class_="db-view__detail__lname_name")

            name = results.text.strip()
            names[locale].append(name)
            print(name)

            time.sleep(1)

    for locale in names:
        with open("duty_names_" + locale + ".txt", "w", encoding='utf-8') as f:
            for name in names[locale]:
                f.write(name + "\n")