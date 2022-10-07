#!/usr/bin/env python3

import argparse
import logging
import pprint
import os
import sys

import mysql_to_sqlite3 as ms3
import mysql.connector as msc

from subprocess import call
 
parser = argparse.ArgumentParser(
    description="Downloads SQL dump files and initialises the DB")
parser.add_argument("--telegram_api_key", required=True)
parser.add_argument("--telegram_chat_id", required=True)
parser.add_argument("--active_dir", default=".")
parser.add_argument("--sql_dump_file", default="flibusta.sql")
parser.add_argument("--mysql_host")
parser.add_argument("--mysql_port", type=int, default=3306)
parser.add_argument("--mysql_user")
parser.add_argument("--mysql_password")
parser.add_argument("--mysql_database")
parser.add_argument("--create_sqlite_file", default="flibusta.db")
parser.add_argument("--skip_download", action="store_true")

args = parser.parse_args()
pprint.pprint(args)

def CreateMySQLDump() -> bool:
    return os.system("./downloader --dump_file " + args.sql_dump_file) == 0

def ImportMySQLDump() -> bool:
  return call(["mysql", "--host", args.mysql_host, 
      "-u" + args.mysql_user, 
      "-p" + args.mysql_password, 
      "--batch", 
      "-e" + "source " + args.sql_dump_file,
      args.mysql_database],
      stdout=sys.stdout,
      stderr=sys.stderr)

def MySQLtoSqlite():
  msq = ms3.MySQLtoSQLite(
      sqlite_file=args.create_sqlite_file,
      mysql_user=args.mysql_user,
      mysql_password=args.mysql_password,
      mysql_database=args.mysql_database,
      mysql_host=args.mysql_host,
      mysql_port=args.mysql_port,
      debug=True)
  msq.transfer()


if not args.skip_download:
    CreateMySQLDump()

if not ImportMySQLDump():
  logging.log(logging.FATAL, "Couldn't import MySQL dump")
  sys.exit(2)

MySQLtoSqlite()