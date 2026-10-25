import json
import time
from pathlib import Path

import requests

# This script refreshes the FFXIV Collect databases used by the frontend
# for every supported language.
# --> Usage: uv run fetch_collect_json.py

dbs_dir = Path(__file__).resolve().parent.parent / 'frontend' / 'src' / 'assets' / 'dbs'

for lang in ['en', 'fr', 'de', 'ja']:
    for page in [
        'achievements',
        'armoires',
        'bardings',
        'emotes',
        'fashions',
        'frames',
        'hairstyles',
        'mounts',
        'minions',
        'orchestrions',
        'records',
        'relics',
        'spells',
        'survey_records',
        'titles'
    ]:
        url = f'https://ffxivcollect.com/api/{page}?language={lang}'
        print(url)
        response = requests.get(url, timeout=30)
        response.raise_for_status()
        response_json = response.json()
        with open(dbs_dir / lang / f'{page}.json', 'w', encoding='utf-8') as f:
            json.dump(response_json, f)
            f.write('\n')
        time.sleep(1)