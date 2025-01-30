#!/bin/bash

if [ `git status --porcelain | wc -l` -gt 0 ]; then
  git --no-pager diff
  git --no-pager status -s
  exit 1
fi

