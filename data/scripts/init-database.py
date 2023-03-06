#!/usr/bin/env python3

import argparse
import logging
import pprint
import os
import sys

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

    ok = os.system(cmd) == 0

    if ok:
        cmd = " ".join(["mysql", "--host", args.mysql_host,
                        "--port", str(args.mysql_port),
                        "-u" + args.mysql_user,
                        "-p" + args.mysql_password,
                        "-e'" + MYSQL_FAST + "source ./scripts/sql/mysql-indexes.sql; commit; '",
                        args.mysql_database])
        logging.info(cmd)

        return os.system(cmd) == 0 & ok
try:

    if not args.skip_download:
        CreateMySQLDump()

    if not ImportMySQLDump():
        logging.log(logging.FATAL, "Couldn't import MySQL dump")
        raise Exception("Failed to import SQL dump")

except Exception as e:
    logging.fatal(e)
finally:
    # Clean-up
    logging.info("Cleaning up dump file.")
    os.system("rm " + args.sql_dump_file)
