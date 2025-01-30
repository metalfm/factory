#!/bin/bash

if [ `git status --porcelain | wc -l` -gt 0 ]; then
  git --no-pager diff
  echo "Err"
  exit 1
fi

echo "OK"
exit 0
