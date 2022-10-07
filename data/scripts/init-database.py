#!/usr/bin/env python3

import argparse
import pprint


parser = argparse.ArgumentParser(description="Downloads SQL dump files and initialises the DB")
parser.add_argument("--telegram_api_key", required=True)
parser.add_argument("--telegram_chat_id", required=True)
parser.add_argument("--active_dir", default=".")
group = parser.add_mutually_exclusive_group()
group.add_argument("--create_sqlite_file", default="flibusta.db")
group.add_argument("--skip_download", action="store_true")

args = parser.parse_args()
pprint.pprint(args)