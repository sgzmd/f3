#!/usr/bin/env python3

import argparse
from asyncio import subprocess
import logging
import pprint
import os
import sys

import mysql_to_sqlite3 as ms3
import mysql.connector as msc

from subprocess import call, Popen

root = logging.getLogger()
root.setLevel(logging.DEBUG)

handler = logging.StreamHandler(sys.stdout)
handler.setLevel(logging.DEBUG)
formatter = logging.Formatter(
    '%(asctime)s - %(name)s - %(levelname)s - %(message)s')
handler.setFormatter(formatter)
root.addHandler(handler)

parser = argparse.ArgumentParser(
    description="Downloads SQL dump files and initialises the DB")
parser.add_argument("--telegram_api_key", required=True)
parser.add_argument("--telegram_chat_id", required=True)
parser.add_argument("--active_dir", default=".")
parser.add_argument("--sql_dump_file", default="flibusta.sql")
parser.add_argument("--mysql_host", default="localhost")
parser.add_argument("--mysql_port", type=int, default=3306)
parser.add_argument("--mysql_user")
parser.add_argument("--mysql_password")
parser.add_argument("--mysql_database")
parser.add_argument("--create_sqlite_file", default="flibusta.db")
parser.add_argument("--skip_download", action="store_true")

args = parser.parse_args()
pprint.pprint(args)

MYSQL_FAST = """
SET foreign_key_checks=0;
SET autocommit=0;
SET unique_checks=0;

"""


def CreateMySQLDump() -> bool:
    return os.system("./downloader --dump_file " + args.sql_dump_file) == 0


def ImportMySQLDump() -> bool:
    cmd = " ".join(["mysql", "--host", args.mysql_host,
                    "--port", str(args.mysql_port),
                    "-u" + args.mysql_user,
                    "-p" + args.mysql_password,
                    "-e'" + MYSQL_FAST + "source ./" + args.sql_dump_file + "; commit; '",
                    args.mysql_database])
    logging.info(cmd)

    return os.system(cmd) == 0


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


try:

    if not args.skip_download:
        CreateMySQLDump()

    if not ImportMySQLDump():
        logging.log(logging.FATAL, "Couldn't import MySQL dump")
        raise Exception("Failed to import SQL dump")

    MySQLtoSqlite()
except Exception as e:
    logging.fatal(e)
finally:
    # Clean-up
    logging.info("Cleaning up dump file.")
    os.system("rm " + args.sql_dump_file)
