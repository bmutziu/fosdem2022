#!/bin/bash
sed -i '' -E 's/(num:)[ ]*[0-9]/\1  '$1'/' quote.cue 
