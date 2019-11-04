#!/bin/sh

rm -rf public/js public/css ./templates/indexpd.plush.html
cp -r ./dist/js public
cp -r ./dist/css public
cp -r ./dist/index.html ./templates/indexpd.plush.html
