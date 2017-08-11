#!/bin/sh

sed 's/{query}/tester/g' websites.csv > tempfile && mv tempfile websites.csv
