#!/usr/bin/env python3
"""
Example of how to parse an gitlab log to find requests.
"""


import argparse
import re
import sys
import time
import json


def getArgs():
    """
    Provie arguments
    """
    parser = argparse.ArgumentParser(
           description='script for finding top ip addresses')


    parser.add_argument('--apijsonlog', required=False, action='store',
                           help='api_json.log file')

    parser.add_argument('--top', required=False, default = 10, action='store',
                           help='Top x of ip adddresses to view.')

    config_args = parser.parse_args()

    return config_args


def createSourceIPList(gitLabLog):
    """
    Extract the source IP addresses
    """
    sourceIPList = []

    for line in gitLabLog:
        ip = re.match(r"^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}", line)
        if ip != None:
            sourceIPList.append(ip.group())

    return sourceIPList


def createSourceIPStats(sourceIPList):
    """
    Create dictionary for source IP stats and sort the list.
    """
    sourceIPStats = {}

    for ip in sourceIPList:
        sourceIPStats.update({ip:sourceIPList.count(ip)})

    #Need to sort the dictionary
    sortedsourceIPStats = [(ipAddr, sourceIPStats[ipAddr]) for ipAddr in sorted(sourceIPStats, key=sourceIPStats.get, reverse=True)]

    return sortedsourceIPStats


def createApiRequestList(apiJsonLog):
    """
    Create list of requests from the api_json.log
    """
    apiRequestList = [json.loads(request) for request in apiJsonLog]

    return apiRequestList

def createRequestStats(apiRequestList):
    """
    Find most requested paths
    """
    apiRequestStats = {}

    pathList = [path['path'] for path in apiRequestList if path['path']]

    for request in pathList:
        #requestPath = request['path']
        #apiRequestStats.update({request:pathList.count(request)})
        print(request)

    #print(apiRequestStats)

def main():

    #Fetch the args
    configArgs = getArgs()

    #Access log to view, default is access.log
    try:
        apiJsonLog = open(configArgs.apijsonlog, "r")
    except Exception as err:
        print("Error opening file:\n{}".format(err))
        sys.exit(1)

    #create a list of all requests from the api_json.log
    apiRequestList = createApiRequestList(apiJsonLog)

    #stats on the number of requests
    createRequestStats(apiRequestList)








if __name__ == "__main__":
    main()