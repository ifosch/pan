# pan [![Build Status](https://travis-ci.org/EDyO/pan.svg)](https://travis-ci.org/EDyO/pan) [![Code Climate](https://codeclimate.com/github/EDyO/pan/badges/gpa.svg)](https://codeclimate.com/github/EDyO/pan) [![Coverage Status](https://coveralls.io/repos/github/EDyO/pan/badge.svg?branch=master)](https://coveralls.io/github/EDyO/pan?branch=master)

`pan` is a RSS feed management tool.

## Installation

You must setup Go in your computer, and, then, go get pan:

    go get github.com/EDyO/pan
    
## Usage

`pan` allows you to convert a RSS feed in XML format an get a YAML representation, or the other way around:

    pan my_feed.xml > my_feed.yml
    pan my_feed.yml > my_feed.xml
